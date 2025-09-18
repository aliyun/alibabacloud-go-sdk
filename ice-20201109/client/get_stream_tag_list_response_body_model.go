// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamTagListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStreamTagListResponseBody
	GetCode() *string
	SetNextToken(v string) *GetStreamTagListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetStreamTagListResponseBody
	GetRequestId() *string
	SetStreamTagList(v []*GetStreamTagListResponseBodyStreamTagList) *GetStreamTagListResponseBody
	GetStreamTagList() []*GetStreamTagListResponseBodyStreamTagList
	SetSuccess(v string) *GetStreamTagListResponseBody
	GetSuccess() *string
	SetTotal(v int64) *GetStreamTagListResponseBody
	GetTotal() *int64
}

type GetStreamTagListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ****73f33c91-d59383e8280b****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StreamTagList []*GetStreamTagListResponseBodyStreamTagList `json:"StreamTagList,omitempty" xml:"StreamTagList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 163
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetStreamTagListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTagListResponseBody) GoString() string {
	return s.String()
}

func (s *GetStreamTagListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStreamTagListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetStreamTagListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStreamTagListResponseBody) GetStreamTagList() []*GetStreamTagListResponseBodyStreamTagList {
	return s.StreamTagList
}

func (s *GetStreamTagListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetStreamTagListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetStreamTagListResponseBody) SetCode(v string) *GetStreamTagListResponseBody {
	s.Code = &v
	return s
}

func (s *GetStreamTagListResponseBody) SetNextToken(v string) *GetStreamTagListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetStreamTagListResponseBody) SetRequestId(v string) *GetStreamTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStreamTagListResponseBody) SetStreamTagList(v []*GetStreamTagListResponseBodyStreamTagList) *GetStreamTagListResponseBody {
	s.StreamTagList = v
	return s
}

func (s *GetStreamTagListResponseBody) SetSuccess(v string) *GetStreamTagListResponseBody {
	s.Success = &v
	return s
}

func (s *GetStreamTagListResponseBody) SetTotal(v int64) *GetStreamTagListResponseBody {
	s.Total = &v
	return s
}

func (s *GetStreamTagListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStreamTagListResponseBodyStreamTagList struct {
	// example:
	//
	// 2025-02-25T02:24:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2025-04-26T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// {"result":"xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetStreamTagListResponseBodyStreamTagList) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTagListResponseBodyStreamTagList) GoString() string {
	return s.String()
}

func (s *GetStreamTagListResponseBodyStreamTagList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStreamTagListResponseBodyStreamTagList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStreamTagListResponseBodyStreamTagList) GetUserData() *string {
	return s.UserData
}

func (s *GetStreamTagListResponseBodyStreamTagList) SetEndTime(v string) *GetStreamTagListResponseBodyStreamTagList {
	s.EndTime = &v
	return s
}

func (s *GetStreamTagListResponseBodyStreamTagList) SetStartTime(v string) *GetStreamTagListResponseBodyStreamTagList {
	s.StartTime = &v
	return s
}

func (s *GetStreamTagListResponseBodyStreamTagList) SetUserData(v string) *GetStreamTagListResponseBodyStreamTagList {
	s.UserData = &v
	return s
}

func (s *GetStreamTagListResponseBodyStreamTagList) Validate() error {
	return dara.Validate(s)
}
