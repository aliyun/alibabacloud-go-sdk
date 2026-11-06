// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskStepResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListTaskStepResponseBody
	GetCount() *int32
	SetData(v []*ListTaskStepResponseBodyData) *ListTaskStepResponseBody
	GetData() []*ListTaskStepResponseBodyData
	SetRequestId(v string) *ListTaskStepResponseBody
	GetRequestId() *string
}

type ListTaskStepResponseBody struct {
	// The total number of records returned in the list.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of step definitions for the specified task type.
	Data []*ListTaskStepResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTaskStepResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskStepResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskStepResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListTaskStepResponseBody) GetData() []*ListTaskStepResponseBodyData {
	return s.Data
}

func (s *ListTaskStepResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskStepResponseBody) SetCount(v int32) *ListTaskStepResponseBody {
	s.Count = &v
	return s
}

func (s *ListTaskStepResponseBody) SetData(v []*ListTaskStepResponseBodyData) *ListTaskStepResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskStepResponseBody) SetRequestId(v string) *ListTaskStepResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskStepResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskStepResponseBodyData struct {
	// The interaction method of the step. Valid values:
	//
	// - **prompt**: Displays step prompts.
	//
	// - **connection**: Provides connection operations to enter the builder or validation environment.
	//
	// - **prompt_connection**: Provides both prompts and connection operations.
	//
	// example:
	//
	// connection
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The Chinese name of the phase to which the step belongs, such as builder environment preparation or test validation.
	//
	// example:
	//
	// 制作环境准备
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The English name of the phase to which the step belongs. This field is returned for image creation steps, but may not be returned for application adaptation steps.
	//
	// example:
	//
	// Builder Environment Preparation
	CategoryEn *string `json:"CategoryEn,omitempty" xml:"CategoryEn,omitempty"`
	// The numeric index of the step definition, which is used to identify the step and its sequence. Do not use this value as the completion percentage.
	//
	// example:
	//
	// 10
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The name identifier of the step, which is used to match the step records returned by GetTask.
	//
	// Steps when Type is set to IMAGE:
	//
	// - **INIT**: Initialize the task.
	//
	// - **CREATE_BUILDER_ENV**: Create the builder environment.
	//
	// - **DEPLOY_BUILDER_ENV**: Deploy the builder environment.
	//
	// - **INIT_BUILDER**: Initialize the builder tool.
	//
	// - **DOWNLOAD_APP**: Download application resources.
	//
	// - **START_BUILD**: Start image creation.
	//
	// - **COMPLETE_INSTALL**: Complete application installation and configuration.
	//
	// - **CREATE_VALIDATE_ENV**: Create the validation environment.
	//
	// - **DEPLOY_VALIDATE_ENV**: Deploy the validation environment.
	//
	// - **COMPLETE_VALIDATE**: Complete image validation.
	//
	// - **ARCHIVE_DATA**: Archive image data.
	//
	// - **CREATE_IMAGE**: Generate the image.
	//
	// - **FINISH_ADAPT**: End the creation process.
	//
	// Steps when Type is set to ADAPTOR:
	//
	// - **INIT**: Initialize the task.
	//
	// - **CREATE_BUILDER_ENV**: Create the builder environment.
	//
	// - **DEPLOY_BUILDER_ENV**: Deploy the builder environment.
	//
	// - **INIT_BUILDER**: Initialize the builder tool.
	//
	// - **DOWNLOAD_APP**: Download application resources.
	//
	// - **START_VHD**: Start application recording.
	//
	// - **BUILD_VHD**: Generate the application VHD.
	//
	// - **UPLOAD_VHD**: Upload the application VHD.
	//
	// - **COMPLETE_BUILDER**: Complete application creation.
	//
	// - **CREATE_VALIDATE_ENV**: Create the validation environment.
	//
	// - **DEPLOY_VALIDATE_ENV**: Deploy the validation environment.
	//
	// - **COMPLETE_VALIDATE**: Complete application validation.
	//
	// - **CREATE_IMAGE**: Generate the application image.
	//
	// - **FINISH_ADAPT**: End the application adaptation process.
	//
	// example:
	//
	// CREATE_BUILDER_ENV
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTaskStepResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTaskStepResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskStepResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *ListTaskStepResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ListTaskStepResponseBodyData) GetCategoryEn() *string {
	return s.CategoryEn
}

func (s *ListTaskStepResponseBodyData) GetIndex() *int64 {
	return s.Index
}

func (s *ListTaskStepResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListTaskStepResponseBodyData) SetAction(v string) *ListTaskStepResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListTaskStepResponseBodyData) SetCategory(v string) *ListTaskStepResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListTaskStepResponseBodyData) SetCategoryEn(v string) *ListTaskStepResponseBodyData {
	s.CategoryEn = &v
	return s
}

func (s *ListTaskStepResponseBodyData) SetIndex(v int64) *ListTaskStepResponseBodyData {
	s.Index = &v
	return s
}

func (s *ListTaskStepResponseBodyData) SetName(v string) *ListTaskStepResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTaskStepResponseBodyData) Validate() error {
	return dara.Validate(s)
}
