// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppKeyPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppKeyPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeAppKeyPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeAppKeyPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeAppKeyPageResponseBodyResultObject) *DescribeAppKeyPageResponseBody
	GetResultObject() []*DescribeAppKeyPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeAppKeyPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeAppKeyPageResponseBody
	GetTotalPage() *int32
}

type DescribeAppKeyPageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// E01E1B4A-6747-5329-9046-B6D6B2D91349
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of items per page, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeAppKeyPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeAppKeyPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppKeyPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAppKeyPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppKeyPageResponseBody) GetResultObject() []*DescribeAppKeyPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAppKeyPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeAppKeyPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeAppKeyPageResponseBody) SetRequestId(v string) *DescribeAppKeyPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppKeyPageResponseBody) SetCurrentPage(v int32) *DescribeAppKeyPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppKeyPageResponseBody) SetPageSize(v int32) *DescribeAppKeyPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppKeyPageResponseBody) SetResultObject(v []*DescribeAppKeyPageResponseBodyResultObject) *DescribeAppKeyPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAppKeyPageResponseBody) SetTotalItem(v int32) *DescribeAppKeyPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeAppKeyPageResponseBody) SetTotalPage(v int32) *DescribeAppKeyPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppKeyPageResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppKeyPageResponseBodyResultObject struct {
	// Android SDK download URL.
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	AndroidSdkUrl *string `json:"androidSdkUrl,omitempty" xml:"androidSdkUrl,omitempty"`
	// Android SDK version number.
	//
	// example:
	//
	// 1
	AndroidSdkVersion *string `json:"androidSdkVersion,omitempty" xml:"androidSdkVersion,omitempty"`
	// Application appkey.
	//
	// example:
	//
	// sh9a71f07fhs556bd767586307e82795
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// iOS SDK download URL.
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	IosSdkUrl *string `json:"iosSdkUrl,omitempty" xml:"iosSdkUrl,omitempty"`
	// iOS SDK version number.
	//
	// example:
	//
	// 1
	IosSdkVersion *string `json:"iosSdkVersion,omitempty" xml:"iosSdkVersion,omitempty"`
	// Memo.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// 已作废
	SdkItems *string `json:"sdkItems,omitempty" xml:"sdkItems,omitempty"`
	// Whether this appKey is integrated.
	//
	// example:
	//
	// false
	Used *string `json:"used,omitempty" xml:"used,omitempty"`
}

func (s DescribeAppKeyPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetAndroidSdkUrl() *string {
	return s.AndroidSdkUrl
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetAndroidSdkVersion() *string {
	return s.AndroidSdkVersion
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetIosSdkUrl() *string {
	return s.IosSdkUrl
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetIosSdkVersion() *string {
	return s.IosSdkVersion
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetSdkItems() *string {
	return s.SdkItems
}

func (s *DescribeAppKeyPageResponseBodyResultObject) GetUsed() *string {
	return s.Used
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetAndroidSdkUrl(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.AndroidSdkUrl = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetAndroidSdkVersion(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.AndroidSdkVersion = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetAppKey(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.AppKey = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetGmtCreate(v int64) *DescribeAppKeyPageResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetIosSdkUrl(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.IosSdkUrl = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetIosSdkVersion(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.IosSdkVersion = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetMemo(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetSdkItems(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.SdkItems = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) SetUsed(v string) *DescribeAppKeyPageResponseBodyResultObject {
	s.Used = &v
	return s
}

func (s *DescribeAppKeyPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
