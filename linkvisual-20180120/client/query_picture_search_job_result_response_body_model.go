// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchJobResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPictureSearchJobResultResponseBody
	GetCode() *string
	SetData(v *QueryPictureSearchJobResultResponseBodyData) *QueryPictureSearchJobResultResponseBody
	GetData() *QueryPictureSearchJobResultResponseBodyData
	SetErrorMessage(v string) *QueryPictureSearchJobResultResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryPictureSearchJobResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPictureSearchJobResultResponseBody
	GetSuccess() *bool
}

type QueryPictureSearchJobResultResponseBody struct {
	// The error code returned if the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryPictureSearchJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// AppInstanceId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a06c4750-7cce-11eb-8010-557d50136505
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

func (s QueryPictureSearchJobResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPictureSearchJobResultResponseBody) GetData() *QueryPictureSearchJobResultResponseBodyData {
	return s.Data
}

func (s *QueryPictureSearchJobResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryPictureSearchJobResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPictureSearchJobResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPictureSearchJobResultResponseBody) SetCode(v string) *QueryPictureSearchJobResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetData(v *QueryPictureSearchJobResultResponseBodyData) *QueryPictureSearchJobResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetErrorMessage(v string) *QueryPictureSearchJobResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetRequestId(v string) *QueryPictureSearchJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetSuccess(v bool) *QueryPictureSearchJobResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPictureSearchJobResultResponseBodyData struct {
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
	// The query results.
	PageData []*QueryPictureSearchJobResultResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
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

func (s QueryPictureSearchJobResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchJobResultResponseBodyData) GetPageCount() *int32 {
	return s.PageCount
}

func (s *QueryPictureSearchJobResultResponseBodyData) GetPageData() []*QueryPictureSearchJobResultResponseBodyDataPageData {
	return s.PageData
}

func (s *QueryPictureSearchJobResultResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchJobResultResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetPageCount(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetPageData(v []*QueryPictureSearchJobResultResponseBodyDataPageData) *QueryPictureSearchJobResultResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetPageSize(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetTotal(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) Validate() error {
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

type QueryPictureSearchJobResultResponseBodyDataPageData struct {
	// The alias of the device.
	//
	// > If no alias is set for the device, this parameter is not returned.
	//
	// example:
	//
	// camera1
	DeviceNickName *string `json:"DeviceNickName,omitempty" xml:"DeviceNickName,omitempty"`
	// The time when the event was created. The value is a 13-digit timestamp in milliseconds.
	//
	// example:
	//
	// 1614333155075
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The IoT ID of the V-series edge all-in-one device.
	//
	// example:
	//
	// pZJUqijOyfpFOOGNPF7S00****
	GatewayIotId *string `json:"GatewayIotId,omitempty" xml:"GatewayIotId,omitempty"`
	// The IoT ID of the queried device.
	//
	// example:
	//
	// 5RUngwC1RknD4lPiHgH300****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The URL of the image to be searched.
	//
	// example:
	//
	// https://xxx.xxx.xxx.com/xxx
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	// The face similarity.
	//
	// example:
	//
	// 0.5
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The ID of the image feature.
	//
	// example:
	//
	// vectorId86d9a5c32b9947278e526c2eef8a****
	VectorId *string `json:"VectorId,omitempty" xml:"VectorId,omitempty"`
}

func (s QueryPictureSearchJobResultResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobResultResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetDeviceNickName() *string {
	return s.DeviceNickName
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetEventTime() *int64 {
	return s.EventTime
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetGatewayIotId() *string {
	return s.GatewayIotId
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetIotId() *string {
	return s.IotId
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetPicUrl() *string {
	return s.PicUrl
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) GetVectorId() *string {
	return s.VectorId
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetDeviceNickName(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.DeviceNickName = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetEventTime(v int64) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.EventTime = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetGatewayIotId(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.GatewayIotId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetPicUrl(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.PicUrl = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetThreshold(v float32) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.Threshold = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetVectorId(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.VectorId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
