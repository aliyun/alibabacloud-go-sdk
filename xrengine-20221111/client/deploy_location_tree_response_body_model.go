// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployLocationTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeployLocationTreeResponseBody
	GetCode() *string
	SetData(v bool) *DeployLocationTreeResponseBody
	GetData() *bool
	SetErrorName(v string) *DeployLocationTreeResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *DeployLocationTreeResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *DeployLocationTreeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployLocationTreeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeployLocationTreeResponseBody
	GetSuccess() *bool
}

type DeployLocationTreeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployLocationTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployLocationTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DeployLocationTreeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeployLocationTreeResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeployLocationTreeResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *DeployLocationTreeResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeployLocationTreeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployLocationTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployLocationTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeployLocationTreeResponseBody) SetCode(v string) *DeployLocationTreeResponseBody {
	s.Code = &v
	return s
}

func (s *DeployLocationTreeResponseBody) SetData(v bool) *DeployLocationTreeResponseBody {
	s.Data = &v
	return s
}

func (s *DeployLocationTreeResponseBody) SetErrorName(v string) *DeployLocationTreeResponseBody {
	s.ErrorName = &v
	return s
}

func (s *DeployLocationTreeResponseBody) SetHttpCode(v int32) *DeployLocationTreeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeployLocationTreeResponseBody) SetMessage(v string) *DeployLocationTreeResponseBody {
	s.Message = &v
	return s
}

func (s *DeployLocationTreeResponseBody) SetRequestId(v string) *DeployLocationTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployLocationTreeResponseBody) SetSuccess(v bool) *DeployLocationTreeResponseBody {
	s.Success = &v
	return s
}

func (s *DeployLocationTreeResponseBody) Validate() error {
	return dara.Validate(s)
}
