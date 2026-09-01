// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int32) *DescribePlaybooksRequest
	GetActive() *int32
	SetEndMillis(v int64) *DescribePlaybooksRequest
	GetEndMillis() *int64
	SetLang(v string) *DescribePlaybooksRequest
	GetLang() *string
	SetName(v string) *DescribePlaybooksRequest
	GetName() *string
	SetOrder(v string) *DescribePlaybooksRequest
	GetOrder() *string
	SetOwnType(v string) *DescribePlaybooksRequest
	GetOwnType() *string
	SetPageNumber(v int64) *DescribePlaybooksRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *DescribePlaybooksRequest
	GetPageSize() *int32
	SetParamTypes(v string) *DescribePlaybooksRequest
	GetParamTypes() *string
	SetPlaybookUuid(v string) *DescribePlaybooksRequest
	GetPlaybookUuid() *string
	SetPlaybookUuids(v string) *DescribePlaybooksRequest
	GetPlaybookUuids() *string
	SetSort(v int32) *DescribePlaybooksRequest
	GetSort() *int32
	SetStartMillis(v int64) *DescribePlaybooksRequest
	GetStartMillis() *int64
}

type DescribePlaybooksRequest struct {
	// The status of the playbook. Valid values:
	//
	// - **1**: The playbook is enabled.
	//
	// - **0**: The playbook is disabled.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The end of the time range to query. This value is a 13-digit timestamp.
	//
	// example:
	//
	// 1683858064361
	EndMillis *int64 `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the playbook.
	//
	// example:
	//
	// demo_playbook
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Default value: **desc**. Valid values:
	//
	// - **desc**: descending.
	//
	// - **asc**: ascending.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The type of the playbook. Valid values:
	//
	// - **preset**: predefined playbook.
	//
	// - **user**: custom playbook.
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 10 entries are returned by default.
	//
	// > Specify a value for this parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trigger type of the playbook. If you do not specify this parameter, playbooks of all trigger types are queried. Valid values:
	//
	// - **template-incident**: security event.
	//
	// - **template-ip**: IP entity.
	//
	// - **template-file**: file entity.
	//
	// - **template-process**: process entity.
	//
	// - **template-alert**: security alert.
	//
	// - **template-domain**: domain name entity.
	//
	// - **template-container**: container entity.
	//
	// - **template-host**: host entity.
	//
	// - **template-custom**: custom.
	//
	// example:
	//
	// template-alert
	ParamTypes *string `json:"ParamTypes,omitempty" xml:"ParamTypes,omitempty"`
	// The UUID of the playbook.
	//
	// > Call the [CreatePlaybook](~~CreatePlaybook~~) operation to obtain this parameter.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// A comma-separated list of playbook UUIDs. You can specify up to 100 UUIDs.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx,8rrt6cff-319e-4ede-97bc-1xxxxxx
	PlaybookUuids *string `json:"PlaybookUuids,omitempty" xml:"PlaybookUuids,omitempty"`
	// The field to sort by. Default value: **1**. Valid values:
	//
	// - **1**: last modification time.
	//
	// - **2**: last running time.
	//
	// example:
	//
	// 1
	Sort *int32 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The start of the time range to query. This value is a 13-digit timestamp.
	//
	// example:
	//
	// 1683526277415
	StartMillis *int64 `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
}

func (s DescribePlaybooksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksRequest) GetActive() *int32 {
	return s.Active
}

func (s *DescribePlaybooksRequest) GetEndMillis() *int64 {
	return s.EndMillis
}

func (s *DescribePlaybooksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybooksRequest) GetName() *string {
	return s.Name
}

func (s *DescribePlaybooksRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribePlaybooksRequest) GetOwnType() *string {
	return s.OwnType
}

func (s *DescribePlaybooksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePlaybooksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePlaybooksRequest) GetParamTypes() *string {
	return s.ParamTypes
}

func (s *DescribePlaybooksRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybooksRequest) GetPlaybookUuids() *string {
	return s.PlaybookUuids
}

func (s *DescribePlaybooksRequest) GetSort() *int32 {
	return s.Sort
}

func (s *DescribePlaybooksRequest) GetStartMillis() *int64 {
	return s.StartMillis
}

func (s *DescribePlaybooksRequest) SetActive(v int32) *DescribePlaybooksRequest {
	s.Active = &v
	return s
}

func (s *DescribePlaybooksRequest) SetEndMillis(v int64) *DescribePlaybooksRequest {
	s.EndMillis = &v
	return s
}

func (s *DescribePlaybooksRequest) SetLang(v string) *DescribePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybooksRequest) SetName(v string) *DescribePlaybooksRequest {
	s.Name = &v
	return s
}

func (s *DescribePlaybooksRequest) SetOrder(v string) *DescribePlaybooksRequest {
	s.Order = &v
	return s
}

func (s *DescribePlaybooksRequest) SetOwnType(v string) *DescribePlaybooksRequest {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPageNumber(v int64) *DescribePlaybooksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPageSize(v int32) *DescribePlaybooksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybooksRequest) SetParamTypes(v string) *DescribePlaybooksRequest {
	s.ParamTypes = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPlaybookUuid(v string) *DescribePlaybooksRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPlaybookUuids(v string) *DescribePlaybooksRequest {
	s.PlaybookUuids = &v
	return s
}

func (s *DescribePlaybooksRequest) SetSort(v int32) *DescribePlaybooksRequest {
	s.Sort = &v
	return s
}

func (s *DescribePlaybooksRequest) SetStartMillis(v int64) *DescribePlaybooksRequest {
	s.StartMillis = &v
	return s
}

func (s *DescribePlaybooksRequest) Validate() error {
	return dara.Validate(s)
}
