// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecretResponseBody
	GetCode() *string
	SetData(v *CreateSecretResponseBodyData) *CreateSecretResponseBody
	GetData() *CreateSecretResponseBodyData
	SetMessage(v string) *CreateSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecretResponseBody
	GetRequestId() *string
}

type CreateSecretResponseBody struct {
	Code    *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CreateSecretResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecretResponseBody) GetData() *CreateSecretResponseBodyData {
	return s.Data
}

func (s *CreateSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretResponseBody) SetCode(v string) *CreateSecretResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecretResponseBody) SetData(v *CreateSecretResponseBodyData) *CreateSecretResponseBody {
	s.Data = v
	return s
}

func (s *CreateSecretResponseBody) SetMessage(v string) *CreateSecretResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecretResponseBody) SetRequestId(v string) *CreateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecretResponseBodyData struct {
	SecretId *string `json:"secretId,omitempty" xml:"secretId,omitempty"`
}

func (s CreateSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBodyData) GetSecretId() *string {
	return s.SecretId
}

func (s *CreateSecretResponseBodyData) SetSecretId(v string) *CreateSecretResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *CreateSecretResponseBodyData) Validate() error {
	return dara.Validate(s)
}
