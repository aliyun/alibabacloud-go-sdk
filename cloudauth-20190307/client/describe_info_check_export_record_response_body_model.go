// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInfoCheckExportRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInfoCheckExportRecordResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *DescribeInfoCheckExportRecordResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeInfoCheckExportRecordResponseBodyItems) *DescribeInfoCheckExportRecordResponseBody
	GetItems() []*DescribeInfoCheckExportRecordResponseBodyItems
	SetMessage(v string) *DescribeInfoCheckExportRecordResponseBody
	GetMessage() *string
	SetPageSize(v int32) *DescribeInfoCheckExportRecordResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInfoCheckExportRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInfoCheckExportRecordResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeInfoCheckExportRecordResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *DescribeInfoCheckExportRecordResponseBody
	GetTotalPage() *int32
}

type DescribeInfoCheckExportRecordResponseBody struct {
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of results.
	Items []*DescribeInfoCheckExportRecordResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeInfoCheckExportRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInfoCheckExportRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetItems() []*DescribeInfoCheckExportRecordResponseBodyItems {
	return s.Items
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInfoCheckExportRecordResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetCode(v string) *DescribeInfoCheckExportRecordResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetCurrentPage(v int32) *DescribeInfoCheckExportRecordResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetItems(v []*DescribeInfoCheckExportRecordResponseBodyItems) *DescribeInfoCheckExportRecordResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetMessage(v string) *DescribeInfoCheckExportRecordResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetPageSize(v int32) *DescribeInfoCheckExportRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetRequestId(v string) *DescribeInfoCheckExportRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetSuccess(v bool) *DescribeInfoCheckExportRecordResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetTotalCount(v int32) *DescribeInfoCheckExportRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) SetTotalPage(v int32) *DescribeInfoCheckExportRecordResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInfoCheckExportRecordResponseBodyItems struct {
	// The download date.
	//
	// example:
	//
	// 1758250979000
	DownloadDate *string `json:"DownloadDate,omitempty" xml:"DownloadDate,omitempty"`
	// The ID of the download task.
	//
	// example:
	//
	// 202510189017278
	DownloadTaskId *string `json:"DownloadTaskId,omitempty" xml:"DownloadTaskId,omitempty"`
	// The error code.
	//
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The file name.
	//
	// example:
	//
	// 手机号二要素统计202509013975081.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type.
	//
	// example:
	//
	// CSV
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The product type name. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification.
	//
	// - **ID_PERIOD**: ID card validity period verification.
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration.
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status.
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition).
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition).
	//
	// - **MOBILE_2_META**: mobile number two-element verification.
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition).
	//
	// - **MOBILE_DETECT**: phone number detection.
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition).
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition.
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition.
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query.
	//
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// example:
	//
	// 身份证二要素
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The task status. Valid values:
	//
	// - **1**: The file is being generated.
	//
	// - **2**: The file has been generated.
	//
	// - **3**: The file failed to be generated.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task URL.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth.oss-cn-shanghai.aliyuncs.com/console/xxxxxxxx.xlsx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInfoCheckExportRecordResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInfoCheckExportRecordResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetDownloadDate() *string {
	return s.DownloadDate
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetDownloadTaskId() *string {
	return s.DownloadTaskId
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetFileName() *string {
	return s.FileName
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetFileType() *string {
	return s.FileType
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetDownloadDate(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.DownloadDate = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetDownloadTaskId(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.DownloadTaskId = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetErrorCode(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetFileName(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.FileName = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetFileType(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.FileType = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetProductType(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.ProductType = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetStatus(v int32) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) SetUrl(v string) *DescribeInfoCheckExportRecordResponseBodyItems {
	s.Url = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
