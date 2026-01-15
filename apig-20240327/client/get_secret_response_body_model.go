// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecretResponseBody
	GetCode() *string
	SetData(v *GetSecretResponseBodyData) *GetSecretResponseBody
	GetData() *GetSecretResponseBodyData
	SetMessage(v string) *GetSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecretResponseBody
	GetRequestId() *string
}

type GetSecretResponseBody struct {
	Code    *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetSecretResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                    `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecretResponseBody) GetData() *GetSecretResponseBodyData {
	return s.Data
}

func (s *GetSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretResponseBody) SetCode(v string) *GetSecretResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretResponseBody) SetData(v *GetSecretResponseBodyData) *GetSecretResponseBody {
	s.Data = v
	return s
}

func (s *GetSecretResponseBody) SetMessage(v string) *GetSecretResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretResponseBody) SetRequestId(v string) *GetSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecretResponseBodyData struct {
	CreateTimestamp *int64     `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	GatewayType     *string    `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	KmsConfig       *KMSConfig `json:"kmsConfig,omitempty" xml:"kmsConfig,omitempty"`
	Name            *string    `json:"name,omitempty" xml:"name,omitempty"`
	ReferenceCount  *int32     `json:"referenceCount,omitempty" xml:"referenceCount,omitempty"`
	SecretId        *string    `json:"secretId,omitempty" xml:"secretId,omitempty"`
	SecretSource    *string    `json:"secretSource,omitempty" xml:"secretSource,omitempty"`
	Status          *string    `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTimestamp *int64     `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecretResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetSecretResponseBodyData) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetSecretResponseBodyData) GetKmsConfig() *KMSConfig {
	return s.KmsConfig
}

func (s *GetSecretResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSecretResponseBodyData) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *GetSecretResponseBodyData) GetSecretId() *string {
	return s.SecretId
}

func (s *GetSecretResponseBodyData) GetSecretSource() *string {
	return s.SecretSource
}

func (s *GetSecretResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSecretResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetSecretResponseBodyData) SetCreateTimestamp(v int64) *GetSecretResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetSecretResponseBodyData) SetGatewayType(v string) *GetSecretResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *GetSecretResponseBodyData) SetKmsConfig(v *KMSConfig) *GetSecretResponseBodyData {
	s.KmsConfig = v
	return s
}

func (s *GetSecretResponseBodyData) SetName(v string) *GetSecretResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSecretResponseBodyData) SetReferenceCount(v int32) *GetSecretResponseBodyData {
	s.ReferenceCount = &v
	return s
}

func (s *GetSecretResponseBodyData) SetSecretId(v string) *GetSecretResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *GetSecretResponseBodyData) SetSecretSource(v string) *GetSecretResponseBodyData {
	s.SecretSource = &v
	return s
}

func (s *GetSecretResponseBodyData) SetStatus(v string) *GetSecretResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSecretResponseBodyData) SetUpdateTimestamp(v int64) *GetSecretResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetSecretResponseBodyData) Validate() error {
	if s.KmsConfig != nil {
		if err := s.KmsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
