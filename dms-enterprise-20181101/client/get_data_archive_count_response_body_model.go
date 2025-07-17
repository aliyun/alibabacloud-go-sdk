// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataArchiveCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataArchiveCountResponseBodyData) *GetDataArchiveCountResponseBody
	GetData() *GetDataArchiveCountResponseBodyData
	SetRequestId(v string) *GetDataArchiveCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataArchiveCountResponseBody
	GetSuccess() *bool
}

type GetDataArchiveCountResponseBody struct {
	// The data returned.
	Data *GetDataArchiveCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B63CAC5-BD7F-5C7C-82C9-59DFFBC3C5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataArchiveCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataArchiveCountResponseBody) GetData() *GetDataArchiveCountResponseBodyData {
	return s.Data
}

func (s *GetDataArchiveCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataArchiveCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataArchiveCountResponseBody) SetData(v *GetDataArchiveCountResponseBodyData) *GetDataArchiveCountResponseBody {
	s.Data = v
	return s
}

func (s *GetDataArchiveCountResponseBody) SetRequestId(v string) *GetDataArchiveCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataArchiveCountResponseBody) SetSuccess(v bool) *GetDataArchiveCountResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataArchiveCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataArchiveCountResponseBodyData struct {
	// The number of tickets that data archiving failed.
	//
	// example:
	//
	// 1**
	FailCount *int64 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The number of tickets that data archiving is in progress.
	//
	// example:
	//
	// 2**
	ProcessingCount *int64 `json:"ProcessingCount,omitempty" xml:"ProcessingCount,omitempty"`
	// The number of tickets that data archiving is successful.
	//
	// example:
	//
	// 3**
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of data archiving tickets.
	//
	// example:
	//
	// 6**
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDataArchiveCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataArchiveCountResponseBodyData) GetFailCount() *int64 {
	return s.FailCount
}

func (s *GetDataArchiveCountResponseBodyData) GetProcessingCount() *int64 {
	return s.ProcessingCount
}

func (s *GetDataArchiveCountResponseBodyData) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *GetDataArchiveCountResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDataArchiveCountResponseBodyData) SetFailCount(v int64) *GetDataArchiveCountResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *GetDataArchiveCountResponseBodyData) SetProcessingCount(v int64) *GetDataArchiveCountResponseBodyData {
	s.ProcessingCount = &v
	return s
}

func (s *GetDataArchiveCountResponseBodyData) SetSuccessCount(v int64) *GetDataArchiveCountResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *GetDataArchiveCountResponseBodyData) SetTotalCount(v int64) *GetDataArchiveCountResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetDataArchiveCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
