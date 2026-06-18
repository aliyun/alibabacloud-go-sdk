// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOuterAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddOuterAccountResponseBody
	GetCode() *string
	SetData(v string) *AddOuterAccountResponseBody
	GetData() *string
	SetMessage(v string) *AddOuterAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOuterAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddOuterAccountResponseBody
	GetSuccess() *bool
}

type AddOuterAccountResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Created Account ID.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddOuterAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOuterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddOuterAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddOuterAccountResponseBody) GetData() *string {
	return s.Data
}

func (s *AddOuterAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOuterAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOuterAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddOuterAccountResponseBody) SetCode(v string) *AddOuterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetData(v string) *AddOuterAccountResponseBody {
	s.Data = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetMessage(v string) *AddOuterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetRequestId(v string) *AddOuterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetSuccess(v bool) *AddOuterAccountResponseBody {
	s.Success = &v
	return s
}

func (s *AddOuterAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
