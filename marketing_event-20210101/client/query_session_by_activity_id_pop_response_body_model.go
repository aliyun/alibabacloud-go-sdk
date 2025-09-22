// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionByActivityIdPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySessionByActivityIdPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*QuerySessionByActivityIdPopResponseBodyData) *QuerySessionByActivityIdPopResponseBody
	GetData() []*QuerySessionByActivityIdPopResponseBodyData
	SetErrCode(v string) *QuerySessionByActivityIdPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QuerySessionByActivityIdPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QuerySessionByActivityIdPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QuerySessionByActivityIdPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySessionByActivityIdPopResponseBody
	GetSuccess() *bool
}

type QuerySessionByActivityIdPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QuerySessionByActivityIdPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s QuerySessionByActivityIdPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByActivityIdPopResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySessionByActivityIdPopResponseBody) GetData() []*QuerySessionByActivityIdPopResponseBodyData {
	return s.Data
}

func (s *QuerySessionByActivityIdPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QuerySessionByActivityIdPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QuerySessionByActivityIdPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySessionByActivityIdPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySessionByActivityIdPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySessionByActivityIdPopResponseBody) SetAccessDeniedDetail(v string) *QuerySessionByActivityIdPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetData(v []*QuerySessionByActivityIdPopResponseBodyData) *QuerySessionByActivityIdPopResponseBody {
	s.Data = v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetErrCode(v string) *QuerySessionByActivityIdPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetErrMessage(v string) *QuerySessionByActivityIdPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetHttpStatusCode(v int32) *QuerySessionByActivityIdPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetRequestId(v string) *QuerySessionByActivityIdPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetSuccess(v bool) *QuerySessionByActivityIdPopResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySessionByActivityIdPopResponseBodyData struct {
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// description
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	EndDateTime *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 地点
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 1234
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// name
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	StartDateTime *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
}

func (s QuerySessionByActivityIdPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByActivityIdPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetEndDateTime() *string {
	return s.EndDateTime
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetNameEn() *string {
	return s.NameEn
}

func (s *QuerySessionByActivityIdPopResponseBodyData) GetStartDateTime() *string {
	return s.StartDateTime
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetDescription(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.Description = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetDescriptionEn(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetEndDateTime(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.EndDateTime = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetId(v int64) *QuerySessionByActivityIdPopResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetLocation(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.Location = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetName(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetNameEn(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.NameEn = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetStartDateTime(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.StartDateTime = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
