// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKmsInstanceQuotaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKmsInstanceId(v string) *GetKmsInstanceQuotaInfosRequest
	GetKmsInstanceId() *string
	SetResourceType(v string) *GetKmsInstanceQuotaInfosRequest
	GetResourceType() *string
}

type GetKmsInstanceQuotaInfosRequest struct {
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetKmsInstanceQuotaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceQuotaInfosRequest) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceQuotaInfosRequest) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *GetKmsInstanceQuotaInfosRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetKmsInstanceQuotaInfosRequest) SetKmsInstanceId(v string) *GetKmsInstanceQuotaInfosRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosRequest) SetResourceType(v string) *GetKmsInstanceQuotaInfosRequest {
	s.ResourceType = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosRequest) Validate() error {
	return dara.Validate(s)
}
