// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTemplateMaterialResponseBody
	GetCode() *string
	SetData(v []*ListTemplateMaterialResponseBodyData) *ListTemplateMaterialResponseBody
	GetData() []*ListTemplateMaterialResponseBodyData
	SetHttpStatusCode(v int32) *ListTemplateMaterialResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListTemplateMaterialResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTemplateMaterialResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTemplateMaterialResponseBody
	GetNextToken() *string
	SetPage(v int64) *ListTemplateMaterialResponseBody
	GetPage() *int64
	SetRequestId(v string) *ListTemplateMaterialResponseBody
	GetRequestId() *string
	SetSize(v int64) *ListTemplateMaterialResponseBody
	GetSize() *int64
	SetSuccess(v bool) *ListTemplateMaterialResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListTemplateMaterialResponseBody
	GetTotal() *int64
	SetTotalCount(v int32) *ListTemplateMaterialResponseBody
	GetTotalCount() *int32
}

type ListTemplateMaterialResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListTemplateMaterialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// M1Ti7gfj7VGDQgD588hxReiw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTemplateMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateMaterialResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTemplateMaterialResponseBody) GetData() []*ListTemplateMaterialResponseBodyData {
	return s.Data
}

func (s *ListTemplateMaterialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTemplateMaterialResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateMaterialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTemplateMaterialResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateMaterialResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListTemplateMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateMaterialResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *ListTemplateMaterialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTemplateMaterialResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListTemplateMaterialResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplateMaterialResponseBody) SetCode(v string) *ListTemplateMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetData(v []*ListTemplateMaterialResponseBodyData) *ListTemplateMaterialResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetHttpStatusCode(v int32) *ListTemplateMaterialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetMaxResults(v int32) *ListTemplateMaterialResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetMessage(v string) *ListTemplateMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetNextToken(v string) *ListTemplateMaterialResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetPage(v int64) *ListTemplateMaterialResponseBody {
	s.Page = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetRequestId(v string) *ListTemplateMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetSize(v int64) *ListTemplateMaterialResponseBody {
	s.Size = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetSuccess(v bool) *ListTemplateMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetTotal(v int64) *ListTemplateMaterialResponseBody {
	s.Total = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) SetTotalCount(v int32) *ListTemplateMaterialResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplateMaterialResponseBody) Validate() error {
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

type ListTemplateMaterialResponseBodyData struct {
	// example:
	//
	// BROADCAST
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// M1Ti7gfj7VGDQgD588hxReiw
	TemplateId   *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// //daily-avatar-property.oss-cn-beijing.aliyuncs.com/material/INPUT_TRAIN_TEMPLATE_UGC_BODY/Mt.CNYEA3CV25QU2/%E5%BA%95%E6%9D%BF3%E6%9B%BF%E6%8D%A2.mp4?Expires=1764326111&OSSAccessKeyId=LTAI5tBRPnF5JkRCidYA8kw9&Signature=3WU2%2FV2XDaQdt00lkAInK6yr9uk%3D
	TemplateURL *string `json:"templateURL,omitempty" xml:"templateURL,omitempty"`
}

func (s ListTemplateMaterialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateMaterialResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTemplateMaterialResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *ListTemplateMaterialResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplateMaterialResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplateMaterialResponseBodyData) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *ListTemplateMaterialResponseBodyData) SetBizType(v string) *ListTemplateMaterialResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListTemplateMaterialResponseBodyData) SetTemplateId(v string) *ListTemplateMaterialResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListTemplateMaterialResponseBodyData) SetTemplateName(v string) *ListTemplateMaterialResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateMaterialResponseBodyData) SetTemplateURL(v string) *ListTemplateMaterialResponseBodyData {
	s.TemplateURL = &v
	return s
}

func (s *ListTemplateMaterialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
