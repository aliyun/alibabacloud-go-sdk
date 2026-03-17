// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRemoteAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySagRemoteAccessResponseBody
	GetCode() *string
	SetMessage(v string) *ModifySagRemoteAccessResponseBody
	GetMessage() *string
	SetRemoteAccessIp(v string) *ModifySagRemoteAccessResponseBody
	GetRemoteAccessIp() *string
	SetRequestId(v string) *ModifySagRemoteAccessResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *ModifySagRemoteAccessResponseBody
	GetSerialNumber() *string
	SetSuccess(v bool) *ModifySagRemoteAccessResponseBody
	GetSuccess() *bool
}

type ModifySagRemoteAccessResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned after calling the API.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The remote access IP address.
	//
	// example:
	//
	// 10.10.10.2
	RemoteAccessIp *string `json:"RemoteAccessIp,omitempty" xml:"RemoteAccessIp,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4FF203D7-462D-498E-94F9-2B2FA462DD23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SN of the SAG device.
	//
	// example:
	//
	// sag61a344**
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySagRemoteAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRemoteAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagRemoteAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySagRemoteAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySagRemoteAccessResponseBody) GetRemoteAccessIp() *string {
	return s.RemoteAccessIp
}

func (s *ModifySagRemoteAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagRemoteAccessResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ModifySagRemoteAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySagRemoteAccessResponseBody) SetCode(v string) *ModifySagRemoteAccessResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySagRemoteAccessResponseBody) SetMessage(v string) *ModifySagRemoteAccessResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySagRemoteAccessResponseBody) SetRemoteAccessIp(v string) *ModifySagRemoteAccessResponseBody {
	s.RemoteAccessIp = &v
	return s
}

func (s *ModifySagRemoteAccessResponseBody) SetRequestId(v string) *ModifySagRemoteAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagRemoteAccessResponseBody) SetSerialNumber(v string) *ModifySagRemoteAccessResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *ModifySagRemoteAccessResponseBody) SetSuccess(v bool) *ModifySagRemoteAccessResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySagRemoteAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
