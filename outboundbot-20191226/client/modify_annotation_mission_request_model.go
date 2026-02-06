// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnnotationMissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotationMissionId(v string) *ModifyAnnotationMissionRequest
	GetAnnotationMissionId() *string
	SetAnnotationMissionName(v string) *ModifyAnnotationMissionRequest
	GetAnnotationMissionName() *string
	SetAnnotationStatus(v int32) *ModifyAnnotationMissionRequest
	GetAnnotationStatus() *int32
	SetDelete(v bool) *ModifyAnnotationMissionRequest
	GetDelete() *bool
}

type ModifyAnnotationMissionRequest struct {
	// example:
	//
	// c88cc004-de69-4eee-aa5f-2efed533a54e
	AnnotationMissionId   *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionName *string `json:"AnnotationMissionName,omitempty" xml:"AnnotationMissionName,omitempty"`
	// example:
	//
	// 2
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// example:
	//
	// false
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
}

func (s ModifyAnnotationMissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnnotationMissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnnotationMissionRequest) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ModifyAnnotationMissionRequest) GetAnnotationMissionName() *string {
	return s.AnnotationMissionName
}

func (s *ModifyAnnotationMissionRequest) GetAnnotationStatus() *int32 {
	return s.AnnotationStatus
}

func (s *ModifyAnnotationMissionRequest) GetDelete() *bool {
	return s.Delete
}

func (s *ModifyAnnotationMissionRequest) SetAnnotationMissionId(v string) *ModifyAnnotationMissionRequest {
	s.AnnotationMissionId = &v
	return s
}

func (s *ModifyAnnotationMissionRequest) SetAnnotationMissionName(v string) *ModifyAnnotationMissionRequest {
	s.AnnotationMissionName = &v
	return s
}

func (s *ModifyAnnotationMissionRequest) SetAnnotationStatus(v int32) *ModifyAnnotationMissionRequest {
	s.AnnotationStatus = &v
	return s
}

func (s *ModifyAnnotationMissionRequest) SetDelete(v bool) *ModifyAnnotationMissionRequest {
	s.Delete = &v
	return s
}

func (s *ModifyAnnotationMissionRequest) Validate() error {
	return dara.Validate(s)
}
