// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckJDBCSourceNetConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionMessage(v string) *CheckJDBCSourceNetConnectionResponseBody
	GetConnectionMessage() *string
	SetConnectionStatus(v string) *CheckJDBCSourceNetConnectionResponseBody
	GetConnectionStatus() *string
	SetRequestId(v string) *CheckJDBCSourceNetConnectionResponseBody
	GetRequestId() *string
}

type CheckJDBCSourceNetConnectionResponseBody struct {
	// Return message: Error message returned when the connection fails, otherwise returns an empty string ("").
	//
	// example:
	//
	// connection timeout
	ConnectionMessage *string `json:"ConnectionMessage,omitempty" xml:"ConnectionMessage,omitempty"`
	// Service status:
	//
	// - Network connected: Success
	//
	// - Network not connected: Failed
	//
	// example:
	//
	// Success
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckJDBCSourceNetConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckJDBCSourceNetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckJDBCSourceNetConnectionResponseBody) GetConnectionMessage() *string {
	return s.ConnectionMessage
}

func (s *CheckJDBCSourceNetConnectionResponseBody) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *CheckJDBCSourceNetConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckJDBCSourceNetConnectionResponseBody) SetConnectionMessage(v string) *CheckJDBCSourceNetConnectionResponseBody {
	s.ConnectionMessage = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionResponseBody) SetConnectionStatus(v string) *CheckJDBCSourceNetConnectionResponseBody {
	s.ConnectionStatus = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionResponseBody) SetRequestId(v string) *CheckJDBCSourceNetConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
