// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliYunSafeCenterResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSimilarSecurityEventsQueryTaskRequest(v *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) *GetAliYunSafeCenterResultRequest
	GetCreateSimilarSecurityEventsQueryTaskRequest() *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest
	SetDescribeInstancesFullStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) *GetAliYunSafeCenterResultRequest
	GetDescribeInstancesFullStatusRequest() *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest
	SetDescribeSecurityEventOperationStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) *GetAliYunSafeCenterResultRequest
	GetDescribeSecurityEventOperationStatusRequest() *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest
	SetDescribeSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest
	GetDescribeSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest
	SetGetAssetDetailByUuidRequest(v *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) *GetAliYunSafeCenterResultRequest
	GetGetAssetDetailByUuidRequest() *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest
	SetHandleSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) *GetAliYunSafeCenterResultRequest
	GetHandleSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest
	SetHandleSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest
	GetHandleSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest
	SetInterfaceCode(v string) *GetAliYunSafeCenterResultRequest
	GetInterfaceCode() *string
	SetListInstancesRequest(v *GetAliYunSafeCenterResultRequestListInstancesRequest) *GetAliYunSafeCenterResultRequest
	GetListInstancesRequest() *GetAliYunSafeCenterResultRequestListInstancesRequest
	SetRegionId(v string) *GetAliYunSafeCenterResultRequest
	GetRegionId() *string
}

