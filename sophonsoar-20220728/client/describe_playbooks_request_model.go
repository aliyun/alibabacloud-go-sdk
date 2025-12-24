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
	// Activation status of the playbook. Values:
	//
	// - **1**: Indicates the playbook is activated.
	//
	// - **0**: Indicates the playbook is not activated.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// End time for the query, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683858064361
	EndMillis *int64 `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	// Specifies the language type for the request and response, default is **zh**. Values:
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
	// The sorting logic, with a default value of **desc**. Values:
	//
	// - **desc**: Descending order.
	//
	// - **asc**: Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Type of the playbook. Values:
	//
	// - **preset**: Predefined playbook.
	//
	// - **user**: Custom playbook.
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// Sets the page number from which to start displaying the query results. The default value is 1, indicating the first page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specifies the maximum number of items to display per page in a paginated query. The default number of items per page is 20. If the PageSize parameter is empty, it will return 10 items by default.
	//
	// > It is recommended that the PageSize value is not empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trigger method for the playbook, with a default value of **query all**. Values:
	//
	// - **template-incident**: Security incident.
	//
	// - **template-ip**: IP entity.
	//
	// - **template-file**: File entity.
	//
	// - **template-process**: Process entity.
	//
	// - **template-alert**: Security alert.
	//
	// - **template-domain**: Domain entity.
	//
	// - **template-container**: Container entity.
	//
	// - **template-host**: Host entity.
	//
	// - **template-custom**: Custom.
	//
	// example:
	//
	// template-alert
	ParamTypes *string `json:"ParamTypes,omitempty" xml:"ParamTypes,omitempty"`
	// The UUID of the playbook.
	//
	// > You can use the UUID to query specific playbook information.
	//
	// > - Call the [CreatePlaybook](~~CreatePlaybook~~) API to obtain this parameter.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// UUID List of the playbook.
	//
	// Note You can use the UUID list to query specific playbook information.
	//
	// Call the DescribePlaybooks API to obtain this parameter.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx,7745e6cff-319e-4ede-97bc-1xxxxxx
	PlaybookUuids *string `json:"PlaybookUuids,omitempty" xml:"PlaybookUuids,omitempty"`
	// The sorting basis, with a default value of **1**. Values:
	//
	// - **1**: Last modified time.
	//
	// - **2**: Most recent execution time.
	//
	// example:
	//
	// 1
	Sort *int32 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start time for the query, in 13-digit timestamp format.
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
