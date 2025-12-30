// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateTTSVoicesCustomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrivateTTSVoicesCustomResponseBody
	GetCode() *string
	SetData(v *ListPrivateTTSVoicesCustomResponseBodyData) *ListPrivateTTSVoicesCustomResponseBody
	GetData() *ListPrivateTTSVoicesCustomResponseBodyData
	SetHttpStatusCode(v int32) *ListPrivateTTSVoicesCustomResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListPrivateTTSVoicesCustomResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPrivateTTSVoicesCustomResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPrivateTTSVoicesCustomResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPrivateTTSVoicesCustomResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPrivateTTSVoicesCustomResponseBody
	GetSuccess() *bool
}

type ListPrivateTTSVoicesCustomResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListPrivateTTSVoicesCustomResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// qht-fc2-test
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPrivateTTSVoicesCustomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateTTSVoicesCustomResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetData() *ListPrivateTTSVoicesCustomResponseBodyData {
	return s.Data
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateTTSVoicesCustomResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetCode(v string) *ListPrivateTTSVoicesCustomResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetData(v *ListPrivateTTSVoicesCustomResponseBodyData) *ListPrivateTTSVoicesCustomResponseBody {
	s.Data = v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetHttpStatusCode(v int32) *ListPrivateTTSVoicesCustomResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetMaxResults(v int32) *ListPrivateTTSVoicesCustomResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetMessage(v string) *ListPrivateTTSVoicesCustomResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetNextToken(v string) *ListPrivateTTSVoicesCustomResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetRequestId(v string) *ListPrivateTTSVoicesCustomResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) SetSuccess(v bool) *ListPrivateTTSVoicesCustomResponseBody {
	s.Success = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPrivateTTSVoicesCustomResponseBodyData struct {
	Data []*ListPrivateTTSVoicesCustomResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 16
	Pages *int32 `json:"pages,omitempty" xml:"pages,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 20
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPrivateTTSVoicesCustomResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateTTSVoicesCustomResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) GetData() []*ListPrivateTTSVoicesCustomResponseBodyDataData {
	return s.Data
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) GetPages() *int32 {
	return s.Pages
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) SetData(v []*ListPrivateTTSVoicesCustomResponseBodyDataData) *ListPrivateTTSVoicesCustomResponseBodyData {
	s.Data = v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) SetPage(v int32) *ListPrivateTTSVoicesCustomResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) SetPages(v int32) *ListPrivateTTSVoicesCustomResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) SetSize(v int32) *ListPrivateTTSVoicesCustomResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) SetTotal(v int64) *ListPrivateTTSVoicesCustomResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyData) Validate() error {
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

type ListPrivateTTSVoicesCustomResponseBodyDataData struct {
	// example:
	//
	// https://xxx-aliyuncs.com/material/INPUT_TRAIN_AUDIO/Mt.CQEG75L4FWIU2/TestTTSVoiceName.mp3?Expires=1764262805&OSSAccessKeyId=LTAI5tK3WcKwKtAyaTSeq7sx&Signature=D%2Fld6gp9Zh6TsGRU9cd6GD2pFY0%3D
	AudioURL *string `json:"audioURL,omitempty" xml:"audioURL,omitempty"`
	// example:
	//
	// PASS
	CensorStatus *string `json:"censorStatus,omitempty" xml:"censorStatus,omitempty"`
	// example:
	//
	// false
	Common *bool `json:"common,omitempty" xml:"common,omitempty"`
	// example:
	//
	// 2024-04-15T09:42:15Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// Optional for WH managernif no issue will be cancel
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ErrorCode   *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Error
	ErrorDetail *string `json:"errorDetail,omitempty" xml:"errorDetail,omitempty"`
	// example:
	//
	// FEMALE
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// M1lhKArheOyYdeYybDFqS1-Q
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ZH_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2024-04-15T09:42:15Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// TestTTSVoice
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 100
	RemainSeconds *int64 `json:"remainSeconds,omitempty" xml:"remainSeconds,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Text   *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// TestTTSVoice
	VoiceKey *string `json:"voiceKey,omitempty" xml:"voiceKey,omitempty"`
}

func (s ListPrivateTTSVoicesCustomResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateTTSVoicesCustomResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetAudioURL() *string {
	return s.AudioURL
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetCensorStatus() *string {
	return s.CensorStatus
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetCommon() *bool {
	return s.Common
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetDescription() *string {
	return s.Description
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetGender() *string {
	return s.Gender
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetId() *string {
	return s.Id
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetLanguage() *string {
	return s.Language
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetName() *string {
	return s.Name
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetRemainSeconds() *int64 {
	return s.RemainSeconds
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetText() *string {
	return s.Text
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) GetVoiceKey() *string {
	return s.VoiceKey
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetAudioURL(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.AudioURL = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetCensorStatus(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.CensorStatus = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetCommon(v bool) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Common = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetCreateTime(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetDescription(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Description = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetErrorCode(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.ErrorCode = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetErrorDetail(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.ErrorDetail = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetGender(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Gender = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetId(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetLanguage(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Language = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetModifiedTime(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.ModifiedTime = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetName(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetRemainSeconds(v int64) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.RemainSeconds = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetStatus(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetText(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.Text = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) SetVoiceKey(v string) *ListPrivateTTSVoicesCustomResponseBodyDataData {
	s.VoiceKey = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
