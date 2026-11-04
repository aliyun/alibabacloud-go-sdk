// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchAiboxesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPictureSearchAiboxesResponseBody
	GetCode() *string
	SetData(v *QueryPictureSearchAiboxesResponseBodyData) *QueryPictureSearchAiboxesResponseBody
	GetData() *QueryPictureSearchAiboxesResponseBodyData
	SetErrorMessage(v string) *QueryPictureSearchAiboxesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryPictureSearchAiboxesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPictureSearchAiboxesResponseBody
	GetSuccess() *bool
}

type QueryPictureSearchAiboxesResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: an error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/288314.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned when the call is successful.
	Data *QueryPictureSearchAiboxesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s QueryPictureSearchAiboxesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPictureSearchAiboxesResponseBody) GetData() *QueryPictureSearchAiboxesResponseBodyData {
	return s.Data
}

func (s *QueryPictureSearchAiboxesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryPictureSearchAiboxesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPictureSearchAiboxesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPictureSearchAiboxesResponseBody) SetCode(v string) *QueryPictureSearchAiboxesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetData(v *QueryPictureSearchAiboxesResponseBodyData) *QueryPictureSearchAiboxesResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetErrorMessage(v string) *QueryPictureSearchAiboxesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetRequestId(v string) *QueryPictureSearchAiboxesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetSuccess(v bool) *QueryPictureSearchAiboxesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPictureSearchAiboxesResponseBodyData struct {
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
	// The list of V-series edge all-in-one machines.
	PageData []*QueryPictureSearchAiboxesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
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

func (s QueryPictureSearchAiboxesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchAiboxesResponseBodyData) GetPageCount() *int32 {
	return s.PageCount
}

func (s *QueryPictureSearchAiboxesResponseBodyData) GetPageData() []*QueryPictureSearchAiboxesResponseBodyDataPageData {
	return s.PageData
}

func (s *QueryPictureSearchAiboxesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchAiboxesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageCount(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageData(v []*QueryPictureSearchAiboxesResponseBodyDataPageData) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageSize(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetTotal(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) Validate() error {
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

type QueryPictureSearchAiboxesResponseBodyDataPageData struct {
	// The device ID of the V-series edge all-in-one machine.
	//
	// example:
	//
	// VrJwPGjC1sJlqPjZA3cxg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The device alias of the V-series edge all-in-one machine.
	//
	// example:
	//
	// box015
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s QueryPictureSearchAiboxesResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) GetIotId() *string {
	return s.IotId
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) GetNickName() *string {
	return s.NickName
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchAiboxesResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) SetNickName(v string) *QueryPictureSearchAiboxesResponseBodyDataPageData {
	s.NickName = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
