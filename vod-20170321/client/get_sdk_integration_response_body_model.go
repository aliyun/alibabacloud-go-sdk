// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSdkIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSdkIntegrationResponseBody
	GetCode() *string
	SetData(v string) *GetSdkIntegrationResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetSdkIntegrationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSdkIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSdkIntegrationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSdkIntegrationResponseBody
	GetSuccess() *bool
}

type GetSdkIntegrationResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSdkIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSdkIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSdkIntegrationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSdkIntegrationResponseBody) GetData() *string {
	return s.Data
}

func (s *GetSdkIntegrationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSdkIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSdkIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSdkIntegrationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSdkIntegrationResponseBody) SetCode(v string) *GetSdkIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *GetSdkIntegrationResponseBody) SetData(v string) *GetSdkIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *GetSdkIntegrationResponseBody) SetHttpStatusCode(v int32) *GetSdkIntegrationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSdkIntegrationResponseBody) SetMessage(v string) *GetSdkIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *GetSdkIntegrationResponseBody) SetRequestId(v string) *GetSdkIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSdkIntegrationResponseBody) SetSuccess(v bool) *GetSdkIntegrationResponseBody {
	s.Success = &v
	return s
}

func (s *GetSdkIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}
