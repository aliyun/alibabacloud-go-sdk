// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageRemovalProResponseBody
	GetCode() *string
	SetData(v *ImageRemovalProResponseBodyData) *ImageRemovalProResponseBody
	GetData() *ImageRemovalProResponseBodyData
	SetMessage(v string) *ImageRemovalProResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageRemovalProResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageRemovalProResponseBody
	GetSuccess() *bool
}

type ImageRemovalProResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The intelligent removal Pro result.
	Data *ImageRemovalProResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageRemovalProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProResponseBody) GoString() string {
	return s.String()
}

func (s *ImageRemovalProResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageRemovalProResponseBody) GetData() *ImageRemovalProResponseBodyData {
	return s.Data
}

func (s *ImageRemovalProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageRemovalProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageRemovalProResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageRemovalProResponseBody) SetCode(v string) *ImageRemovalProResponseBody {
	s.Code = &v
	return s
}

func (s *ImageRemovalProResponseBody) SetData(v *ImageRemovalProResponseBodyData) *ImageRemovalProResponseBody {
	s.Data = v
	return s
}

func (s *ImageRemovalProResponseBody) SetMessage(v string) *ImageRemovalProResponseBody {
	s.Message = &v
	return s
}

func (s *ImageRemovalProResponseBody) SetRequestId(v string) *ImageRemovalProResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageRemovalProResponseBody) SetSuccess(v bool) *ImageRemovalProResponseBody {
	s.Success = &v
	return s
}

func (s *ImageRemovalProResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageRemovalProResponseBodyData struct {
	// The URL of the result image. This parameter is returned in synchronous mode.
	//
	// example:
	//
	// https://aib-image.oss-ap-southeast-1.aliyuncs.com/ai_desc%2F250cc947-9cd5-4df0-9c23-44eba5d0dfc30.jpg?OSSAccessKeyId=LTAI5tSEGjGp5wixZgHLc3bV&Expires=4999655814&Signature=shvGNDmkyv9MLTw4%2BOxYglJCpAE%3D
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The asynchronous task ID. This parameter is returned in asynchronous mode and is used to query the task result.
	//
	// example:
	//
	// ed9d8504-6141-9fbb-8345-4fa36433483f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage details. This parameter is returned in synchronous mode.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s ImageRemovalProResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageRemovalProResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemovalProResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageRemovalProResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageRemovalProResponseBodyData) SetImageUrl(v string) *ImageRemovalProResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemovalProResponseBodyData) SetTaskId(v string) *ImageRemovalProResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageRemovalProResponseBodyData) SetUsageMap(v map[string]*int64) *ImageRemovalProResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageRemovalProResponseBodyData) Validate() error {
	return dara.Validate(s)
}
