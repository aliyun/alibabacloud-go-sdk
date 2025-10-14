// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluateAndImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteEvaluateAndImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *DeleteEvaluateAndImportTaskRequest
	GetSlinkTaskId() *string
}

type DeleteEvaluateAndImportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s DeleteEvaluateAndImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluateAndImportTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteEvaluateAndImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEvaluateAndImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DeleteEvaluateAndImportTaskRequest) SetRegionId(v string) *DeleteEvaluateAndImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskRequest) SetSlinkTaskId(v string) *DeleteEvaluateAndImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
