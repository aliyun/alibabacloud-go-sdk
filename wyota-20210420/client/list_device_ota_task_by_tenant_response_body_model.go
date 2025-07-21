// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceOtaTaskByTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDeviceOtaTaskByTenantResponseBody
	GetCode() *string
	SetData(v *ListDeviceOtaTaskByTenantResponseBodyData) *ListDeviceOtaTaskByTenantResponseBody
	GetData() *ListDeviceOtaTaskByTenantResponseBodyData
	SetMessage(v string) *ListDeviceOtaTaskByTenantResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListDeviceOtaTaskByTenantResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeviceOtaTaskByTenantResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDeviceOtaTaskByTenantResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDeviceOtaTaskByTenantResponseBody
	GetTotalCount() *int64
}

type ListDeviceOtaTaskByTenantResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *ListDeviceOtaTaskByTenantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceOtaTaskByTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetData() *ListDeviceOtaTaskByTenantResponseBodyData {
	return s.Data
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceOtaTaskByTenantResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetCode(v string) *ListDeviceOtaTaskByTenantResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetData(v *ListDeviceOtaTaskByTenantResponseBodyData) *ListDeviceOtaTaskByTenantResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetMessage(v string) *ListDeviceOtaTaskByTenantResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetPageNumber(v int32) *ListDeviceOtaTaskByTenantResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetPageSize(v int32) *ListDeviceOtaTaskByTenantResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetRequestId(v string) *ListDeviceOtaTaskByTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetTotalCount(v int64) *ListDeviceOtaTaskByTenantResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeviceOtaTaskByTenantResponseBodyData struct {
	TenantDeviceOtaTasks []*ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks `json:"TenantDeviceOtaTasks,omitempty" xml:"TenantDeviceOtaTasks,omitempty" type:"Repeated"`
}

func (s ListDeviceOtaTaskByTenantResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponseBodyData) GetTenantDeviceOtaTasks() []*ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	return s.TenantDeviceOtaTasks
}

func (s *ListDeviceOtaTaskByTenantResponseBodyData) SetTenantDeviceOtaTasks(v []*ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) *ListDeviceOtaTaskByTenantResponseBodyData {
	s.TenantDeviceOtaTasks = v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks struct {
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OperationStatus *int32  `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	PublishTime     *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpgradeCount    *int64  `json:"UpgradeCount,omitempty" xml:"UpgradeCount,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetModel() *string {
	return s.Model
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetOperationStatus() *int32 {
	return s.OperationStatus
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetStatus() *int32 {
	return s.Status
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetUpgradeCount() *int64 {
	return s.UpgradeCount
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GetVersion() *string {
	return s.Version
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetModel(v string) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.Model = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetOperationStatus(v int32) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.OperationStatus = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetPublishTime(v string) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.PublishTime = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetStatus(v int32) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.Status = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetTaskId(v int32) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.TaskId = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetUpgradeCount(v int64) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.UpgradeCount = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetVersion(v string) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.Version = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) Validate() error {
	return dara.Validate(s)
}
