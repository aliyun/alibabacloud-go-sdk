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
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The knowledge base job ID, which is the `Data.Id` parameter returned by the [SubmitIndexJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexjob) or [SubmitIndexAddDocumentsJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexadddocumentsjob) operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230718xxxx-146c93bf
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Both the [SubmitIndexJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexjob) and [SubmitIndexAddDocumentsJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexadddocumentsjob) operations support batch import of documents. This operation returns both the overall `Status` of the job and the `Document.Status` of each document. If there are a large number of documents, you can use the `PageNumber` parameter to perform a paged query. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of document import jobs that are displayed on each page. No maximum value. Default value: 10.
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
