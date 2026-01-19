// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsForResourceServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOrganizationalUnitsForResourceServerResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOrganizationalUnitsForResourceServerResponseBody
	GetNextToken() *string
	SetOrganizationalUnits(v []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) *ListOrganizationalUnitsForResourceServerResponseBody
	GetOrganizationalUnits() []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits
	SetRequestId(v string) *ListOrganizationalUnitsForResourceServerResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOrganizationalUnitsForResourceServerResponseBody
	GetTotalCount() *int64
}

type ListOrganizationalUnitsForResourceServerResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken           *string                                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrganizationalUnits []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrganizationalUnitsForResourceServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForResourceServerResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) GetOrganizationalUnits() []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits {
	return s.OrganizationalUnits
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) SetMaxResults(v int32) *ListOrganizationalUnitsForResourceServerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) SetNextToken(v string) *ListOrganizationalUnitsForResourceServerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) SetOrganizationalUnits(v []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) *ListOrganizationalUnitsForResourceServerResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) SetRequestId(v string) *ListOrganizationalUnitsForResourceServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsForResourceServerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBody) Validate() error {
	if s.OrganizationalUnits != nil {
		for _, item := range s.OrganizationalUnits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits struct {
	// 实例唯一标识
	//
	// example:
	//
	// idaas_qsw77zl5vrllwzyrrfwbmpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 组织的唯一标识
	//
	// example:
	//
	// ou_nbsomva32b6utec3hgi7scxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// 资源服务Scope权限集合
	ResourceServerScopes []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes `json:"ResourceServerScopes,omitempty" xml:"ResourceServerScopes,omitempty" type:"Repeated"`
}

func (s ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) GetResourceServerScopes() []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes {
	return s.ResourceServerScopes
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) SetInstanceId(v string) *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) SetResourceServerScopes(v []*ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits {
	s.ResourceServerScopes = v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnits) Validate() error {
	if s.ResourceServerScopes != nil {
		for _, item := range s.ResourceServerScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes struct {
	// ResourceServerScope唯一标识
	//
	// example:
	//
	// ress_nbte4bb3qqqnaq73rlmkqixxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
	// ResourceServerScope名称
	//
	// example:
	//
	// 获取资源信息
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
}

func (s ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) SetResourceServerScopeId(v string) *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) SetResourceServerScopeName(v string) *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes {
	s.ResourceServerScopeName = &v
	return s
}

func (s *ListOrganizationalUnitsForResourceServerResponseBodyOrganizationalUnitsResourceServerScopes) Validate() error {
	return dara.Validate(s)
}
