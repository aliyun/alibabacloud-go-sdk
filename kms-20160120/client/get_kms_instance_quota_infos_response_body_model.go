// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKmsInstanceQuotaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsInstanceId(v string) *GetKmsInstanceQuotaInfosResponseBody
	GetKmsInstanceId() *string
	SetKmsInstanceQuotaInfos(v []*GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) *GetKmsInstanceQuotaInfosResponseBody
	GetKmsInstanceQuotaInfos() []*GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos
	SetRequestId(v string) *GetKmsInstanceQuotaInfosResponseBody
	GetRequestId() *string
}

type GetKmsInstanceQuotaInfosResponseBody struct {
	KmsInstanceId         *string                                                      `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
	KmsInstanceQuotaInfos []*GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos `json:"KmsInstanceQuotaInfos,omitempty" xml:"KmsInstanceQuotaInfos,omitempty" type:"Repeated"`
	RequestId             *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKmsInstanceQuotaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceQuotaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceQuotaInfosResponseBody) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *GetKmsInstanceQuotaInfosResponseBody) GetKmsInstanceQuotaInfos() []*GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos {
	return s.KmsInstanceQuotaInfos
}

func (s *GetKmsInstanceQuotaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKmsInstanceQuotaInfosResponseBody) SetKmsInstanceId(v string) *GetKmsInstanceQuotaInfosResponseBody {
	s.KmsInstanceId = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponseBody) SetKmsInstanceQuotaInfos(v []*GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) *GetKmsInstanceQuotaInfosResponseBody {
	s.KmsInstanceQuotaInfos = v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponseBody) SetRequestId(v string) *GetKmsInstanceQuotaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponseBody) Validate() error {
	if s.KmsInstanceQuotaInfos != nil {
		for _, item := range s.KmsInstanceQuotaInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos struct {
	ResourceQuota *int64  `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UsedQuantity  *int64  `json:"UsedQuantity,omitempty" xml:"UsedQuantity,omitempty"`
}

func (s GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) GetResourceQuota() *int64 {
	return s.ResourceQuota
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) GetUsedQuantity() *int64 {
	return s.UsedQuantity
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) SetResourceQuota(v int64) *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos {
	s.ResourceQuota = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) SetResourceType(v string) *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos {
	s.ResourceType = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) SetUsedQuantity(v int64) *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos {
	s.UsedQuantity = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponseBodyKmsInstanceQuotaInfos) Validate() error {
	return dara.Validate(s)
}
