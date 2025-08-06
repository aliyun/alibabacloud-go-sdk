// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobType(v string) *ListJobInfoRequest
	GetJobType() *string
	SetMediaId(v string) *ListJobInfoRequest
	GetMediaId() *string
}

type ListJobInfoRequest struct {
	// The type of the task. Valid values:
	//
	// 	- transcode
	//
	// 	- snapshot
	//
	// 	- ai
	//
	// This parameter is required.
	//
	// example:
	//
	// transcode
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30e5d7**********bd900764de7c0102
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s ListJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfoRequest) GoString() string {
	return s.String()
}

func (s *ListJobInfoRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListJobInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *ListJobInfoRequest) SetJobType(v string) *ListJobInfoRequest {
	s.JobType = &v
	return s
}

func (s *ListJobInfoRequest) SetMediaId(v string) *ListJobInfoRequest {
	s.MediaId = &v
	return s
}

func (s *ListJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
