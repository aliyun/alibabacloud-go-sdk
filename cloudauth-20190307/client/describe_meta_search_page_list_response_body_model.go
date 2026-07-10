// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaSearchPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeMetaSearchPageListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeMetaSearchPageListResponseBodyItems) *DescribeMetaSearchPageListResponseBody
	GetItems() []*DescribeMetaSearchPageListResponseBodyItems
	SetPageSize(v int32) *DescribeMetaSearchPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMetaSearchPageListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMetaSearchPageListResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *DescribeMetaSearchPageListResponseBody
	GetTotalPage() *int32
}

type DescribeMetaSearchPageListResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The paginated list data.
	Items []*DescribeMetaSearchPageListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// 5176EB42-6EE7-510B-9388-35018DF3175B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 0
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeMetaSearchPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeMetaSearchPageListResponseBody) GetItems() []*DescribeMetaSearchPageListResponseBodyItems {
	return s.Items
}

func (s *DescribeMetaSearchPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetaSearchPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetaSearchPageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMetaSearchPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeMetaSearchPageListResponseBody) SetCurrentPage(v int32) *DescribeMetaSearchPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBody) SetItems(v []*DescribeMetaSearchPageListResponseBodyItems) *DescribeMetaSearchPageListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMetaSearchPageListResponseBody) SetPageSize(v int32) *DescribeMetaSearchPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBody) SetRequestId(v string) *DescribeMetaSearchPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBody) SetTotalCount(v int32) *DescribeMetaSearchPageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBody) SetTotalPage(v int32) *DescribeMetaSearchPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBody) Validate() error {
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

type DescribeMetaSearchPageListResponseBodyItems struct {
	// The product API. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification
	//
	// - **ID_PERIOD**: ID card validity period verification
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition)
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition)
	//
	// - **MOBILE_2_META**: mobile number two-element verification
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition)
	//
	// - **MOBILE_DETECT**: phone number detection
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition)
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query
	//
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// example:
	//
	// ID_CARD_2_META
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The name corresponding to the API. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification
	//
	// - **ID_PERIOD**: ID card validity period verification
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition)
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition)
	//
	// - **MOBILE_2_META**: mobile number two-element verification
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition)
	//
	// - **MOBILE_DETECT**: phone number detection
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition)
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query
	//
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// example:
	//
	// 身份证二要素
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The bank card ID.
	//
	// example:
	//
	// 6214837145546986
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// The verification status. Valid values:
	//
	// - **1**: Verification passed.
	//
	// - **2**: Verification failed.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The verification date.
	//
	// example:
	//
	// 2025-10-17 10:00:11
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The ID card number.
	//
	// example:
	//
	// 4****************X
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The name of the telecommunications service provider. Valid values:
	//
	// - **CMCC**: China Mobile
	//
	// - **CUCC**: China Unicom
	//
	// - **CTCC**: China Telecom.
	//
	// example:
	//
	// CUCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 1500000xxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The request parameters.
	Request *DescribeMetaSearchPageListResponseBodyItemsRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B0102BEF-4411-57C3-860D-CFE7DE0A64C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP request parameters.
	//
	// example:
	//
	// {\\"productCode\\":\\"cloudauth_infoverify_public_cn\\",\\"identifyNum\\":\\"4****************X\\",\\"userName\\":\\"**豪\\",\\"userId\\":\\"12600512516xxxxx\\"}
	RequestJson *string `json:"RequestJson,omitempty" xml:"RequestJson,omitempty"`
	// The response content.
	Response *DescribeMetaSearchPageListResponseBodyItemsResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// The response content.
	//
	// example:
	//
	// {\\"code\\":\\"200\\",\\"data\\":{\\"bizCode\\":\\"1\\",\\"ispName\\":\\"CMCC\\",\\"subCode\\":\\"101\\"},\\"message\\":\\"调用成功\\"}
	ResponseJson *string `json:"ResponseJson,omitempty" xml:"ResponseJson,omitempty"`
	// The result code. For more information, see [official documentation](https://www.alibabacloud.com/help/en/id-verification/information-verification/).
	//
	// example:
	//
	// 207
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The name.
	//
	// example:
	//
	// **豪
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The vehicle number.
	//
	// example:
	//
	// 浙J3H91H
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
}

func (s DescribeMetaSearchPageListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetBankCard() *string {
	return s.BankCard
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetDate() *string {
	return s.Date
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetIspName() *string {
	return s.IspName
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetMobile() *string {
	return s.Mobile
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetRequest() *DescribeMetaSearchPageListResponseBodyItemsRequest {
	return s.Request
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetRequestJson() *string {
	return s.RequestJson
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetResponse() *DescribeMetaSearchPageListResponseBodyItemsResponse {
	return s.Response
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetResponseJson() *string {
	return s.ResponseJson
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetSubCode() *string {
	return s.SubCode
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMetaSearchPageListResponseBodyItems) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetApi(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.Api = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetApiName(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.ApiName = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetBankCard(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.BankCard = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetBizCode(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.BizCode = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetDate(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetIdentifyNum(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.IdentifyNum = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetIspName(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.IspName = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetMobile(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.Mobile = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetRequest(v *DescribeMetaSearchPageListResponseBodyItemsRequest) *DescribeMetaSearchPageListResponseBodyItems {
	s.Request = v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetRequestId(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetRequestJson(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.RequestJson = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetResponse(v *DescribeMetaSearchPageListResponseBodyItemsResponse) *DescribeMetaSearchPageListResponseBodyItems {
	s.Response = v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetResponseJson(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.ResponseJson = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetSubCode(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.SubCode = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetUserName(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) SetVehicleNum(v string) *DescribeMetaSearchPageListResponseBodyItems {
	s.VehicleNum = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItems) Validate() error {
	if s.Request != nil {
		if err := s.Request.Validate(); err != nil {
			return err
		}
	}
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetaSearchPageListResponseBodyItemsRequest struct {
	// The vehicle number.
	//
	// example:
	//
	// 浙J3H91H
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// The vehicle type. Valid values:
	//
	// - **01**: large vehicle
	//
	// - **02**: small vehicle
	//
	// - **03**: embassy vehicle
	//
	// - **04**: consulate vehicle
	//
	// - **05**: overseas vehicle
	//
	// - **06**: foreign vehicle
	//
	// - **07**: standard motorcycle
	//
	// - **08**: light motorcycle
	//
	// - **09**: embassy motorcycle
	//
	// - **10**: consulate motorcycle
	//
	// - **11**: overseas motorcycle
	//
	// - **12**: foreign motorcycle
	//
	// - **13**: low-speed vehicle
	//
	// - **14**: tractor
	//
	// - **15**: trailer
	//
	// - **16**: training vehicle
	//
	// - **17**: training motorcycle
	//
	// - **20**: temporary entry vehicle
	//
	// - **21**: temporary entry motorcycle
	//
	// - **22**: temporary driving vehicle
	//
	// - **23**: police vehicle
	//
	// - **24**: police motorcycle
	//
	// - **51**: new energy large vehicle
	//
	// - **52**: new energy small vehicle.
	//
	// example:
	//
	// 52
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// The vehicle type name.
	//
	// example:
	//
	// 新能源小型车
	VehicleTypeName *string `json:"VehicleTypeName,omitempty" xml:"VehicleTypeName,omitempty"`
}

func (s DescribeMetaSearchPageListResponseBodyItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListResponseBodyItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) GetVehicleType() *string {
	return s.VehicleType
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) GetVehicleTypeName() *string {
	return s.VehicleTypeName
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) SetVehicleNum(v string) *DescribeMetaSearchPageListResponseBodyItemsRequest {
	s.VehicleNum = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) SetVehicleType(v string) *DescribeMetaSearchPageListResponseBodyItemsRequest {
	s.VehicleType = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) SetVehicleTypeName(v string) *DescribeMetaSearchPageListResponseBodyItemsRequest {
	s.VehicleTypeName = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeMetaSearchPageListResponseBodyItemsResponse struct {
	// The HTTP return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *DescribeMetaSearchPageListResponseBodyItemsResponseData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeMetaSearchPageListResponseBodyItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListResponseBodyItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) GetCode() *string {
	return s.Code
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) GetData() *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	return s.Data
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) SetCode(v string) *DescribeMetaSearchPageListResponseBodyItemsResponse {
	s.Code = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) SetData(v *DescribeMetaSearchPageListResponseBodyItemsResponseData) *DescribeMetaSearchPageListResponseBodyItemsResponse {
	s.Data = v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) SetMessage(v string) *DescribeMetaSearchPageListResponseBodyItemsResponse {
	s.Message = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponse) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetaSearchPageListResponseBodyItemsResponseData struct {
	// The approved number of passengers.
	//
	// example:
	//
	// 5
	ApprovedCount *string `json:"ApprovedCount,omitempty" xml:"ApprovedCount,omitempty"`
	// The approved passenger capacity.
	//
	// example:
	//
	// 0
	ApprovedLoad *string `json:"ApprovedLoad,omitempty" xml:"ApprovedLoad,omitempty"`
	// The number of axles.
	//
	// example:
	//
	// 2
	AxleCount *string `json:"AxleCount,omitempty" xml:"AxleCount,omitempty"`
	// The rear wheel track.
	//
	// example:
	//
	// 1630
	BackWheelDistance *string `json:"BackWheelDistance,omitempty" xml:"BackWheelDistance,omitempty"`
	// The bank card number.
	//
	// example:
	//
	// 622848001714440xxxx
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// The verification status. Valid values:
	//
	// - **1**: Verification passed.
	//
	// - **2**: Verification failed.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 2
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The brand name.
	//
	// example:
	//
	// 奔驰
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// The body color.
	//
	// example:
	//
	// 白
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The displacement.
	//
	// example:
	//
	// 0
	Displacement *string `json:"Displacement,omitempty" xml:"Displacement,omitempty"`
	// The engine number.
	//
	// example:
	//
	// N7YJ113PA
	EngineNum *string `json:"EngineNum,omitempty" xml:"EngineNum,omitempty"`
	// The engine model.
	//
	// example:
	//
	// CA6GV30TD
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The front wheel track.
	//
	// example:
	//
	// 1630
	FrontWheelDistance *string `json:"FrontWheelDistance,omitempty" xml:"FrontWheelDistance,omitempty"`
	// The fuel type.
	//
	// example:
	//
	// 电
	FuelType *string `json:"FuelType,omitempty" xml:"FuelType,omitempty"`
	// The inspection validity expiration date.
	//
	// example:
	//
	// 2026-07-31 00:00:00
	InspectionDate *string `json:"InspectionDate,omitempty" xml:"InspectionDate,omitempty"`
	// The vehicle model number.
	//
	// example:
	//
	// HQ7002BEV67
	ModelNum *string `json:"ModelNum,omitempty" xml:"ModelNum,omitempty"`
	// The power.
	//
	// example:
	//
	// 120
	Power *string `json:"Power,omitempty" xml:"Power,omitempty"`
	// The initial registration date.
	//
	// example:
	//
	// 2015-08-24 00:00:00
	RegistrationDate *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	// The manufacturing date.
	//
	// example:
	//
	// 2022-08-13 00:00:00
	ReleaseDate *string `json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	// The mandatory retirement date.
	//
	// example:
	//
	// 2099-12-31 00:00:00
	RetirementDate *string `json:"RetirementDate,omitempty" xml:"RetirementDate,omitempty"`
	// The gross mass.
	//
	// example:
	//
	// 2350
	TotalMass *string `json:"TotalMass,omitempty" xml:"TotalMass,omitempty"`
	// The vehicle type.
	//
	// example:
	//
	// 小型轿车
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The curb weight.
	//
	// example:
	//
	// 1900
	UnladenMass *string `json:"UnladenMass,omitempty" xml:"UnladenMass,omitempty"`
	// The usage property.
	//
	// example:
	//
	// 0
	UseProperty *string `json:"UseProperty,omitempty" xml:"UseProperty,omitempty"`
	// The motor vehicle status.
	//
	// example:
	//
	// 正常
	VehicleState *string `json:"VehicleState,omitempty" xml:"VehicleState,omitempty"`
	// The vehicle identification number (VIN).
	//
	// example:
	//
	// LDC643T44G3667219
	Vin *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
	// The wheelbase.
	//
	// example:
	//
	// 2750
	WheelBase *string `json:"WheelBase,omitempty" xml:"WheelBase,omitempty"`
}

func (s DescribeMetaSearchPageListResponseBodyItemsResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListResponseBodyItemsResponseData) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetApprovedCount() *string {
	return s.ApprovedCount
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetApprovedLoad() *string {
	return s.ApprovedLoad
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetAxleCount() *string {
	return s.AxleCount
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetBackWheelDistance() *string {
	return s.BackWheelDistance
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetBankCard() *string {
	return s.BankCard
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetBrand() *string {
	return s.Brand
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetColor() *string {
	return s.Color
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetDisplacement() *string {
	return s.Displacement
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetEngineNum() *string {
	return s.EngineNum
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetFrontWheelDistance() *string {
	return s.FrontWheelDistance
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetFuelType() *string {
	return s.FuelType
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetInspectionDate() *string {
	return s.InspectionDate
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetModelNum() *string {
	return s.ModelNum
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetPower() *string {
	return s.Power
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetRegistrationDate() *string {
	return s.RegistrationDate
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetRetirementDate() *string {
	return s.RetirementDate
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetTotalMass() *string {
	return s.TotalMass
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetType() *string {
	return s.Type
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetUnladenMass() *string {
	return s.UnladenMass
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetUseProperty() *string {
	return s.UseProperty
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetVehicleState() *string {
	return s.VehicleState
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetVin() *string {
	return s.Vin
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) GetWheelBase() *string {
	return s.WheelBase
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetApprovedCount(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.ApprovedCount = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetApprovedLoad(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.ApprovedLoad = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetAxleCount(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.AxleCount = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetBackWheelDistance(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.BackWheelDistance = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetBankCard(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.BankCard = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetBizCode(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.BizCode = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetBrand(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.Brand = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetColor(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.Color = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetDisplacement(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.Displacement = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetEngineNum(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.EngineNum = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetEngineType(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.EngineType = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetFrontWheelDistance(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.FrontWheelDistance = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetFuelType(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.FuelType = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetInspectionDate(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.InspectionDate = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetModelNum(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.ModelNum = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetPower(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.Power = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetRegistrationDate(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.RegistrationDate = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetReleaseDate(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.ReleaseDate = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetRetirementDate(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.RetirementDate = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetTotalMass(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.TotalMass = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetType(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.Type = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetUnladenMass(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.UnladenMass = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetUseProperty(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.UseProperty = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetVehicleState(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.VehicleState = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetVin(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.Vin = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) SetWheelBase(v string) *DescribeMetaSearchPageListResponseBodyItemsResponseData {
	s.WheelBase = &v
	return s
}

func (s *DescribeMetaSearchPageListResponseBodyItemsResponseData) Validate() error {
	return dara.Validate(s)
}