type GetAliYunSafeCenterResultRequest struct {
	CreateSimilarSecurityEventsQueryTaskRequest *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest `json:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty" xml:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty" type:"Struct"`
	DescribeInstancesFullStatusRequest          *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest          `json:"DescribeInstancesFullStatusRequest,omitempty" xml:"DescribeInstancesFullStatusRequest,omitempty" type:"Struct"`
	DescribeSecurityEventOperationStatusRequest *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest `json:"DescribeSecurityEventOperationStatusRequest,omitempty" xml:"DescribeSecurityEventOperationStatusRequest,omitempty" type:"Struct"`
	DescribeSimilarSecurityEventsRequest        *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest        `json:"DescribeSimilarSecurityEventsRequest,omitempty" xml:"DescribeSimilarSecurityEventsRequest,omitempty" type:"Struct"`
	GetAssetDetailByUuidRequest                 *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest                 `json:"GetAssetDetailByUuidRequest,omitempty" xml:"GetAssetDetailByUuidRequest,omitempty" type:"Struct"`
	HandleSecurityEventsRequest                 *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest                 `json:"HandleSecurityEventsRequest,omitempty" xml:"HandleSecurityEventsRequest,omitempty" type:"Struct"`
	HandleSimilarSecurityEventsRequest          *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest          `json:"HandleSimilarSecurityEventsRequest,omitempty" xml:"HandleSimilarSecurityEventsRequest,omitempty" type:"Struct"`
	// This parameter is required.
	InterfaceCode        *string                                               `json:"InterfaceCode,omitempty" xml:"InterfaceCode,omitempty"`
	ListInstancesRequest *GetAliYunSafeCenterResultRequestListInstancesRequest `json:"ListInstancesRequest,omitempty" xml:"ListInstancesRequest,omitempty" type:"Struct"`
	RegionId             *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequest) GetCreateSimilarSecurityEventsQueryTaskRequest() *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	return s.CreateSimilarSecurityEventsQueryTaskRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetDescribeInstancesFullStatusRequest() *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest {
	return s.DescribeInstancesFullStatusRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetDescribeSecurityEventOperationStatusRequest() *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	return s.DescribeSecurityEventOperationStatusRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetDescribeSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest {
	return s.DescribeSimilarSecurityEventsRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetGetAssetDetailByUuidRequest() *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest {
	return s.GetAssetDetailByUuidRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetHandleSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	return s.HandleSecurityEventsRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetHandleSimilarSecurityEventsRequest() *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	return s.HandleSimilarSecurityEventsRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetInterfaceCode() *string {
	return s.InterfaceCode
}

func (s *GetAliYunSafeCenterResultRequest) GetListInstancesRequest() *GetAliYunSafeCenterResultRequestListInstancesRequest {
	return s.ListInstancesRequest
}

func (s *GetAliYunSafeCenterResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequest) SetCreateSimilarSecurityEventsQueryTaskRequest(v *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) *GetAliYunSafeCenterResultRequest {
	s.CreateSimilarSecurityEventsQueryTaskRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetDescribeInstancesFullStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) *GetAliYunSafeCenterResultRequest {
	s.DescribeInstancesFullStatusRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetDescribeSecurityEventOperationStatusRequest(v *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) *GetAliYunSafeCenterResultRequest {
	s.DescribeSecurityEventOperationStatusRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetDescribeSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest {
	s.DescribeSimilarSecurityEventsRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetGetAssetDetailByUuidRequest(v *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) *GetAliYunSafeCenterResultRequest {
	s.GetAssetDetailByUuidRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetHandleSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) *GetAliYunSafeCenterResultRequest {
	s.HandleSecurityEventsRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetHandleSimilarSecurityEventsRequest(v *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) *GetAliYunSafeCenterResultRequest {
	s.HandleSimilarSecurityEventsRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetInterfaceCode(v string) *GetAliYunSafeCenterResultRequest {
	s.InterfaceCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetListInstancesRequest(v *GetAliYunSafeCenterResultRequestListInstancesRequest) *GetAliYunSafeCenterResultRequest {
	s.ListInstancesRequest = v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequest) Validate() error {
	if s.CreateSimilarSecurityEventsQueryTaskRequest != nil {
		if err := s.CreateSimilarSecurityEventsQueryTaskRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DescribeInstancesFullStatusRequest != nil {
		if err := s.DescribeInstancesFullStatusRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DescribeSecurityEventOperationStatusRequest != nil {
		if err := s.DescribeSecurityEventOperationStatusRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DescribeSimilarSecurityEventsRequest != nil {
		if err := s.DescribeSimilarSecurityEventsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.GetAssetDetailByUuidRequest != nil {
		if err := s.GetAssetDetailByUuidRequest.Validate(); err != nil {
			return err
		}
	}
	if s.HandleSecurityEventsRequest != nil {
		if err := s.HandleSecurityEventsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.HandleSimilarSecurityEventsRequest != nil {
		if err := s.HandleSimilarSecurityEventsRequest.Validate(); err != nil {
			return err
		}
	}
	if s.ListInstancesRequest != nil {
		if err := s.ListInstancesRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityEventId          *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	SimilarEventScenarioCode *string `json:"SimilarEventScenarioCode,omitempty" xml:"SimilarEventScenarioCode,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) GetSimilarEventScenarioCode() *string {
	return s.SimilarEventScenarioCode
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) SetSecurityEventId(v int64) *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	s.SecurityEventId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) SetSimilarEventScenarioCode(v string) *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest {
	s.SimilarEventScenarioCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestCreateSimilarSecurityEventsQueryTaskRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) SetInstanceId(v []*string) *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest {
	s.InstanceId = v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeInstancesFullStatusRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	TaskId           *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) SetSecurityEventIds(v []*string) *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	s.SecurityEventIds = v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSecurityEventOperationStatusRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) SetTaskId(v int64) *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestDescribeSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) SetUuid(v string) *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest {
	s.Uuid = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestGetAssetDetailByUuidRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest struct {
	AlertType        *string   `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	FileMd5          *string   `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	FilePath         *string   `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip               *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MarkBatch        *string   `json:"MarkBatch,omitempty" xml:"MarkBatch,omitempty"`
	MarkMissParam    *string   `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	OperationCode    *string   `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams  *string   `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark           *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
}

func (s GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetIp() *string {
	return s.Ip
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetMarkBatch() *string {
	return s.MarkBatch
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetMarkMissParam() *string {
	return s.MarkMissParam
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetAlertType(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.AlertType = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetFileMd5(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.FileMd5 = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetFilePath(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.FilePath = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetInstanceId(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetIp(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.Ip = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetMarkBatch(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.MarkBatch = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetMarkMissParam(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetOperationCode(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetOperationParams(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetRemark(v string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.Remark = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) SetSecurityEventIds(v []*string) *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest {
	s.SecurityEventIds = v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest struct {
	AlertType       *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	OperationCode   *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetIp() *string {
	return s.Ip
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetAlertType(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.AlertType = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetInstanceId(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetIp(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.Ip = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetOperationCode(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetOperationParams(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetSourceIp(v string) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) SetTaskId(v int64) *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestHandleSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAliYunSafeCenterResultRequestListInstancesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultRequestListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultRequestListInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) SetInstanceIds(v string) *GetAliYunSafeCenterResultRequestListInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) SetRegionId(v string) *GetAliYunSafeCenterResultRequestListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultRequestListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
