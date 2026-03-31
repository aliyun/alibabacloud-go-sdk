// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataSourceResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateDataSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataSourceResponseBody
	GetSuccess() *bool
}

type UpdateDataSourceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1F22A695-***-***-***-DDC66F5E5ACB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourceResponseBody) SetCode(v string) *UpdateDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetMessage(v string) *UpdateDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetSuccess(v bool) *UpdateDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
