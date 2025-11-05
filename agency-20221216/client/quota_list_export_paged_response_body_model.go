// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaListExportPagedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuotaListExportPagedResponseBody
	GetCode() *string
	SetData(v []*QuotaListExportPagedResponseBodyData) *QuotaListExportPagedResponseBody
	GetData() []*QuotaListExportPagedResponseBodyData
	SetMsg(v string) *QuotaListExportPagedResponseBody
	GetMsg() *string
	SetPageNo(v int32) *QuotaListExportPagedResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *QuotaListExportPagedResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QuotaListExportPagedResponseBody
	GetRequestId() *string
	SetTotal(v int32) *QuotaListExportPagedResponseBody
	GetTotal() *int32
}

type QuotaListExportPagedResponseBody struct {
	// Status code of returning result, 200 means success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Listed data of returning result
	Data []*QuotaListExportPagedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of returning result
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Record number on each page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of the Request
	//
	// example:
	//
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total volume
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuotaListExportPagedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuotaListExportPagedResponseBody) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuotaListExportPagedResponseBody) GetData() []*QuotaListExportPagedResponseBodyData {
	return s.Data
}

func (s *QuotaListExportPagedResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *QuotaListExportPagedResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QuotaListExportPagedResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuotaListExportPagedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuotaListExportPagedResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *QuotaListExportPagedResponseBody) SetCode(v string) *QuotaListExportPagedResponseBody {
	s.Code = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetData(v []*QuotaListExportPagedResponseBodyData) *QuotaListExportPagedResponseBody {
	s.Data = v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetMsg(v string) *QuotaListExportPagedResponseBody {
	s.Msg = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetPageNo(v int32) *QuotaListExportPagedResponseBody {
	s.PageNo = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetPageSize(v int32) *QuotaListExportPagedResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetRequestId(v string) *QuotaListExportPagedResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetTotal(v int32) *QuotaListExportPagedResponseBody {
	s.Total = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) Validate() error {
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

type QuotaListExportPagedResponseBodyData struct {
	// Create Time
	//
	// example:
	//
	// 2023-12-21 21:31:57 UTC+8
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// File Name
	//
	// example:
	//
	// 5113766248601929_quota_2023-06-22_2023-12-21_all_2023122121310057
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Notification Message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Display of Task Status
	//
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task Status Enum</br>
	//
	// 2: Exporting</br>
	//
	// 3: Export Success</br>
	//
	// -1: Export Fail</br>
	//
	// example:
	//
	// Export Success
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The link to download exported file.
	//
	// example:
	//
	// //aliyun-eco-market-servic-singapore.oss-ap-southeast-1.aliyuncs.com/5113766248601929_quota_2023-06-22_2023-12-21_all_2023122121310057
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QuotaListExportPagedResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuotaListExportPagedResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QuotaListExportPagedResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *QuotaListExportPagedResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *QuotaListExportPagedResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QuotaListExportPagedResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *QuotaListExportPagedResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *QuotaListExportPagedResponseBodyData) SetCreateTime(v string) *QuotaListExportPagedResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetFileName(v string) *QuotaListExportPagedResponseBodyData {
	s.FileName = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetMessage(v string) *QuotaListExportPagedResponseBodyData {
	s.Message = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetStatus(v string) *QuotaListExportPagedResponseBodyData {
	s.Status = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetStatusCode(v string) *QuotaListExportPagedResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetUrl(v string) *QuotaListExportPagedResponseBodyData {
	s.Url = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) Validate() error {
	return dara.Validate(s)
}
