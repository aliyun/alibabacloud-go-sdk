// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHadoopNetConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionMessage(v string) *CheckHadoopNetConnectionResponseBody
	GetConnectionMessage() *string
	SetConnectionStatus(v string) *CheckHadoopNetConnectionResponseBody
	GetConnectionStatus() *string
	SetRequestId(v string) *CheckHadoopNetConnectionResponseBody
	GetRequestId() *string
}

type CheckHadoopNetConnectionResponseBody struct {
	// The response message. If the connection fails, an error message is returned. Otherwise, an empty string is returned.
	//
	// example:
	//
	// connection timeout
	ConnectionMessage *string `json:"ConnectionMessage,omitempty" xml:"ConnectionMessage,omitempty"`
	// The connection status. Valid values:
	//
	// - Success: The network is connected.
	//
	// - Failed: The network is disconnected.
	//
	// example:
	//
	// Success
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckHadoopNetConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckHadoopNetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckHadoopNetConnectionResponseBody) GetConnectionMessage() *string {
	return s.ConnectionMessage
}

func (s *CheckHadoopNetConnectionResponseBody) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *CheckHadoopNetConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckHadoopNetConnectionResponseBody) SetConnectionMessage(v string) *CheckHadoopNetConnectionResponseBody {
	s.ConnectionMessage = &v
	return s
}

func (s *CheckHadoopNetConnectionResponseBody) SetConnectionStatus(v string) *CheckHadoopNetConnectionResponseBody {
	s.ConnectionStatus = &v
	return s
}

func (s *CheckHadoopNetConnectionResponseBody) SetRequestId(v string) *CheckHadoopNetConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckHadoopNetConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
