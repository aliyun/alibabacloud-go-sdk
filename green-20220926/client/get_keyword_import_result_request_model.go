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
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID. Although this parameter is marked as required:false in the specification, it is actually required. If you do not provide this parameter, the service returns a 400 error. You can obtain the task ID from the TaskId field in the response of the AddKeywordsToLib operation when you import keywords in batches from a file.
	//
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
