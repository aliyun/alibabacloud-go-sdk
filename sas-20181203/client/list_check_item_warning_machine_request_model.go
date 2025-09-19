// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemWarningMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ListCheckItemWarningMachineRequest
	GetCheckId() *int64
	SetContainerFieldName(v string) *ListCheckItemWarningMachineRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *ListCheckItemWarningMachineRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v int32) *ListCheckItemWarningMachineRequest
	GetCurrentPage() *int32
	SetGroupId(v int64) *ListCheckItemWarningMachineRequest
	GetGroupId() *int64
	SetLang(v string) *ListCheckItemWarningMachineRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckItemWarningMachineRequest
	GetPageSize() *int32
	SetRemark(v string) *ListCheckItemWarningMachineRequest
	GetRemark() *string
	SetRiskType(v string) *ListCheckItemWarningMachineRequest
	GetRiskType() *string
	SetSource(v string) *ListCheckItemWarningMachineRequest
	GetSource() *string
	SetStatus(v int32) *ListCheckItemWarningMachineRequest
	GetStatus() *int32
	SetUuidList(v []*string) *ListCheckItemWarningMachineRequest
	GetUuidList() []*string
}

type ListCheckItemWarningMachineRequest struct {
	// The ID of the check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the field that is used to query containers.
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the field that is used to query containers.
	//
	// example:
	//
	// ce89cdd0ea732472a8703821b19e****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset group.
	//
	// > You can call the [DescribeAllGroups](https://help.aliyun.com/document_detail/130972.html) operation to query the ID of the asset group.
	//
	// example:
	//
	// 1161****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used to query servers in fuzzy match mode.
	//
	// example:
	//
	// 225
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the check item.
	//
	// example:
	//
	// cis
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// The data source. Default value: default. Valid values:
	//
	// 	- **default**: The check items of baselines for hosts.
	//
	// 	- **agentless**: The check items of baselines for agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- 1: failed
	//
	// 	- 2: verifying
	//
	// 	- 3: passed
	//
	// 	- 6: ignored
	//
	// 	- 7: fixing
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID array of the servers.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s ListCheckItemWarningMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineRequest) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckItemWarningMachineRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *ListCheckItemWarningMachineRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *ListCheckItemWarningMachineRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemWarningMachineRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListCheckItemWarningMachineRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckItemWarningMachineRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemWarningMachineRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListCheckItemWarningMachineRequest) GetRiskType() *string {
	return s.RiskType
}

func (s *ListCheckItemWarningMachineRequest) GetSource() *string {
	return s.Source
}

func (s *ListCheckItemWarningMachineRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListCheckItemWarningMachineRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ListCheckItemWarningMachineRequest) SetCheckId(v int64) *ListCheckItemWarningMachineRequest {
	s.CheckId = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetContainerFieldName(v string) *ListCheckItemWarningMachineRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetContainerFieldValue(v string) *ListCheckItemWarningMachineRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetCurrentPage(v int32) *ListCheckItemWarningMachineRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetGroupId(v int64) *ListCheckItemWarningMachineRequest {
	s.GroupId = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetLang(v string) *ListCheckItemWarningMachineRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetPageSize(v int32) *ListCheckItemWarningMachineRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetRemark(v string) *ListCheckItemWarningMachineRequest {
	s.Remark = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetRiskType(v string) *ListCheckItemWarningMachineRequest {
	s.RiskType = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetSource(v string) *ListCheckItemWarningMachineRequest {
	s.Source = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetStatus(v int32) *ListCheckItemWarningMachineRequest {
	s.Status = &v
	return s
}

func (s *ListCheckItemWarningMachineRequest) SetUuidList(v []*string) *ListCheckItemWarningMachineRequest {
	s.UuidList = v
	return s
}

func (s *ListCheckItemWarningMachineRequest) Validate() error {
	return dara.Validate(s)
}
