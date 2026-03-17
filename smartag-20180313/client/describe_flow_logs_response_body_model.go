// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogs(v *DescribeFlowLogsResponseBodyFlowLogs) *DescribeFlowLogsResponseBody
	GetFlowLogs() *DescribeFlowLogsResponseBodyFlowLogs
	SetPageNumber(v int32) *DescribeFlowLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFlowLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFlowLogsResponseBody
	GetTotalCount() *int32
}

type DescribeFlowLogsResponseBody struct {
	FlowLogs *DescribeFlowLogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44F4E2D0-77F7-4DEC-969B-061688946143
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of flow logs.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBody) GetFlowLogs() *DescribeFlowLogsResponseBodyFlowLogs {
	return s.FlowLogs
}

func (s *DescribeFlowLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFlowLogsResponseBody) SetFlowLogs(v *DescribeFlowLogsResponseBodyFlowLogs) *DescribeFlowLogsResponseBody {
	s.FlowLogs = v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetPageNumber(v int32) *DescribeFlowLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetPageSize(v int32) *DescribeFlowLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetRequestId(v string) *DescribeFlowLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetTotalCount(v int32) *DescribeFlowLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) Validate() error {
	if s.FlowLogs != nil {
		if err := s.FlowLogs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFlowLogsResponseBodyFlowLogs struct {
	FlowLogSetType []*DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType `json:"FlowLogSetType,omitempty" xml:"FlowLogSetType,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogsResponseBodyFlowLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetFlowLogSetType() []*DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	return s.FlowLogSetType
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetFlowLogSetType(v []*DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) *DescribeFlowLogsResponseBodyFlowLogs {
	s.FlowLogSetType = v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) Validate() error {
	if s.FlowLogSetType != nil {
		for _, item := range s.FlowLogSetType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType struct {
	ActiveAging       *int32  `json:"ActiveAging,omitempty" xml:"ActiveAging,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogId         *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	InactiveAging     *int32  `json:"InactiveAging,omitempty" xml:"InactiveAging,omitempty"`
	LogstoreName      *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetflowServerIp   *string `json:"NetflowServerIp,omitempty" xml:"NetflowServerIp,omitempty"`
	NetflowServerPort *string `json:"NetflowServerPort,omitempty" xml:"NetflowServerPort,omitempty"`
	NetflowVersion    *string `json:"NetflowVersion,omitempty" xml:"NetflowVersion,omitempty"`
	OutputType        *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	ProjectName       *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlsRegionId       *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalSagNum       *int32  `json:"TotalSagNum,omitempty" xml:"TotalSagNum,omitempty"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetActiveAging() *int32 {
	return s.ActiveAging
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetInactiveAging() *int32 {
	return s.InactiveAging
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetName() *string {
	return s.Name
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetNetflowServerIp() *string {
	return s.NetflowServerIp
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetNetflowServerPort() *string {
	return s.NetflowServerPort
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetNetflowVersion() *string {
	return s.NetflowVersion
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetOutputType() *string {
	return s.OutputType
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) GetTotalSagNum() *int32 {
	return s.TotalSagNum
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetActiveAging(v int32) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.ActiveAging = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetDescription(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.Description = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetFlowLogId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetInactiveAging(v int32) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.InactiveAging = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetLogstoreName(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.LogstoreName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetName(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.Name = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetNetflowServerIp(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.NetflowServerIp = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetNetflowServerPort(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.NetflowServerPort = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetNetflowVersion(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.NetflowVersion = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetOutputType(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.OutputType = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetProjectName(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetResourceGroupId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetSlsRegionId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetStatus(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.Status = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) SetTotalSagNum(v int32) *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType {
	s.TotalSagNum = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogSetType) Validate() error {
	return dara.Validate(s)
}
