// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVaultResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteVaultResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVaultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVaultResponseBody
	GetSuccess() *bool
}

type DeleteVaultResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVaultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVaultResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVaultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVaultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVaultResponseBody) SetCode(v string) *DeleteVaultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVaultResponseBody) SetMessage(v string) *DeleteVaultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVaultResponseBody) SetRequestId(v string) *DeleteVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVaultResponseBody) SetSuccess(v bool) *DeleteVaultResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVaultResponseBody) Validate() error {
	return dara.Validate(s)
}
