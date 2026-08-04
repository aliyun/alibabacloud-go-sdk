// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyContacterResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyContacterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyContacterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyContacterResponseBody
	GetSuccess() *bool
}

type ModifyContacterResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyContacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContacterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContacterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyContacterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyContacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContacterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyContacterResponseBody) SetCode(v string) *ModifyContacterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyContacterResponseBody) SetMessage(v string) *ModifyContacterResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyContacterResponseBody) SetRequestId(v string) *ModifyContacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContacterResponseBody) SetSuccess(v bool) *ModifyContacterResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyContacterResponseBody) Validate() error {
	return dara.Validate(s)
}
