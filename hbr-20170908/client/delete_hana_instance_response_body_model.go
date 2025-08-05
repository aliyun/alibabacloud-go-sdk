// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHanaInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHanaInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHanaInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHanaInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHanaInstanceResponseBody
	GetSuccess() *bool
}

type DeleteHanaInstanceResponseBody struct {
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
	// 28EAF89A-E0D8-5C04-9A1D-B373B29BCFB9
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

func (s DeleteHanaInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHanaInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHanaInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHanaInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHanaInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHanaInstanceResponseBody) SetCode(v string) *DeleteHanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) SetMessage(v string) *DeleteHanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) SetRequestId(v string) *DeleteHanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) SetSuccess(v bool) *DeleteHanaInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
