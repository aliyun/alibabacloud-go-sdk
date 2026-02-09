// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTuringTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *SubmitTuringTaskRequest
	GetDuration() *int32
	SetIdempotentKey(v string) *SubmitTuringTaskRequest
	GetIdempotentKey() *string
	SetImgUrl(v string) *SubmitTuringTaskRequest
	GetImgUrl() *string
	SetResolution(v string) *SubmitTuringTaskRequest
	GetResolution() *string
	SetResourceType(v string) *SubmitTuringTaskRequest
	GetResourceType() *string
	SetTaskType(v string) *SubmitTuringTaskRequest
	GetTaskType() *string
}

type SubmitTuringTaskRequest struct {
	// example:
	//
	// 5
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 82veb0a6wc41asdv80
	IdempotentKey *string `json:"idempotentKey,omitempty" xml:"idempotentKey,omitempty"`
	// example:
	//
	// https://www.example.com/1.png
	ImgUrl *string `json:"imgUrl,omitempty" xml:"imgUrl,omitempty"`
	// example:
	//
	// 720P
	Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// example:
	//
	// basic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// live-wallpaper
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s SubmitTuringTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTuringTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTuringTaskRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *SubmitTuringTaskRequest) GetIdempotentKey() *string {
	return s.IdempotentKey
}

func (s *SubmitTuringTaskRequest) GetImgUrl() *string {
	return s.ImgUrl
}

func (s *SubmitTuringTaskRequest) GetResolution() *string {
	return s.Resolution
}

func (s *SubmitTuringTaskRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *SubmitTuringTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *SubmitTuringTaskRequest) SetDuration(v int32) *SubmitTuringTaskRequest {
	s.Duration = &v
	return s
}

func (s *SubmitTuringTaskRequest) SetIdempotentKey(v string) *SubmitTuringTaskRequest {
	s.IdempotentKey = &v
	return s
}

func (s *SubmitTuringTaskRequest) SetImgUrl(v string) *SubmitTuringTaskRequest {
	s.ImgUrl = &v
	return s
}

func (s *SubmitTuringTaskRequest) SetResolution(v string) *SubmitTuringTaskRequest {
	s.Resolution = &v
	return s
}

func (s *SubmitTuringTaskRequest) SetResourceType(v string) *SubmitTuringTaskRequest {
	s.ResourceType = &v
	return s
}

func (s *SubmitTuringTaskRequest) SetTaskType(v string) *SubmitTuringTaskRequest {
	s.TaskType = &v
	return s
}

func (s *SubmitTuringTaskRequest) Validate() error {
	return dara.Validate(s)
}
