// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSecretsResponseBody
	GetCode() *string
	SetData(v *ListSecretsResponseBodyData) *ListSecretsResponseBody
	GetData() *ListSecretsResponseBodyData
	SetMessage(v string) *ListSecretsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecretsResponseBody
	GetRequestId() *string
}

type ListSecretsResponseBody struct {
	Code    *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data    *ListSecretsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                      `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSecretsResponseBody) GetData() *ListSecretsResponseBodyData {
	return s.Data
}

func (s *ListSecretsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretsResponseBody) SetCode(v string) *ListSecretsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecretsResponseBody) SetData(v *ListSecretsResponseBodyData) *ListSecretsResponseBody {
	s.Data = v
	return s
}

func (s *ListSecretsResponseBody) SetMessage(v string) *ListSecretsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecretsResponseBody) SetRequestId(v string) *ListSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretsResponseBodyData struct {
	Items      []*ListSecretsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32                              `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalSize  *int32                              `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListSecretsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodyData) GetItems() []*ListSecretsResponseBodyDataItems {
	return s.Items
}

func (s *ListSecretsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListSecretsResponseBodyData) SetItems(v []*ListSecretsResponseBodyDataItems) *ListSecretsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSecretsResponseBodyData) SetPageNumber(v int32) *ListSecretsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsResponseBodyData) SetPageSize(v int32) *ListSecretsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSecretsResponseBodyData) SetTotalSize(v int32) *ListSecretsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListSecretsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecretsResponseBodyDataItems struct {
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

func (s ListSecretsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListSecretsResponseBodyDataItems) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListSecretsResponseBodyDataItems) GetKmsConfig() *KMSConfig {
	return s.KmsConfig
}

func (s *ListSecretsResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListSecretsResponseBodyDataItems) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *ListSecretsResponseBodyDataItems) GetSecretId() *string {
	return s.SecretId
}

func (s *ListSecretsResponseBodyDataItems) GetSecretSource() *string {
	return s.SecretSource
}

func (s *ListSecretsResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListSecretsResponseBodyDataItems) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListSecretsResponseBodyDataItems) SetCreateTimestamp(v int64) *ListSecretsResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetGatewayType(v string) *ListSecretsResponseBodyDataItems {
	s.GatewayType = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetKmsConfig(v *KMSConfig) *ListSecretsResponseBodyDataItems {
	s.KmsConfig = v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetName(v string) *ListSecretsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetReferenceCount(v int32) *ListSecretsResponseBodyDataItems {
	s.ReferenceCount = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetSecretId(v string) *ListSecretsResponseBodyDataItems {
	s.SecretId = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetSecretSource(v string) *ListSecretsResponseBodyDataItems {
	s.SecretSource = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetStatus(v string) *ListSecretsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListSecretsResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListSecretsResponseBodyDataItems) Validate() error {
	if s.KmsConfig != nil {
		if err := s.KmsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
