// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionListPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySessionListPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*QuerySessionListPopResponseBodyData) *QuerySessionListPopResponseBody
	GetData() []*QuerySessionListPopResponseBodyData
	SetErrCode(v string) *QuerySessionListPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QuerySessionListPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QuerySessionListPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QuerySessionListPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySessionListPopResponseBody
	GetSuccess() *bool
}

type QuerySessionListPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QuerySessionListPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySessionListPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionListPopResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySessionListPopResponseBody) GetData() []*QuerySessionListPopResponseBodyData {
	return s.Data
}

func (s *QuerySessionListPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QuerySessionListPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QuerySessionListPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySessionListPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySessionListPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySessionListPopResponseBody) SetAccessDeniedDetail(v string) *QuerySessionListPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetData(v []*QuerySessionListPopResponseBodyData) *QuerySessionListPopResponseBody {
	s.Data = v
	return s
}

func (s *QuerySessionListPopResponseBody) SetErrCode(v string) *QuerySessionListPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetErrMessage(v string) *QuerySessionListPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetHttpStatusCode(v int32) *QuerySessionListPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetRequestId(v string) *QuerySessionListPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetSuccess(v bool) *QuerySessionListPopResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySessionListPopResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySessionListPopResponseBodyData struct {
	// code
	//
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// id
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// location
	//
	// example:
	//
	// location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// name
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QuerySessionListPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionListPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *QuerySessionListPopResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySessionListPopResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QuerySessionListPopResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *QuerySessionListPopResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QuerySessionListPopResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySessionListPopResponseBodyData) SetCode(v string) *QuerySessionListPopResponseBodyData {
	s.Code = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetEndTime(v string) *QuerySessionListPopResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetId(v int64) *QuerySessionListPopResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetLocation(v string) *QuerySessionListPopResponseBodyData {
	s.Location = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetName(v string) *QuerySessionListPopResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetStartTime(v string) *QuerySessionListPopResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
