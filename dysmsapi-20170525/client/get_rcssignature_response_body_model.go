// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRCSSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetRCSSignatureResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetRCSSignatureResponseBody
	GetCode() *string
	SetData(v *GetRCSSignatureResponseBodyData) *GetRCSSignatureResponseBody
	GetData() *GetRCSSignatureResponseBodyData
	SetMessage(v string) *GetRCSSignatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRCSSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRCSSignatureResponseBody
	GetSuccess() *bool
}

type GetRCSSignatureResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRCSSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRCSSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetRCSSignatureResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRCSSignatureResponseBody) GetData() *GetRCSSignatureResponseBodyData {
	return s.Data
}

func (s *GetRCSSignatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRCSSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRCSSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRCSSignatureResponseBody) SetAccessDeniedDetail(v string) *GetRCSSignatureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetRCSSignatureResponseBody) SetCode(v string) *GetRCSSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *GetRCSSignatureResponseBody) SetData(v *GetRCSSignatureResponseBodyData) *GetRCSSignatureResponseBody {
	s.Data = v
	return s
}

func (s *GetRCSSignatureResponseBody) SetMessage(v string) *GetRCSSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *GetRCSSignatureResponseBody) SetRequestId(v string) *GetRCSSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRCSSignatureResponseBody) SetSuccess(v bool) *GetRCSSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *GetRCSSignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRCSSignatureResponseBodyData struct {
	// example:
	//
	// 示例值
	BackgroundImage *string `json:"BackgroundImage,omitempty" xml:"BackgroundImage,omitempty"`
	// example:
	//
	// 示例值
	BubbleColor *string `json:"BubbleColor,omitempty" xml:"BubbleColor,omitempty"`
	// example:
	//
	// 90
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ChatbotCode *string `json:"ChatbotCode,omitempty" xml:"ChatbotCode,omitempty"`
	// example:
	//
	// 示例值示例值
	ChatbotName *string `json:"ChatbotName,omitempty" xml:"ChatbotName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 示例值示例值
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 示例值示例值
	Logo *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OfficeAddress      *string                                              `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	RegisterResultList []*GetRCSSignatureResponseBodyDataRegisterResultList `json:"RegisterResultList,omitempty" xml:"RegisterResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
	ServiceEmail *string `json:"ServiceEmail,omitempty" xml:"ServiceEmail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ServicePhone *string `json:"ServicePhone,omitempty" xml:"ServicePhone,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ServiceTerms *string `json:"ServiceTerms,omitempty" xml:"ServiceTerms,omitempty"`
	// example:
	//
	// 示例值示例值
	ServiceWebsite  *string                                           `json:"ServiceWebsite,omitempty" xml:"ServiceWebsite,omitempty"`
	ShelfResultList []*GetRCSSignatureResponseBodyDataShelfResultList `json:"ShelfResultList,omitempty" xml:"ShelfResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 32
	SignId *int64 `json:"SignId,omitempty" xml:"SignId,omitempty"`
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s GetRCSSignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponseBodyData) GetBackgroundImage() *string {
	return s.BackgroundImage
}

func (s *GetRCSSignatureResponseBodyData) GetBubbleColor() *string {
	return s.BubbleColor
}

func (s *GetRCSSignatureResponseBodyData) GetCategory() *int64 {
	return s.Category
}

func (s *GetRCSSignatureResponseBodyData) GetChatbotCode() *string {
	return s.ChatbotCode
}

func (s *GetRCSSignatureResponseBodyData) GetChatbotName() *string {
	return s.ChatbotName
}

func (s *GetRCSSignatureResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetRCSSignatureResponseBodyData) GetLatitude() *string {
	return s.Latitude
}

func (s *GetRCSSignatureResponseBodyData) GetLogo() *string {
	return s.Logo
}

func (s *GetRCSSignatureResponseBodyData) GetLongitude() *string {
	return s.Longitude
}

func (s *GetRCSSignatureResponseBodyData) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *GetRCSSignatureResponseBodyData) GetRegisterResultList() []*GetRCSSignatureResponseBodyDataRegisterResultList {
	return s.RegisterResultList
}

func (s *GetRCSSignatureResponseBodyData) GetServiceEmail() *string {
	return s.ServiceEmail
}

func (s *GetRCSSignatureResponseBodyData) GetServicePhone() *string {
	return s.ServicePhone
}

func (s *GetRCSSignatureResponseBodyData) GetServiceTerms() *string {
	return s.ServiceTerms
}

func (s *GetRCSSignatureResponseBodyData) GetServiceWebsite() *string {
	return s.ServiceWebsite
}

func (s *GetRCSSignatureResponseBodyData) GetShelfResultList() []*GetRCSSignatureResponseBodyDataShelfResultList {
	return s.ShelfResultList
}

func (s *GetRCSSignatureResponseBodyData) GetSignId() *int64 {
	return s.SignId
}

func (s *GetRCSSignatureResponseBodyData) GetSignName() *string {
	return s.SignName
}

func (s *GetRCSSignatureResponseBodyData) SetBackgroundImage(v string) *GetRCSSignatureResponseBodyData {
	s.BackgroundImage = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetBubbleColor(v string) *GetRCSSignatureResponseBodyData {
	s.BubbleColor = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetCategory(v int64) *GetRCSSignatureResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetChatbotCode(v string) *GetRCSSignatureResponseBodyData {
	s.ChatbotCode = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetChatbotName(v string) *GetRCSSignatureResponseBodyData {
	s.ChatbotName = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetDescription(v string) *GetRCSSignatureResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetLatitude(v string) *GetRCSSignatureResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetLogo(v string) *GetRCSSignatureResponseBodyData {
	s.Logo = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetLongitude(v string) *GetRCSSignatureResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetOfficeAddress(v string) *GetRCSSignatureResponseBodyData {
	s.OfficeAddress = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetRegisterResultList(v []*GetRCSSignatureResponseBodyDataRegisterResultList) *GetRCSSignatureResponseBodyData {
	s.RegisterResultList = v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetServiceEmail(v string) *GetRCSSignatureResponseBodyData {
	s.ServiceEmail = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetServicePhone(v string) *GetRCSSignatureResponseBodyData {
	s.ServicePhone = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetServiceTerms(v string) *GetRCSSignatureResponseBodyData {
	s.ServiceTerms = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetServiceWebsite(v string) *GetRCSSignatureResponseBodyData {
	s.ServiceWebsite = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetShelfResultList(v []*GetRCSSignatureResponseBodyDataShelfResultList) *GetRCSSignatureResponseBodyData {
	s.ShelfResultList = v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetSignId(v int64) *GetRCSSignatureResponseBodyData {
	s.SignId = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) SetSignName(v string) *GetRCSSignatureResponseBodyData {
	s.SignName = &v
	return s
}

func (s *GetRCSSignatureResponseBodyData) Validate() error {
	if s.RegisterResultList != nil {
		for _, item := range s.RegisterResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ShelfResultList != nil {
		for _, item := range s.ShelfResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRCSSignatureResponseBodyDataRegisterResultList struct {
	// example:
	//
	// 示例值示例值
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
	// example:
	//
	// 46
	ProductType *int64 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RegisterCompleteTime *string `json:"RegisterCompleteTime,omitempty" xml:"RegisterCompleteTime,omitempty"`
	// example:
	//
	// 72
	RegisterStatus        *int64                                                                    `json:"RegisterStatus,omitempty" xml:"RegisterStatus,omitempty"`
	RegisterStatusReasons []*GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons `json:"RegisterStatusReasons,omitempty" xml:"RegisterStatusReasons,omitempty" type:"Repeated"`
}

func (s GetRCSSignatureResponseBodyDataRegisterResultList) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponseBodyDataRegisterResultList) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) GetRegisterCompleteTime() *string {
	return s.RegisterCompleteTime
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) GetRegisterStatus() *int64 {
	return s.RegisterStatus
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) GetRegisterStatusReasons() []*GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons {
	return s.RegisterStatusReasons
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) SetOperatorCode(v string) *GetRCSSignatureResponseBodyDataRegisterResultList {
	s.OperatorCode = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) SetProductType(v int64) *GetRCSSignatureResponseBodyDataRegisterResultList {
	s.ProductType = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) SetRegisterCompleteTime(v string) *GetRCSSignatureResponseBodyDataRegisterResultList {
	s.RegisterCompleteTime = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) SetRegisterStatus(v int64) *GetRCSSignatureResponseBodyDataRegisterResultList {
	s.RegisterStatus = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) SetRegisterStatusReasons(v []*GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) *GetRCSSignatureResponseBodyDataRegisterResultList {
	s.RegisterStatusReasons = v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultList) Validate() error {
	if s.RegisterStatusReasons != nil {
		for _, item := range s.RegisterStatusReasons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons struct {
	// example:
	//
	// 示例值示例值
	ReasonCode     *string   `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonDescList []*string `json:"ReasonDescList,omitempty" xml:"ReasonDescList,omitempty" type:"Repeated"`
}

