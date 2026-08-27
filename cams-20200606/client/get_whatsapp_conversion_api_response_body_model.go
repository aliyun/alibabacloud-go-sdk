// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappConversionApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetWhatsappConversionApiResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetWhatsappConversionApiResponseBody
	GetCode() *string
	SetData(v []*GetWhatsappConversionApiResponseBodyData) *GetWhatsappConversionApiResponseBody
	GetData() []*GetWhatsappConversionApiResponseBodyData
	SetMessage(v string) *GetWhatsappConversionApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWhatsappConversionApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWhatsappConversionApiResponseBody
	GetSuccess() *bool
}

type GetWhatsappConversionApiResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetWhatsappConversionApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each API request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWhatsappConversionApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConversionApiResponseBody) GoString() string {
	return s.String()
}

func (s *GetWhatsappConversionApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetWhatsappConversionApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWhatsappConversionApiResponseBody) GetData() []*GetWhatsappConversionApiResponseBodyData {
	return s.Data
}

func (s *GetWhatsappConversionApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWhatsappConversionApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWhatsappConversionApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWhatsappConversionApiResponseBody) SetAccessDeniedDetail(v string) *GetWhatsappConversionApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBody) SetCode(v string) *GetWhatsappConversionApiResponseBody {
	s.Code = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBody) SetData(v []*GetWhatsappConversionApiResponseBodyData) *GetWhatsappConversionApiResponseBody {
	s.Data = v
	return s
}

func (s *GetWhatsappConversionApiResponseBody) SetMessage(v string) *GetWhatsappConversionApiResponseBody {
	s.Message = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBody) SetRequestId(v string) *GetWhatsappConversionApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBody) SetSuccess(v bool) *GetWhatsappConversionApiResponseBody {
	s.Success = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBody) Validate() error {
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

type GetWhatsappConversionApiResponseBodyData struct {
	// The dataset ID.
	//
	// example:
	//
	// 111
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The PageId of Messenger.
	//
	// example:
	//
	// 7832312
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The page name.
	//
	// example:
	//
	// aaa
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
}

func (s GetWhatsappConversionApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConversionApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWhatsappConversionApiResponseBodyData) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetWhatsappConversionApiResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *GetWhatsappConversionApiResponseBodyData) GetPageName() *string {
	return s.PageName
}

func (s *GetWhatsappConversionApiResponseBodyData) SetDatasetId(v string) *GetWhatsappConversionApiResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBodyData) SetPageId(v string) *GetWhatsappConversionApiResponseBodyData {
	s.PageId = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBodyData) SetPageName(v string) *GetWhatsappConversionApiResponseBodyData {
	s.PageName = &v
	return s
}

func (s *GetWhatsappConversionApiResponseBodyData) Validate() error {
	return dara.Validate(s)
}
