// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPictureSearchDevicesResponseBody
	GetCode() *string
	SetData(v *QueryPictureSearchDevicesResponseBodyData) *QueryPictureSearchDevicesResponseBody
	GetData() *QueryPictureSearchDevicesResponseBodyData
	SetErrorMessage(v string) *QueryPictureSearchDevicesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryPictureSearchDevicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPictureSearchDevicesResponseBody
	GetSuccess() *bool
}

type QueryPictureSearchDevicesResponseBody struct {
	// The error code returned if the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryPictureSearchDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Specified parameter CurrentPage is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DF7DC571-FE48-4C25-8BE0-B13C3038C85B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPictureSearchDevicesResponseBody) GetData() *QueryPictureSearchDevicesResponseBodyData {
	return s.Data
}

func (s *QueryPictureSearchDevicesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryPictureSearchDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPictureSearchDevicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPictureSearchDevicesResponseBody) SetCode(v string) *QueryPictureSearchDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetData(v *QueryPictureSearchDevicesResponseBodyData) *QueryPictureSearchDevicesResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetErrorMessage(v string) *QueryPictureSearchDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetRequestId(v string) *QueryPictureSearchDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetSuccess(v bool) *QueryPictureSearchDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPictureSearchDevicesResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The list of applications.
	PageData []*QueryPictureSearchDevicesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchDevicesResponseBodyData) GetPageCount() *int32 {
	return s.PageCount
}

func (s *QueryPictureSearchDevicesResponseBodyData) GetPageData() []*QueryPictureSearchDevicesResponseBodyDataPageData {
	return s.PageData
}

func (s *QueryPictureSearchDevicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchDevicesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageCount(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageData(v []*QueryPictureSearchDevicesResponseBodyDataPageData) *QueryPictureSearchDevicesResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageSize(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetTotal(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPictureSearchDevicesResponseBodyDataPageData struct {
	// The ID of the device.
	//
	// example:
	//
	// VrJwPGjC1sJlqPjZA3cxg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The alias of the device.
	//
	// > This parameter is returned only for IPC devices.
	//
	// example:
	//
	// camera1
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) GetIotId() *string {
	return s.IotId
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) GetNickName() *string {
	return s.NickName
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchDevicesResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) SetNickName(v string) *QueryPictureSearchDevicesResponseBodyDataPageData {
	s.NickName = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
