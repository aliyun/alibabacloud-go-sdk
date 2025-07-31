// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordImportResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetKeywordImportResultRequest
	GetRegionId() *string
	SetTaskId(v string) *GetKeywordImportResultRequest
	GetTaskId() *string
}

type GetKeywordImportResultRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// xxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetKeywordImportResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordImportResultRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKeywordImportResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKeywordImportResultRequest) SetRegionId(v string) *GetKeywordImportResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetKeywordImportResultRequest) SetTaskId(v string) *GetKeywordImportResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetKeywordImportResultRequest) Validate() error {
	return dara.Validate(s)
}
