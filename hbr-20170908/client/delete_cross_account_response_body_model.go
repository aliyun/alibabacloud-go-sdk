// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrossAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCrossAccountResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteCrossAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCrossAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCrossAccountResponseBody
	GetSuccess() *bool
}

type DeleteCrossAccountResponseBody struct {
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
	// 921E9989-735C-5254-A29F-6B8A5DDC1ED1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCrossAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrossAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrossAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCrossAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCrossAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCrossAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCrossAccountResponseBody) SetCode(v string) *DeleteCrossAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCrossAccountResponseBody) SetMessage(v string) *DeleteCrossAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCrossAccountResponseBody) SetRequestId(v string) *DeleteCrossAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCrossAccountResponseBody) SetSuccess(v bool) *DeleteCrossAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCrossAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
