// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetTaskInstanceLogRequest
	GetId() *int64
	SetRunNumber(v int32) *GetTaskInstanceLogRequest
	GetRunNumber() *int32
}

type GetTaskInstanceLogRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sequence number of an instance run. Minimum value: 1. By default, the latest run is used.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
}

func (s GetTaskInstanceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceLogRequest) GetId() *int64 {
	return s.Id
}

func (s *GetTaskInstanceLogRequest) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *GetTaskInstanceLogRequest) SetId(v int64) *GetTaskInstanceLogRequest {
	s.Id = &v
	return s
}

func (s *GetTaskInstanceLogRequest) SetRunNumber(v int32) *GetTaskInstanceLogRequest {
	s.RunNumber = &v
	return s
}

func (s *GetTaskInstanceLogRequest) Validate() error {
	return dara.Validate(s)
}
