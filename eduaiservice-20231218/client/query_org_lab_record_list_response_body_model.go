// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgLabRecordListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryOrgLabRecordListResponseBodyData) *QueryOrgLabRecordListResponseBody
	GetData() []*QueryOrgLabRecordListResponseBodyData
	SetErrCode(v string) *QueryOrgLabRecordListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryOrgLabRecordListResponseBody
	GetErrMessage() *string
	SetPageIndex(v int64) *QueryOrgLabRecordListResponseBody
	GetPageIndex() *int64
	SetPageSize(v int64) *QueryOrgLabRecordListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *QueryOrgLabRecordListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryOrgLabRecordListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *QueryOrgLabRecordListResponseBody
	GetTotalCount() *int64
}

type QueryOrgLabRecordListResponseBody struct {
	Data       []*QueryOrgLabRecordListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode    *string                                  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                                  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageIndex  *int64                                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int64                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryOrgLabRecordListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgLabRecordListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgLabRecordListResponseBody) GetData() []*QueryOrgLabRecordListResponseBodyData {
	return s.Data
}

func (s *QueryOrgLabRecordListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryOrgLabRecordListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryOrgLabRecordListResponseBody) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *QueryOrgLabRecordListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryOrgLabRecordListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrgLabRecordListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOrgLabRecordListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryOrgLabRecordListResponseBody) SetData(v []*QueryOrgLabRecordListResponseBodyData) *QueryOrgLabRecordListResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetErrCode(v string) *QueryOrgLabRecordListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetErrMessage(v string) *QueryOrgLabRecordListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetPageIndex(v int64) *QueryOrgLabRecordListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetPageSize(v int64) *QueryOrgLabRecordListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetRequestId(v string) *QueryOrgLabRecordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetSuccess(v bool) *QueryOrgLabRecordListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) SetTotalCount(v int64) *QueryOrgLabRecordListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBody) Validate() error {
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

type QueryOrgLabRecordListResponseBodyData struct {
	CreatedAt   *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	LabRecordId *string `json:"LabRecordId,omitempty" xml:"LabRecordId,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmittedAt *int64  `json:"SubmittedAt,omitempty" xml:"SubmittedAt,omitempty"`
}

func (s QueryOrgLabRecordListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgLabRecordListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrgLabRecordListResponseBodyData) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *QueryOrgLabRecordListResponseBodyData) GetLabRecordId() *string {
	return s.LabRecordId
}

func (s *QueryOrgLabRecordListResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryOrgLabRecordListResponseBodyData) GetSubmittedAt() *int64 {
	return s.SubmittedAt
}

func (s *QueryOrgLabRecordListResponseBodyData) SetCreatedAt(v int64) *QueryOrgLabRecordListResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBodyData) SetLabRecordId(v string) *QueryOrgLabRecordListResponseBodyData {
	s.LabRecordId = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBodyData) SetStatus(v int32) *QueryOrgLabRecordListResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBodyData) SetSubmittedAt(v int64) *QueryOrgLabRecordListResponseBodyData {
	s.SubmittedAt = &v
	return s
}

func (s *QueryOrgLabRecordListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
