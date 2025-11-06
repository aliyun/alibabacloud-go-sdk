// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStaticAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListStaticAccountsResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *ListStaticAccountsResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *ListStaticAccountsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStaticAccountsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStaticAccountsResponseBody
	GetSuccess() *bool
}

type ListStaticAccountsResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListStaticAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStaticAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStaticAccountsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListStaticAccountsResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ListStaticAccountsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStaticAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStaticAccountsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStaticAccountsResponseBody) SetCode(v int32) *ListStaticAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *ListStaticAccountsResponseBody) SetData(v map[string]interface{}) *ListStaticAccountsResponseBody {
	s.Data = v
	return s
}

func (s *ListStaticAccountsResponseBody) SetMessage(v string) *ListStaticAccountsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStaticAccountsResponseBody) SetRequestId(v string) *ListStaticAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStaticAccountsResponseBody) SetSuccess(v bool) *ListStaticAccountsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStaticAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
