// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyBrandUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateProxyBrandUserResponseBody
	GetCode() *int32
	SetData(v bool) *CreateProxyBrandUserResponseBody
	GetData() *bool
	SetErrorMsg(v string) *CreateProxyBrandUserResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateProxyBrandUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProxyBrandUserResponseBody
	GetSuccess() *bool
}

type CreateProxyBrandUserResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProxyBrandUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyBrandUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProxyBrandUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateProxyBrandUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateProxyBrandUserResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateProxyBrandUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProxyBrandUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProxyBrandUserResponseBody) SetCode(v int32) *CreateProxyBrandUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProxyBrandUserResponseBody) SetData(v bool) *CreateProxyBrandUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProxyBrandUserResponseBody) SetErrorMsg(v string) *CreateProxyBrandUserResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProxyBrandUserResponseBody) SetRequestId(v string) *CreateProxyBrandUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProxyBrandUserResponseBody) SetSuccess(v bool) *CreateProxyBrandUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProxyBrandUserResponseBody) Validate() error {
	return dara.Validate(s)
}
