// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccountContactDeleteResponseBody
	GetCode() *string
	SetData(v bool) *AccountContactDeleteResponseBody
	GetData() *bool
	SetMessage(v string) *AccountContactDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *AccountContactDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccountContactDeleteResponseBody
	GetSuccess() *bool
}

type AccountContactDeleteResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {\\"count\\": 1}
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8CCD6B37-98E7-5A68-B1F7-A900C9BFF45C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccountContactDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountContactDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *AccountContactDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccountContactDeleteResponseBody) GetData() *bool {
	return s.Data
}

func (s *AccountContactDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccountContactDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountContactDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountContactDeleteResponseBody) SetCode(v string) *AccountContactDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *AccountContactDeleteResponseBody) SetData(v bool) *AccountContactDeleteResponseBody {
	s.Data = &v
	return s
}

func (s *AccountContactDeleteResponseBody) SetMessage(v string) *AccountContactDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *AccountContactDeleteResponseBody) SetRequestId(v string) *AccountContactDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountContactDeleteResponseBody) SetSuccess(v bool) *AccountContactDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *AccountContactDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
