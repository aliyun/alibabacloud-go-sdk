// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeAvailableModelsResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAvailableModelsResponseBody
	GetEngineVersion() *string
	SetItems(v []*DescribeAvailableModelsResponseBodyItems) *DescribeAvailableModelsResponseBody
	GetItems() []*DescribeAvailableModelsResponseBodyItems
	SetModelType(v string) *DescribeAvailableModelsResponseBody
	GetModelType() *string
	SetRequestId(v string) *DescribeAvailableModelsResponseBody
	GetRequestId() *string
	SetTuneArch(v string) *DescribeAvailableModelsResponseBody
	GetTuneArch() *string
}

type DescribeAvailableModelsResponseBody struct {
	// The database engine.
	//
	// example:
	//
	// polardb_ai
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 3.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The list of models.
	Items []*DescribeAvailableModelsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The model type corresponding to the request.
	//
	// example:
	//
	// custom
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AA69096-757C-4647-B36C-29EBC2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The model fine-tuning architecture.
	//
	// example:
	//
	// swift
	TuneArch *string `json:"TuneArch,omitempty" xml:"TuneArch,omitempty"`
}

func (s DescribeAvailableModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableModelsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableModelsResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableModelsResponseBody) GetItems() []*DescribeAvailableModelsResponseBodyItems {
	return s.Items
}

func (s *DescribeAvailableModelsResponseBody) GetModelType() *string {
	return s.ModelType
}

func (s *DescribeAvailableModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableModelsResponseBody) GetTuneArch() *string {
	return s.TuneArch
}

func (s *DescribeAvailableModelsResponseBody) SetEngine(v string) *DescribeAvailableModelsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableModelsResponseBody) SetEngineVersion(v string) *DescribeAvailableModelsResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableModelsResponseBody) SetItems(v []*DescribeAvailableModelsResponseBodyItems) *DescribeAvailableModelsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAvailableModelsResponseBody) SetModelType(v string) *DescribeAvailableModelsResponseBody {
	s.ModelType = &v
	return s
}

func (s *DescribeAvailableModelsResponseBody) SetRequestId(v string) *DescribeAvailableModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableModelsResponseBody) SetTuneArch(v string) *DescribeAvailableModelsResponseBody {
	s.TuneArch = &v
	return s
}

func (s *DescribeAvailableModelsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableModelsResponseBodyItems struct {
	// The custom model registration key.
	//
	// example:
	//
	// Qwen3-32B
	CustomModelName *string `json:"CustomModelName,omitempty" xml:"CustomModelName,omitempty"`
	// The display name of the model.
	//
	// example:
	//
	// My Qwen3 32B
	DisplayModelName *string `json:"DisplayModelName,omitempty" xml:"DisplayModelName,omitempty"`
	// The last modified time of the registration.
	//
	// example:
	//
	// 2026-09-07 20:57:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The supported GPU types.
	GpuRequired []*DescribeAvailableModelsResponseBodyItemsGpuRequired `json:"GpuRequired,omitempty" xml:"GpuRequired,omitempty" type:"Repeated"`
	// The minimum number of CPUs.
	//
	// example:
	//
	// 8
	MinimumCpu *int64 `json:"MinimumCpu,omitempty" xml:"MinimumCpu,omitempty"`
	// The minimum memory size.
	//
	// example:
	//
	// 16384
	MinimumMemory *int64 `json:"MinimumMemory,omitempty" xml:"MinimumMemory,omitempty"`
	// The model name.
	//
	// example:
	//
	// Qwen3-32B-GPTQ-Int4
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model series.
	//
	// example:
	//
	// BERT
	ModelSeries *string `json:"ModelSeries,omitempty" xml:"ModelSeries,omitempty"`
	// The model type.
	//
	// example:
	//
	// custom
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The OSS path of the custom model.
	//
	// example:
	//
	// /my-model-bucket/models
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// The supported GPU models.
	SupportedGpuModels []*string `json:"SupportedGpuModels,omitempty" xml:"SupportedGpuModels,omitempty" type:"Repeated"`
	TuneArch           *string   `json:"TuneArch,omitempty" xml:"TuneArch,omitempty"`
}

func (s DescribeAvailableModelsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableModelsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAvailableModelsResponseBodyItems) GetCustomModelName() *string {
	return s.CustomModelName
}

func (s *DescribeAvailableModelsResponseBodyItems) GetDisplayModelName() *string {
	return s.DisplayModelName
}

func (s *DescribeAvailableModelsResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAvailableModelsResponseBodyItems) GetGpuRequired() []*DescribeAvailableModelsResponseBodyItemsGpuRequired {
	return s.GpuRequired
}

func (s *DescribeAvailableModelsResponseBodyItems) GetMinimumCpu() *int64 {
	return s.MinimumCpu
}

func (s *DescribeAvailableModelsResponseBodyItems) GetMinimumMemory() *int64 {
	return s.MinimumMemory
}

func (s *DescribeAvailableModelsResponseBodyItems) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeAvailableModelsResponseBodyItems) GetModelSeries() *string {
	return s.ModelSeries
}

