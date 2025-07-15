// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAudioFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAudioFilesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListAudioFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAudioFilesRequest
	GetPageSize() *int32
	SetStatus(v string) *ListAudioFilesRequest
	GetStatus() *string
	SetUsage(v string) *ListAudioFilesRequest
	GetUsage() *string
}

type ListAudioFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Usage    *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListAudioFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAudioFilesRequest) GoString() string {
	return s.String()
}

func (s *ListAudioFilesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAudioFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAudioFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAudioFilesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAudioFilesRequest) GetUsage() *string {
	return s.Usage
}

func (s *ListAudioFilesRequest) SetInstanceId(v string) *ListAudioFilesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAudioFilesRequest) SetPageNumber(v int32) *ListAudioFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAudioFilesRequest) SetPageSize(v int32) *ListAudioFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAudioFilesRequest) SetStatus(v string) *ListAudioFilesRequest {
	s.Status = &v
	return s
}

func (s *ListAudioFilesRequest) SetUsage(v string) *ListAudioFilesRequest {
	s.Usage = &v
	return s
}

func (s *ListAudioFilesRequest) Validate() error {
	return dara.Validate(s)
}
