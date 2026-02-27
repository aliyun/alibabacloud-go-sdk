// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataServiceAppSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetDataServiceAppSecretResponseBody
	GetCode() *string
	SetData(v *ResetDataServiceAppSecretResponseBodyData) *ResetDataServiceAppSecretResponseBody
	GetData() *ResetDataServiceAppSecretResponseBodyData
	SetHttpStatusCode(v int32) *ResetDataServiceAppSecretResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetDataServiceAppSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetDataServiceAppSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetDataServiceAppSecretResponseBody
	GetSuccess() *bool
}

type ResetDataServiceAppSecretResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ResetDataServiceAppSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetDataServiceAppSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDataServiceAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDataServiceAppSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetDataServiceAppSecretResponseBody) GetData() *ResetDataServiceAppSecretResponseBodyData {
	return s.Data
}

func (s *ResetDataServiceAppSecretResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetDataServiceAppSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetDataServiceAppSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDataServiceAppSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetDataServiceAppSecretResponseBody) SetCode(v string) *ResetDataServiceAppSecretResponseBody {
	s.Code = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBody) SetData(v *ResetDataServiceAppSecretResponseBodyData) *ResetDataServiceAppSecretResponseBody {
	s.Data = v
	return s
}

func (s *ResetDataServiceAppSecretResponseBody) SetHttpStatusCode(v int32) *ResetDataServiceAppSecretResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBody) SetMessage(v string) *ResetDataServiceAppSecretResponseBody {
	s.Message = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBody) SetRequestId(v string) *ResetDataServiceAppSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBody) SetSuccess(v bool) *ResetDataServiceAppSecretResponseBody {
	s.Success = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetDataServiceAppSecretResponseBodyData struct {
	// example:
	//
	// 200000001
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// abc123456789
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
}

func (s ResetDataServiceAppSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetDataServiceAppSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetDataServiceAppSecretResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *ResetDataServiceAppSecretResponseBodyData) GetAppSecret() *string {
	return s.AppSecret
}

func (s *ResetDataServiceAppSecretResponseBodyData) SetAppKey(v string) *ResetDataServiceAppSecretResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBodyData) SetAppSecret(v string) *ResetDataServiceAppSecretResponseBodyData {
	s.AppSecret = &v
	return s
}

func (s *ResetDataServiceAppSecretResponseBodyData) Validate() error {
	return dara.Validate(s)
}
