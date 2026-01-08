// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpecReviewTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSpecReviewTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListSpecReviewTasksResponseBodyResult) *ListSpecReviewTasksResponseBody
	GetResult() []*ListSpecReviewTasksResponseBodyResult
	SetTotalCount(v int32) *ListSpecReviewTasksResponseBody
	GetTotalCount() *int32
}

type ListSpecReviewTasksResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 55F7B3FE-05D8-5F0F-BD55-A18967D447DC
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSpecReviewTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSpecReviewTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSpecReviewTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSpecReviewTasksResponseBody) GetResult() []*ListSpecReviewTasksResponseBodyResult {
	return s.Result
}

func (s *ListSpecReviewTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSpecReviewTasksResponseBody) SetRequestId(v string) *ListSpecReviewTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpecReviewTasksResponseBody) SetResult(v []*ListSpecReviewTasksResponseBodyResult) *ListSpecReviewTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListSpecReviewTasksResponseBody) SetTotalCount(v int32) *ListSpecReviewTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSpecReviewTasksResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSpecReviewTasksResponseBodyResult struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 339
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	AppName     *string `json:"appName,omitempty" xml:"appName,omitempty"`
	ApplyReason *string `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// example:
	//
	// 2024-05-27T10:13:22.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// QUOTA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSpecReviewTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSpecReviewTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListSpecReviewTasksResponseBodyResult) GetAppName() *string {
	return s.AppName
}

func (s *ListSpecReviewTasksResponseBodyResult) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *ListSpecReviewTasksResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListSpecReviewTasksResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListSpecReviewTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListSpecReviewTasksResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListSpecReviewTasksResponseBodyResult) SetId(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetAppName(v string) *ListSpecReviewTasksResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetApplyReason(v string) *ListSpecReviewTasksResponseBodyResult {
	s.ApplyReason = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetGmtCreate(v string) *ListSpecReviewTasksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetSource(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetStatus(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetType(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
