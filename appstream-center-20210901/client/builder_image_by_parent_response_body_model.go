// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuilderImageByParentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BuilderImageByParentResponseBody
	GetCode() *string
	SetData(v *BuilderImageByParentResponseBodyData) *BuilderImageByParentResponseBody
	GetData() *BuilderImageByParentResponseBodyData
	SetMessage(v string) *BuilderImageByParentResponseBody
	GetMessage() *string
	SetRequestId(v string) *BuilderImageByParentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuilderImageByParentResponseBody
	GetSuccess() *bool
}

type BuilderImageByParentResponseBody struct {
	// The error code. This parameter is returned when the request fails. A successful response may not contain this parameter.
	//
	// example:
	//
	// ExceedBuilderQuota
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned by this operation.
	Data *BuilderImageByParentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error or result message. This parameter describes the cause when the request fails. A successful response may not contain this parameter.
	//
	// example:
	//
	// The maximum number of builders is reached.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is processed successfully. A value of true indicates that the request is processed successfully. To check whether the asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuilderImageByParentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuilderImageByParentResponseBody) GoString() string {
	return s.String()
}

func (s *BuilderImageByParentResponseBody) GetCode() *string {
	return s.Code
}

func (s *BuilderImageByParentResponseBody) GetData() *BuilderImageByParentResponseBodyData {
	return s.Data
}

func (s *BuilderImageByParentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BuilderImageByParentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuilderImageByParentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuilderImageByParentResponseBody) SetCode(v string) *BuilderImageByParentResponseBody {
	s.Code = &v
	return s
}

func (s *BuilderImageByParentResponseBody) SetData(v *BuilderImageByParentResponseBodyData) *BuilderImageByParentResponseBody {
	s.Data = v
	return s
}

func (s *BuilderImageByParentResponseBody) SetMessage(v string) *BuilderImageByParentResponseBody {
	s.Message = &v
	return s
}

func (s *BuilderImageByParentResponseBody) SetRequestId(v string) *BuilderImageByParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuilderImageByParentResponseBody) SetSuccess(v bool) *BuilderImageByParentResponseBody {
	s.Success = &v
	return s
}

func (s *BuilderImageByParentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuilderImageByParentResponseBodyData struct {
	// The image ID, which is used for subsequent queries or modifications of the image.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The task ID, which can be used to call GetTask to query the task status.
	//
	// example:
	//
	// tid-bp1example****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The image version ID, which is used to query or manage the specified version of the image.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s BuilderImageByParentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BuilderImageByParentResponseBodyData) GoString() string {
	return s.String()
}

func (s *BuilderImageByParentResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *BuilderImageByParentResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BuilderImageByParentResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *BuilderImageByParentResponseBodyData) SetImageId(v string) *BuilderImageByParentResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *BuilderImageByParentResponseBodyData) SetTaskId(v string) *BuilderImageByParentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BuilderImageByParentResponseBodyData) SetVersionId(v string) *BuilderImageByParentResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *BuilderImageByParentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
