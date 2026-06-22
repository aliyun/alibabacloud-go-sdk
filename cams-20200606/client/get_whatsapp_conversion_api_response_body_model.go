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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetWhatsappConversionApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 示例值
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// 示例值
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
