// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAudienceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListCustomAudienceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListCustomAudienceResponseBody
	GetCode() *string
	SetData(v []*ListCustomAudienceResponseBodyData) *ListCustomAudienceResponseBody
	GetData() []*ListCustomAudienceResponseBodyData
	SetMessage(v string) *ListCustomAudienceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCustomAudienceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomAudienceResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListCustomAudienceResponseBody
	GetTotal() *int64
}

type ListCustomAudienceResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - A value of OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListCustomAudienceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful.
	//
	// - **true**: successful.
	//
	// - **false**: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 69
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCustomAudienceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAudienceResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAudienceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListCustomAudienceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomAudienceResponseBody) GetData() []*ListCustomAudienceResponseBodyData {
	return s.Data
}

func (s *ListCustomAudienceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomAudienceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAudienceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomAudienceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCustomAudienceResponseBody) SetAccessDeniedDetail(v string) *ListCustomAudienceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListCustomAudienceResponseBody) SetCode(v string) *ListCustomAudienceResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomAudienceResponseBody) SetData(v []*ListCustomAudienceResponseBodyData) *ListCustomAudienceResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomAudienceResponseBody) SetMessage(v string) *ListCustomAudienceResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomAudienceResponseBody) SetRequestId(v string) *ListCustomAudienceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAudienceResponseBody) SetSuccess(v bool) *ListCustomAudienceResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomAudienceResponseBody) SetTotal(v int64) *ListCustomAudienceResponseBody {
	s.Total = &v
	return s
}

func (s *ListCustomAudienceResponseBody) Validate() error {
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

type ListCustomAudienceResponseBodyData struct {
	// The Meta ad account ID.
	//
	// example:
	//
	// 339**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// The time when the audience was created.
	//
	// example:
	//
	// 1720356898
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the custom audience.
	//
	// example:
	//
	// 339**
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// The name of the custom audience.
	//
	// example:
	//
	// custom audience name
	CustomAudienceName *string `json:"CustomAudienceName,omitempty" xml:"CustomAudienceName,omitempty"`
	// The description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Page ID for Messenger.
	//
	// example:
	//
	// 239**
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The status.
	//
	// example:
	//
	// None
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of tokens.
	//
	// example:
	//
	// 70
	TokenTotal *int64 `json:"TokenTotal,omitempty" xml:"TokenTotal,omitempty"`
	// The token type.
	//
	// example:
	//
	// custom
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
	// The time when the audience was last updated.
	//
	// example:
	//
	// 51
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The upload type.
	//
	// example:
	//
	// excel
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s ListCustomAudienceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAudienceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomAudienceResponseBodyData) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListCustomAudienceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCustomAudienceResponseBodyData) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *ListCustomAudienceResponseBodyData) GetCustomAudienceName() *string {
	return s.CustomAudienceName
}

func (s *ListCustomAudienceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListCustomAudienceResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *ListCustomAudienceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListCustomAudienceResponseBodyData) GetTokenTotal() *int64 {
	return s.TokenTotal
}

func (s *ListCustomAudienceResponseBodyData) GetTokenType() *string {
	return s.TokenType
}

func (s *ListCustomAudienceResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCustomAudienceResponseBodyData) GetUploadType() *string {
	return s.UploadType
}

func (s *ListCustomAudienceResponseBodyData) SetAdAccountId(v string) *ListCustomAudienceResponseBodyData {
	s.AdAccountId = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetCreateTime(v int64) *ListCustomAudienceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetCustomAudienceId(v string) *ListCustomAudienceResponseBodyData {
	s.CustomAudienceId = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetCustomAudienceName(v string) *ListCustomAudienceResponseBodyData {
	s.CustomAudienceName = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetDescription(v string) *ListCustomAudienceResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetPageId(v string) *ListCustomAudienceResponseBodyData {
	s.PageId = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetStatus(v string) *ListCustomAudienceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetTokenTotal(v int64) *ListCustomAudienceResponseBodyData {
	s.TokenTotal = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetTokenType(v string) *ListCustomAudienceResponseBodyData {
	s.TokenType = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetUpdateTime(v int64) *ListCustomAudienceResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) SetUploadType(v string) *ListCustomAudienceResponseBodyData {
	s.UploadType = &v
	return s
}

func (s *ListCustomAudienceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
