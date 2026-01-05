// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTrademarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySmsTrademarkResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySmsTrademarkResponseBody
	GetCode() *string
	SetData(v []*QuerySmsTrademarkResponseBodyData) *QuerySmsTrademarkResponseBody
	GetData() []*QuerySmsTrademarkResponseBodyData
	SetMessage(v string) *QuerySmsTrademarkResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySmsTrademarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySmsTrademarkResponseBody
	GetSuccess() *bool
}

type QuerySmsTrademarkResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QuerySmsTrademarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 569E9DBD-23FF-1785-99AD-E4B23608C104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySmsTrademarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTrademarkResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsTrademarkResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySmsTrademarkResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsTrademarkResponseBody) GetData() []*QuerySmsTrademarkResponseBodyData {
	return s.Data
}

func (s *QuerySmsTrademarkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsTrademarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsTrademarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySmsTrademarkResponseBody) SetAccessDeniedDetail(v string) *QuerySmsTrademarkResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySmsTrademarkResponseBody) SetCode(v string) *QuerySmsTrademarkResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsTrademarkResponseBody) SetData(v []*QuerySmsTrademarkResponseBodyData) *QuerySmsTrademarkResponseBody {
	s.Data = v
	return s
}

func (s *QuerySmsTrademarkResponseBody) SetMessage(v string) *QuerySmsTrademarkResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsTrademarkResponseBody) SetRequestId(v string) *QuerySmsTrademarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsTrademarkResponseBody) SetSuccess(v bool) *QuerySmsTrademarkResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySmsTrademarkResponseBody) Validate() error {
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

type QuerySmsTrademarkResponseBodyData struct {
	// 申请人名称
	//
	// example:
	//
	// 示例值示例值
	TrademarkApplicantName *string `json:"TrademarkApplicantName,omitempty" xml:"TrademarkApplicantName,omitempty"`
	// 专用权生失效日期
	//
	// example:
	//
	// 2025-11-01~2025-12-19
	TrademarkEffExpDate *string `json:"TrademarkEffExpDate,omitempty" xml:"TrademarkEffExpDate,omitempty"`
	// 商标材料id
	//
	// example:
	//
	// 10000*******
	TrademarkId *int64 `json:"TrademarkId,omitempty" xml:"TrademarkId,omitempty"`
	// 商标名称
	//
	// example:
	//
	// 示例值示例值
	TrademarkName *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	// 商标截图Osskey（给签名传工单用）
	//
	// example:
	//
	// 100000*****802/afdebd46-****-46e4-899d-b4375826c898_mhk9oz0p_1762****31542.png
	TrademarkPic *string `json:"TrademarkPic,omitempty" xml:"TrademarkPic,omitempty"`
	// 商标截图url地址
	//
	// example:
	//
	// https://alicom-fc-media.oss-cn-zhangjiakou.aliyuncs.com/100000****50802/afde****-496d-46e4-899d-b43758****8_mhk9oz0p_176224****542.png?Expires=1762****6&OSSAccessKeyId=bypFN****73PsLI&Signature=BygI9X****h7%2FXmFIo****FB2c%3D
	TrademarkPicUrl *string `json:"TrademarkPicUrl,omitempty" xml:"TrademarkPicUrl,omitempty"`
	// 商标注册号
	//
	// example:
	//
	// 1234
	TrademarkRegistrationNumber *string `json:"TrademarkRegistrationNumber,omitempty" xml:"TrademarkRegistrationNumber,omitempty"`
}

func (s QuerySmsTrademarkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTrademarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkApplicantName() *string {
	return s.TrademarkApplicantName
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkEffExpDate() *string {
	return s.TrademarkEffExpDate
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkId() *int64 {
	return s.TrademarkId
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkName() *string {
	return s.TrademarkName
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkPic() *string {
	return s.TrademarkPic
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkPicUrl() *string {
	return s.TrademarkPicUrl
}

func (s *QuerySmsTrademarkResponseBodyData) GetTrademarkRegistrationNumber() *string {
	return s.TrademarkRegistrationNumber
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkApplicantName(v string) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkApplicantName = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkEffExpDate(v string) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkEffExpDate = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkId(v int64) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkId = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkName(v string) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkName = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkPic(v string) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkPic = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkPicUrl(v string) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkPicUrl = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) SetTrademarkRegistrationNumber(v string) *QuerySmsTrademarkResponseBodyData {
	s.TrademarkRegistrationNumber = &v
	return s
}

func (s *QuerySmsTrademarkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
