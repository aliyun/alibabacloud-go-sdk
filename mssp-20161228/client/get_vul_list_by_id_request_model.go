// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulListByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetVulListByIdRequest
	GetCurrentPage() *int32
	SetDealed(v string) *GetVulListByIdRequest
	GetDealed() *string
	SetId(v int64) *GetVulListByIdRequest
	GetId() *int64
	SetNecessity(v string) *GetVulListByIdRequest
	GetNecessity() *string
	SetPageSize(v int32) *GetVulListByIdRequest
	GetPageSize() *int32
	SetRemark(v string) *GetVulListByIdRequest
	GetRemark() *string
	SetUuids(v string) *GetVulListByIdRequest
	GetUuids() *string
}

type GetVulListByIdRequest struct {
	// Current page
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Whether it has been processed; y: processed; n: not processed
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 4209205
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Risk level
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Asset information of the vulnerability to be queried, which can be set as asset name, public IP, or private IP.
	//
	// example:
	//
	// production_nat_cn-hangzhou_zone_105
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// UUID of the server with the vulnerability to be queried. Multiple UUIDs should be separated by a comma (,).
	//
	// example:
	//
	// 3615b908-995a-4edb-bc85-1981b4e94ba0,9c52cf9a-d8ba-4e31-ae06-500b879ee4e6,4b7de3cf-c4ac-42fc-8804-35070493dc29,f3c01525-0777-4c97-88d9-bec11afd4a6a,a80bd516-c4f3-4c27-a169-c8abfaf9e89e
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s GetVulListByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulListByIdRequest) GoString() string {
	return s.String()
}

func (s *GetVulListByIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetVulListByIdRequest) GetDealed() *string {
	return s.Dealed
}

func (s *GetVulListByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *GetVulListByIdRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *GetVulListByIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVulListByIdRequest) GetRemark() *string {
	return s.Remark
}

func (s *GetVulListByIdRequest) GetUuids() *string {
	return s.Uuids
}

func (s *GetVulListByIdRequest) SetCurrentPage(v int32) *GetVulListByIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetVulListByIdRequest) SetDealed(v string) *GetVulListByIdRequest {
	s.Dealed = &v
	return s
}

func (s *GetVulListByIdRequest) SetId(v int64) *GetVulListByIdRequest {
	s.Id = &v
	return s
}

func (s *GetVulListByIdRequest) SetNecessity(v string) *GetVulListByIdRequest {
	s.Necessity = &v
	return s
}

func (s *GetVulListByIdRequest) SetPageSize(v int32) *GetVulListByIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetVulListByIdRequest) SetRemark(v string) *GetVulListByIdRequest {
	s.Remark = &v
	return s
}

func (s *GetVulListByIdRequest) SetUuids(v string) *GetVulListByIdRequest {
	s.Uuids = &v
	return s
}

func (s *GetVulListByIdRequest) Validate() error {
	return dara.Validate(s)
}
