// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorksEmbedListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorksEmbedListResponseBody
	GetRequestId() *string
	SetResult(v *GetWorksEmbedListResponseBodyResult) *GetWorksEmbedListResponseBody
	GetResult() *GetWorksEmbedListResponseBodyResult
	SetSuccess(v bool) *GetWorksEmbedListResponseBody
	GetSuccess() *bool
}

type GetWorksEmbedListResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 38C0F*****0-415****9F1-*****422BDB65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Array of report objects
	Result *GetWorksEmbedListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Whether the request was successful
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorksEmbedListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorksEmbedListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorksEmbedListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorksEmbedListResponseBody) GetResult() *GetWorksEmbedListResponseBodyResult {
	return s.Result
}

func (s *GetWorksEmbedListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorksEmbedListResponseBody) SetRequestId(v string) *GetWorksEmbedListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorksEmbedListResponseBody) SetResult(v *GetWorksEmbedListResponseBodyResult) *GetWorksEmbedListResponseBody {
	s.Result = v
	return s
}

func (s *GetWorksEmbedListResponseBody) SetSuccess(v bool) *GetWorksEmbedListResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorksEmbedListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorksEmbedListResponseBodyResult struct {
	// Array of reports
	Data []*GetWorksEmbedListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 18
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 2
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s GetWorksEmbedListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetWorksEmbedListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWorksEmbedListResponseBodyResult) GetData() []*GetWorksEmbedListResponseBodyResultData {
	return s.Data
}

func (s *GetWorksEmbedListResponseBodyResult) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetWorksEmbedListResponseBodyResult) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetWorksEmbedListResponseBodyResult) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *GetWorksEmbedListResponseBodyResult) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *GetWorksEmbedListResponseBodyResult) SetData(v []*GetWorksEmbedListResponseBodyResultData) *GetWorksEmbedListResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetWorksEmbedListResponseBodyResult) SetPageNo(v int64) *GetWorksEmbedListResponseBodyResult {
	s.PageNo = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResult) SetPageSize(v int64) *GetWorksEmbedListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResult) SetTotalNum(v int64) *GetWorksEmbedListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResult) SetTotalPages(v int64) *GetWorksEmbedListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorksEmbedListResponseBodyResultData struct {
	// Embed time
	//
	// example:
	//
	// YYYY-mm-DD hh:MM:ss
	EmbedTime *string `json:"EmbedTime,omitempty" xml:"EmbedTime,omitempty"`
	// Report ID
	//
	// example:
	//
	// 897ce25e-****-****-af84-d13c5610****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// Report name
	//
	// example:
	//
	// test
	WorksName *string `json:"WorksName,omitempty" xml:"WorksName,omitempty"`
	// Report type
	//
	// example:
	//
	// page
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// 87c6b145-****-43e1-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorksEmbedListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetWorksEmbedListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetWorksEmbedListResponseBodyResultData) GetEmbedTime() *string {
	return s.EmbedTime
}

func (s *GetWorksEmbedListResponseBodyResultData) GetWorksId() *string {
	return s.WorksId
}

func (s *GetWorksEmbedListResponseBodyResultData) GetWorksName() *string {
	return s.WorksName
}

func (s *GetWorksEmbedListResponseBodyResultData) GetWorksType() *string {
	return s.WorksType
}

func (s *GetWorksEmbedListResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorksEmbedListResponseBodyResultData) SetEmbedTime(v string) *GetWorksEmbedListResponseBodyResultData {
	s.EmbedTime = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResultData) SetWorksId(v string) *GetWorksEmbedListResponseBodyResultData {
	s.WorksId = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResultData) SetWorksName(v string) *GetWorksEmbedListResponseBodyResultData {
	s.WorksName = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResultData) SetWorksType(v string) *GetWorksEmbedListResponseBodyResultData {
	s.WorksType = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResultData) SetWorkspaceId(v string) *GetWorksEmbedListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorksEmbedListResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
