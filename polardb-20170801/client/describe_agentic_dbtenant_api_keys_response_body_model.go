// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBTenantApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAgenticDBTenantApiKeysResponseBodyItems) *DescribeAgenticDBTenantApiKeysResponseBody
	GetItems() []*DescribeAgenticDBTenantApiKeysResponseBodyItems
	SetPageNumber(v int32) *DescribeAgenticDBTenantApiKeysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBTenantApiKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAgenticDBTenantApiKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAgenticDBTenantApiKeysResponseBody
	GetTotalCount() *int32
}

type DescribeAgenticDBTenantApiKeysResponseBody struct {
	Items []*DescribeAgenticDBTenantApiKeysResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C3D4E5F6-A7B8-9012-CDEF-123456789012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAgenticDBTenantApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBTenantApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) GetItems() []*DescribeAgenticDBTenantApiKeysResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) SetItems(v []*DescribeAgenticDBTenantApiKeysResponseBodyItems) *DescribeAgenticDBTenantApiKeysResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) SetPageNumber(v int32) *DescribeAgenticDBTenantApiKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) SetPageSize(v int32) *DescribeAgenticDBTenantApiKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) SetRequestId(v string) *DescribeAgenticDBTenantApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) SetTotalCount(v int32) *DescribeAgenticDBTenantApiKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBody) Validate() error {
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

type DescribeAgenticDBTenantApiKeysResponseBodyItems struct {
	// example:
	//
	// ak-71304e39c7e841a1
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// pagc_key_cGFnYy1icDFhMmIz****
	ApiKeyMasked *string `json:"ApiKeyMasked,omitempty" xml:"ApiKeyMasked,omitempty"`
	// example:
	//
	// 2026-06-10T10:30:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MCP server token
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// my-saas-app
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeAgenticDBTenantApiKeysResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBTenantApiKeysResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetApiKeyMasked() *string {
	return s.ApiKeyMasked
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetApiKeyId(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.ApiKeyId = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetApiKeyMasked(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.ApiKeyMasked = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetCreateTime(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetDescription(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetExpireTime(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetStatus(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetTenantId(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) SetTenantName(v string) *DescribeAgenticDBTenantApiKeysResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
