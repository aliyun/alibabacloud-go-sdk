// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexId(v string) *GetIndexJobStatusRequest
	GetIndexId() *string
	SetJobId(v string) *GetIndexJobStatusRequest
	GetJobId() *string
	SetPageNumber(v int32) *GetIndexJobStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetIndexJobStatusRequest
	GetPageSize() *int32
}

type GetIndexJobStatusRequest struct {
	// The knowledge base ID, which is the Data.Id returned by the CreateIndex operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The knowledge base job task ID, which is the Data.Id returned by the SubmitIndexJob or SubmitIndexAddDocumentsJob operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230718xxxx-146c93bf
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Both the **SubmitIndexJob*	- and **SubmitIndexAddDocumentsJob*	- operations support batch file import. The fields returned by this operation include the overall status of the knowledge base job task (`Status`) and the import status of each file (`Document.Status`). If a large number of files are imported, you can use the `PageNumber` parameter for paging. The starting value is 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of file import tasks to return on each page during paging. Maximum value: unlimited. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetIndexJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIndexJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *GetIndexJobStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetIndexJobStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetIndexJobStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetIndexJobStatusRequest) SetIndexId(v string) *GetIndexJobStatusRequest {
	s.IndexId = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetJobId(v string) *GetIndexJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetPageNumber(v int32) *GetIndexJobStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetPageSize(v int32) *GetIndexJobStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetIndexJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