func (s GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) GetReasonDescList() []*string {
	return s.ReasonDescList
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) SetReasonCode(v string) *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons {
	s.ReasonCode = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) SetReasonDescList(v []*string) *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons {
	s.ReasonDescList = v
	return s
}

func (s *GetRCSSignatureResponseBodyDataRegisterResultListRegisterStatusReasons) Validate() error {
	return dara.Validate(s)
}

type GetRCSSignatureResponseBodyDataShelfResultList struct {
	// example:
	//
	// 示例值示例值
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
	// example:
	//
	// 81
	ProductType *int64 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 51
	ShelfStatus        *int64                                                              `json:"ShelfStatus,omitempty" xml:"ShelfStatus,omitempty"`
	ShelfStatusReasons []*GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons `json:"ShelfStatusReasons,omitempty" xml:"ShelfStatusReasons,omitempty" type:"Repeated"`
}

func (s GetRCSSignatureResponseBodyDataShelfResultList) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponseBodyDataShelfResultList) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) GetShelfStatus() *int64 {
	return s.ShelfStatus
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) GetShelfStatusReasons() []*GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons {
	return s.ShelfStatusReasons
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) SetOperatorCode(v string) *GetRCSSignatureResponseBodyDataShelfResultList {
	s.OperatorCode = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) SetProductType(v int64) *GetRCSSignatureResponseBodyDataShelfResultList {
	s.ProductType = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) SetShelfStatus(v int64) *GetRCSSignatureResponseBodyDataShelfResultList {
	s.ShelfStatus = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) SetShelfStatusReasons(v []*GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) *GetRCSSignatureResponseBodyDataShelfResultList {
	s.ShelfStatusReasons = v
	return s
}

func (s *GetRCSSignatureResponseBodyDataShelfResultList) Validate() error {
	if s.ShelfStatusReasons != nil {
		for _, item := range s.ShelfStatusReasons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons struct {
	// example:
	//
	// 示例值示例值示例值
	ReasonCode     *string   `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonDescList []*string `json:"ReasonDescList,omitempty" xml:"ReasonDescList,omitempty" type:"Repeated"`
}

func (s GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) GetReasonDescList() []*string {
	return s.ReasonDescList
}

func (s *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) SetReasonCode(v string) *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons {
	s.ReasonCode = &v
	return s
}

func (s *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) SetReasonDescList(v []*string) *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons {
	s.ReasonDescList = v
	return s
}

func (s *GetRCSSignatureResponseBodyDataShelfResultListShelfStatusReasons) Validate() error {
	return dara.Validate(s)
}