func (s *DescribeAvailableModelsResponseBodyItems) GetModelType() *string {
	return s.ModelType
}

func (s *DescribeAvailableModelsResponseBodyItems) GetOssPath() *string {
	return s.OssPath
}

func (s *DescribeAvailableModelsResponseBodyItems) GetSupportedGpuModels() []*string {
	return s.SupportedGpuModels
}

func (s *DescribeAvailableModelsResponseBodyItems) GetTuneArch() *string {
	return s.TuneArch
}

func (s *DescribeAvailableModelsResponseBodyItems) SetCustomModelName(v string) *DescribeAvailableModelsResponseBodyItems {
	s.CustomModelName = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetDisplayModelName(v string) *DescribeAvailableModelsResponseBodyItems {
	s.DisplayModelName = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetGmtModified(v string) *DescribeAvailableModelsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetGpuRequired(v []*DescribeAvailableModelsResponseBodyItemsGpuRequired) *DescribeAvailableModelsResponseBodyItems {
	s.GpuRequired = v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetMinimumCpu(v int64) *DescribeAvailableModelsResponseBodyItems {
	s.MinimumCpu = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetMinimumMemory(v int64) *DescribeAvailableModelsResponseBodyItems {
	s.MinimumMemory = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetModelName(v string) *DescribeAvailableModelsResponseBodyItems {
	s.ModelName = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetModelSeries(v string) *DescribeAvailableModelsResponseBodyItems {
	s.ModelSeries = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetModelType(v string) *DescribeAvailableModelsResponseBodyItems {
	s.ModelType = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetOssPath(v string) *DescribeAvailableModelsResponseBodyItems {
	s.OssPath = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetSupportedGpuModels(v []*string) *DescribeAvailableModelsResponseBodyItems {
	s.SupportedGpuModels = v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) SetTuneArch(v string) *DescribeAvailableModelsResponseBodyItems {
	s.TuneArch = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItems) Validate() error {
	if s.GpuRequired != nil {
		for _, item := range s.GpuRequired {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableModelsResponseBodyItemsGpuRequired struct {
	// The minimum number of GPUs.
	//
	// example:
	//
	// xx
	GpuMinCount *string `json:"GpuMinCount,omitempty" xml:"GpuMinCount,omitempty"`
	// The GPU model.
	//
	// example:
	//
	// xxx
	GpuModel *string `json:"GpuModel,omitempty" xml:"GpuModel,omitempty"`
}

func (s DescribeAvailableModelsResponseBodyItemsGpuRequired) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableModelsResponseBodyItemsGpuRequired) GoString() string {
	return s.String()
}

func (s *DescribeAvailableModelsResponseBodyItemsGpuRequired) GetGpuMinCount() *string {
	return s.GpuMinCount
}

func (s *DescribeAvailableModelsResponseBodyItemsGpuRequired) GetGpuModel() *string {
	return s.GpuModel
}

func (s *DescribeAvailableModelsResponseBodyItemsGpuRequired) SetGpuMinCount(v string) *DescribeAvailableModelsResponseBodyItemsGpuRequired {
	s.GpuMinCount = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItemsGpuRequired) SetGpuModel(v string) *DescribeAvailableModelsResponseBodyItemsGpuRequired {
	s.GpuModel = &v
	return s
}

func (s *DescribeAvailableModelsResponseBodyItemsGpuRequired) Validate() error {
	return dara.Validate(s)
}
