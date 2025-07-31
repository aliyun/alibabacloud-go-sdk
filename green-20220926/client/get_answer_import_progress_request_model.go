// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnswerImportProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetAnswerImportProgressRequest
	GetRegionId() *string
	SetTaskId(v string) *GetAnswerImportProgressRequest
	GetTaskId() *string
}

type GetAnswerImportProgressRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAnswerImportProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAnswerImportProgressRequest) GoString() string {
	return s.String()
}

func (s *GetAnswerImportProgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAnswerImportProgressRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAnswerImportProgressRequest) SetRegionId(v string) *GetAnswerImportProgressRequest {
	s.RegionId = &v
	return s
}

func (s *GetAnswerImportProgressRequest) SetTaskId(v string) *GetAnswerImportProgressRequest {
	s.TaskId = &v
	return s
}

func (s *GetAnswerImportProgressRequest) Validate() error {
	return dara.Validate(s)
}
