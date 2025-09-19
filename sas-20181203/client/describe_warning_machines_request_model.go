// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarningMachinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeWarningMachinesRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeWarningMachinesRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeWarningMachinesRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v int32) *DescribeWarningMachinesRequest
	GetCurrentPage() *int32
	SetGroupId(v int64) *DescribeWarningMachinesRequest
	GetGroupId() *int64
	SetHaveRisk(v int32) *DescribeWarningMachinesRequest
	GetHaveRisk() *int32
	SetLang(v string) *DescribeWarningMachinesRequest
	GetLang() *string
	SetMachineName(v string) *DescribeWarningMachinesRequest
	GetMachineName() *string
	SetPageSize(v int32) *DescribeWarningMachinesRequest
	GetPageSize() *int32
	SetRiskId(v int64) *DescribeWarningMachinesRequest
	GetRiskId() *int64
	SetSourceIp(v string) *DescribeWarningMachinesRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *DescribeWarningMachinesRequest
	GetStrategyId() *int64
	SetTargetType(v string) *DescribeWarningMachinesRequest
	GetTargetType() *string
	SetUuids(v string) *DescribeWarningMachinesRequest
	GetUuids() *string
}

type DescribeWarningMachinesRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c7e3c5b420a7947c2933303144688****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the field that is used to search for the container. Valid values:
	//
	// 	- **CONTAINER_ID**: the ID of the container
	//
	// 	- **IMAGE**: the name of the image
	//
	// 	- **NAMESPACE**: the namespace
	//
	// 	- **NODE_NAME**: the name of the node
	//
	// 	- **POD_IP**: the IP address of the pod
	//
	// 	- **HOST_IP**: the IP address of the host
	//
	// 	- **INSTANCE_ID**: the ID of the instance
	//
	// example:
	//
	// containerId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the field that is used to search for the container.
	//
	// example:
	//
	// c8bb3ef0f5ccf45508f0fd1ffc200****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset group.
	//
	// > You can call the [DescribeAllGroups](https://help.aliyun.com/document_detail/130972.html) operation to query the IDs of asset groups.
	//
	// example:
	//
	// 123
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether risks were detected. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	HaveRisk *int32 `json:"HaveRisk,omitempty" xml:"HaveRisk,omitempty"`
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
	// The name of the server on which the baseline check is performed.
	//
	// example:
	//
	// oracle-win-001****
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The number of entries per page. Default value: **10**, which indicates that 10 entries of server information are displayed on each page. A maximum of 100 entries can be returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the risk item.
	//
	// > You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the IDs of risk items.
	//
	// This parameter is required.
	//
	// example:
	//
	// 196
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 16
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **containerId**: the ID of the container
	//
	// 	- **uuid**: the UUID of the asset
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUID of the server on which the baseline check is performed. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// 9888955c-0076-49da-bd9c-34f5492b****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeWarningMachinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningMachinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeWarningMachinesRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeWarningMachinesRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeWarningMachinesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWarningMachinesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeWarningMachinesRequest) GetHaveRisk() *int32 {
	return s.HaveRisk
}

func (s *DescribeWarningMachinesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWarningMachinesRequest) GetMachineName() *string {
	return s.MachineName
}

func (s *DescribeWarningMachinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWarningMachinesRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeWarningMachinesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWarningMachinesRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWarningMachinesRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeWarningMachinesRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeWarningMachinesRequest) SetClusterId(v string) *DescribeWarningMachinesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetContainerFieldName(v string) *DescribeWarningMachinesRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetContainerFieldValue(v string) *DescribeWarningMachinesRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetCurrentPage(v int32) *DescribeWarningMachinesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetGroupId(v int64) *DescribeWarningMachinesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetHaveRisk(v int32) *DescribeWarningMachinesRequest {
	s.HaveRisk = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetLang(v string) *DescribeWarningMachinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetMachineName(v string) *DescribeWarningMachinesRequest {
	s.MachineName = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetPageSize(v int32) *DescribeWarningMachinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetRiskId(v int64) *DescribeWarningMachinesRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetSourceIp(v string) *DescribeWarningMachinesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetStrategyId(v int64) *DescribeWarningMachinesRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetTargetType(v string) *DescribeWarningMachinesRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetUuids(v string) *DescribeWarningMachinesRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeWarningMachinesRequest) Validate() error {
	return dara.Validate(s)
}
