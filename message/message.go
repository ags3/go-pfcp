// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "io"

// MessageType definitions.
const (
	MsgTypeHeartbeatRequest            uint8 = 1
	MsgTypeHeartbeatResponse           uint8 = 2
	MsgTypePFDManagementRequest        uint8 = 3
	MsgTypePFDManagementResponse       uint8 = 4
	MsgTypeAssociationSetupRequest     uint8 = 5
	MsgTypeAssociationSetupResponse    uint8 = 6
	MsgTypeAssociationUpdateRequest    uint8 = 7
	MsgTypeAssociationUpdateResponse   uint8 = 8
	MsgTypeAssociationReleaseRequest   uint8 = 9
	MsgTypeAssociationReleaseResponse  uint8 = 10
	MsgTypeVersionNotSupportedResponse uint8 = 11
	MsgTypeNodeReportRequest           uint8 = 12
	MsgTypeNodeReportResponse          uint8 = 13
	MsgTypeSessionSetDeletionRequest   uint8 = 14
	MsgTypeSessionSetDeletionResponse  uint8 = 15

	// 16 to 49: For future use

	MsgTypeSessionEstablishmentRequest  uint8 = 50
	MsgTypeSessionEstablishmentResponse uint8 = 51
	MsgTypeSessionModificationRequest   uint8 = 52
	MsgTypeSessionModificationResponse  uint8 = 53
	MsgTypeSessionDeletionRequest       uint8 = 54
	MsgTypeSessionDeletionResponse      uint8 = 55
	MsgTypeSessionReportRequest         uint8 = 56
	MsgTypeSessionReportResponse        uint8 = 57

	// 58 to 99: for future use
)

// Message is an interface that defines PFCP messages.
type Message interface {
	MarshalTo([]byte) error
	UnmarshalBinary(b []byte) error
	MarshalLen() int
	MessageType() uint8
	MessageTypeName() string
	Version() int
	SEID() uint64
	Sequence() uint32
}

// Parse parses the given bytes as Message.
func Parse(b []byte) (Message, error) {
	if len(b) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	var m Message
	switch b[1] {
	case MsgTypeHeartbeatRequest:
		m = &HeartbeatRequest{}
	case MsgTypeHeartbeatResponse:
		m = &HeartbeatResponse{}
	default:
		m = &Generic{}
	}

	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}
