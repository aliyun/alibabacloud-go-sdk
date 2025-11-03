// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChains(v []*ListChainResponseBodyChains) *ListChainResponseBody
	GetChains() []*ListChainResponseBodyChains
	SetCode(v string) *ListChainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListChainResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListChainResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChainResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChainResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListChainResponseBody
	GetTotalCount() *int32
}

type ListChainResponseBody struct {
	// The list of delivery chains.
	Chains []*ListChainResponseBodyChains `json:"Chains,omitempty" xml:"Chains,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85A99B10-3926-5201-958E-C06FA47F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of delivery chains.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListChainResponseBody) GetChains() []*ListChainResponseBodyChains {
	return s.Chains
}

func (s *ListChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListChainResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChainResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChainResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListChainResponseBody) SetChains(v []*ListChainResponseBodyChains) *ListChainResponseBody {
	s.Chains = v
	return s
}

func (s *ListChainResponseBody) SetCode(v string) *ListChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListChainResponseBody) SetIsSuccess(v bool) *ListChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChainResponseBody) SetPageNo(v int32) *ListChainResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChainResponseBody) SetPageSize(v int32) *ListChainResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChainResponseBody) SetRequestId(v string) *ListChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChainResponseBody) SetTotalCount(v int32) *ListChainResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChainResponseBody) Validate() error {
	if s.Chains != nil {
		for _, item := range s.Chains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChainResponseBodyChains struct {
	// The ID of the delivery chain.
	//
	// example:
	//
	// chi-0ops0gsmw5x2****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The time when the delivery chain was created.
	//
	// example:
	//
	// 1638255427000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery chain.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the delivery chain was last modified.
	//
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the delivery chain.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Repositories to which the delivery chain does not apply.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
	// The ID of the delivery chain scope.
	//
	// example:
	//
	// crr-nyrh2oko32xb****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The type of the delivery chain scope.
	//
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s ListChainResponseBodyChains) String() string {
	return dara.Prettify(s)
}

func (s ListChainResponseBodyChains) GoString() string {
	return s.String()
}

func (s *ListChainResponseBodyChains) GetChainId() *string {
	return s.ChainId
}

func (s *ListChainResponseBodyChains) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListChainResponseBodyChains) GetDescription() *string {
	return s.Description
}

func (s *ListChainResponseBodyChains) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChainResponseBodyChains) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListChainResponseBodyChains) GetName() *string {
	return s.Name
}

func (s *ListChainResponseBodyChains) GetScopeExclude() []*string {
	return s.ScopeExclude
}

func (s *ListChainResponseBodyChains) GetScopeId() *string {
	return s.ScopeId
}

func (s *ListChainResponseBodyChains) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListChainResponseBodyChains) SetChainId(v string) *ListChainResponseBodyChains {
	s.ChainId = &v
	return s
}

func (s *ListChainResponseBodyChains) SetCreateTime(v int64) *ListChainResponseBodyChains {
	s.CreateTime = &v
	return s
}

func (s *ListChainResponseBodyChains) SetDescription(v string) *ListChainResponseBodyChains {
	s.Description = &v
	return s
}

func (s *ListChainResponseBodyChains) SetInstanceId(v string) *ListChainResponseBodyChains {
	s.InstanceId = &v
	return s
}

func (s *ListChainResponseBodyChains) SetModifiedTime(v int64) *ListChainResponseBodyChains {
	s.ModifiedTime = &v
	return s
}

func (s *ListChainResponseBodyChains) SetName(v string) *ListChainResponseBodyChains {
	s.Name = &v
	return s
}

func (s *ListChainResponseBodyChains) SetScopeExclude(v []*string) *ListChainResponseBodyChains {
	s.ScopeExclude = v
	return s
}

func (s *ListChainResponseBodyChains) SetScopeId(v string) *ListChainResponseBodyChains {
	s.ScopeId = &v
	return s
}

func (s *ListChainResponseBodyChains) SetScopeType(v string) *ListChainResponseBodyChains {
	s.ScopeType = &v
	return s
}

func (s *ListChainResponseBodyChains) Validate() error {
	return dara.Validate(s)
}
