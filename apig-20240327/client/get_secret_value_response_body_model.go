// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecretValueResponseBody
	GetCode() *string
	SetData(v *GetSecretValueResponseBodyData) *GetSecretValueResponseBody
	GetData() *GetSecretValueResponseBodyData
	SetMessage(v string) *GetSecretValueResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecretValueResponseBody
	GetRequestId() *string
}

type GetSecretValueResponseBody struct {
	Code    *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetSecretValueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecretValueResponseBody) GetData() *GetSecretValueResponseBodyData {
	return s.Data
}

func (s *GetSecretValueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretValueResponseBody) SetCode(v string) *GetSecretValueResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretValueResponseBody) SetData(v *GetSecretValueResponseBodyData) *GetSecretValueResponseBody {
	s.Data = v
	return s
}

func (s *GetSecretValueResponseBody) SetMessage(v string) *GetSecretValueResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRequestId(v string) *GetSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecretValueResponseBodyData struct {
	SecretData *string `json:"secretData,omitempty" xml:"secretData,omitempty"`
}

func (s GetSecretValueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBodyData) GetSecretData() *string {
	return s.SecretData
}

func (s *GetSecretValueResponseBodyData) SetSecretData(v string) *GetSecretValueResponseBodyData {
	s.SecretData = &v
	return s
}

func (s *GetSecretValueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
