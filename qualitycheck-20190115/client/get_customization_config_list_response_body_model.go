// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizationConfigListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomizationConfigListResponseBody
	GetCode() *string
	SetData(v *GetCustomizationConfigListResponseBodyData) *GetCustomizationConfigListResponseBody
	GetData() *GetCustomizationConfigListResponseBodyData
	SetMessage(v string) *GetCustomizationConfigListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomizationConfigListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomizationConfigListResponseBody
	GetSuccess() *bool
}

type GetCustomizationConfigListResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomizationConfigListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomizationConfigListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizationConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomizationConfigListResponseBody) GetData() *GetCustomizationConfigListResponseBodyData {
	return s.Data
}

func (s *GetCustomizationConfigListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomizationConfigListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomizationConfigListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomizationConfigListResponseBody) SetCode(v string) *GetCustomizationConfigListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetData(v *GetCustomizationConfigListResponseBodyData) *GetCustomizationConfigListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetMessage(v string) *GetCustomizationConfigListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetRequestId(v string) *GetCustomizationConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetSuccess(v bool) *GetCustomizationConfigListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCustomizationConfigListResponseBodyData struct {
	ModelCustomizationDataSetPo []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo `json:"ModelCustomizationDataSetPo,omitempty" xml:"ModelCustomizationDataSetPo,omitempty" type:"Repeated"`
}

func (s GetCustomizationConfigListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizationConfigListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBodyData) GetModelCustomizationDataSetPo() []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	return s.ModelCustomizationDataSetPo
}

func (s *GetCustomizationConfigListResponseBodyData) SetModelCustomizationDataSetPo(v []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) *GetCustomizationConfigListResponseBodyData {
	s.ModelCustomizationDataSetPo = v
	return s
}

func (s *GetCustomizationConfigListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo struct {
	AsrVersion *int32 `json:"AsrVersion,omitempty" xml:"AsrVersion,omitempty"`
	// example:
	//
	// 2019-01-08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cdae396590bb479a9ec40f3476e274fc
	ModeCustomizationId *string `json:"ModeCustomizationId,omitempty" xml:"ModeCustomizationId,omitempty"`
	// example:
	//
	// 1
	ModelId   *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 5
	ModelStatus *int32 `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetAsrVersion() *int32 {
	return s.AsrVersion
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetModeCustomizationId() *string {
	return s.ModeCustomizationId
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetModelId() *int64 {
	return s.ModelId
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetModelName() *string {
	return s.ModelName
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetModelStatus() *int32 {
	return s.ModelStatus
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetAsrVersion(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.AsrVersion = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetCreateTime(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.CreateTime = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModeCustomizationId(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModeCustomizationId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelId(v int64) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelName(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelName = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelStatus(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelStatus = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetTaskType(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.TaskType = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) Validate() error {
	return dara.Validate(s)
}
