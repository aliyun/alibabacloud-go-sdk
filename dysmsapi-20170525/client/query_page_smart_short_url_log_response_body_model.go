// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPageSmartShortUrlLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPageSmartShortUrlLogResponseBody
	GetCode() *string
	SetMessage(v string) *QueryPageSmartShortUrlLogResponseBody
	GetMessage() *string
	SetModel(v *QueryPageSmartShortUrlLogResponseBodyModel) *QueryPageSmartShortUrlLogResponseBody
	GetModel() *QueryPageSmartShortUrlLogResponseBodyModel
	SetRequestId(v string) *QueryPageSmartShortUrlLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPageSmartShortUrlLogResponseBody
	GetSuccess() *bool
}

type QueryPageSmartShortUrlLogResponseBody struct {
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *QueryPageSmartShortUrlLogResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPageSmartShortUrlLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPageSmartShortUrlLogResponseBody) GetModel() *QueryPageSmartShortUrlLogResponseBodyModel {
	return s.Model
}

func (s *QueryPageSmartShortUrlLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPageSmartShortUrlLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetCode(v string) *QueryPageSmartShortUrlLogResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetMessage(v string) *QueryPageSmartShortUrlLogResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetModel(v *QueryPageSmartShortUrlLogResponseBodyModel) *QueryPageSmartShortUrlLogResponseBody {
	s.Model = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetRequestId(v string) *QueryPageSmartShortUrlLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetSuccess(v bool) *QueryPageSmartShortUrlLogResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPageSmartShortUrlLogResponseBodyModel struct {
	List []*QueryPageSmartShortUrlLogResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 74
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 15
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 66
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 86
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) GetList() []*QueryPageSmartShortUrlLogResponseBodyModelList {
	return s.List
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetList(v []*QueryPageSmartShortUrlLogResponseBodyModelList) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.List = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetPageNo(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetPageSize(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetTotalCount(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetTotalPage(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

type QueryPageSmartShortUrlLogResponseBodyModelList struct {
	// example:
	//
	// 87
	ClickState *int64 `json:"ClickState,omitempty" xml:"ClickState,omitempty"`
	// example:
	//
	// 51
	ClickTime *int64 `json:"ClickTime,omitempty" xml:"ClickTime,omitempty"`
	// example:
	//
	// 64
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 示例值示例值
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) GetClickState() *int64 {
	return s.ClickState
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) GetClickTime() *int64 {
	return s.ClickTime
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) GetShortName() *string {
	return s.ShortName
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) GetShortUrl() *string {
	return s.ShortUrl
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetClickState(v int64) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ClickState = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetClickTime(v int64) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ClickTime = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetCreateTime(v int64) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.CreateTime = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetPhoneNumber(v string) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.PhoneNumber = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetShortName(v string) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ShortName = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetShortUrl(v string) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ShortUrl = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
