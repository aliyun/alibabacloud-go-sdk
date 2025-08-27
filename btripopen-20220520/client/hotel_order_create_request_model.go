// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderCreateRequest
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelOrderCreateRequest
	GetCheckIn() *string
	SetCheckOut(v string) *HotelOrderCreateRequest
	GetCheckOut() *string
	SetContractEmail(v string) *HotelOrderCreateRequest
	GetContractEmail() *string
	SetContractName(v string) *HotelOrderCreateRequest
	GetContractName() *string
	SetContractPhone(v string) *HotelOrderCreateRequest
	GetContractPhone() *string
	SetCorpPayPrice(v int64) *HotelOrderCreateRequest
	GetCorpPayPrice() *int64
	SetDisOrderId(v string) *HotelOrderCreateRequest
	GetDisOrderId() *string
	SetExtra(v string) *HotelOrderCreateRequest
	GetExtra() *string
	SetInvoiceInfo(v *HotelOrderCreateRequestInvoiceInfo) *HotelOrderCreateRequest
	GetInvoiceInfo() *HotelOrderCreateRequestInvoiceInfo
	SetItemId(v int64) *HotelOrderCreateRequest
	GetItemId() *int64
	SetItineraryNo(v string) *HotelOrderCreateRequest
	GetItineraryNo() *string
	SetOccupantInfoList(v []*HotelOrderCreateRequestOccupantInfoList) *HotelOrderCreateRequest
	GetOccupantInfoList() []*HotelOrderCreateRequestOccupantInfoList
	SetPersonPayPrice(v int64) *HotelOrderCreateRequest
	GetPersonPayPrice() *int64
	SetPromotionInfo(v *HotelOrderCreateRequestPromotionInfo) *HotelOrderCreateRequest
	GetPromotionInfo() *HotelOrderCreateRequestPromotionInfo
	SetRatePlanId(v int64) *HotelOrderCreateRequest
	GetRatePlanId() *int64
	SetRoomId(v int64) *HotelOrderCreateRequest
	GetRoomId() *int64
	SetRoomNum(v int32) *HotelOrderCreateRequest
	GetRoomNum() *int32
	SetSellerId(v int64) *HotelOrderCreateRequest
	GetSellerId() *int64
	SetShid(v int64) *HotelOrderCreateRequest
	GetShid() *int64
	SetTotalOrderPrice(v int64) *HotelOrderCreateRequest
	GetTotalOrderPrice() *int64
	SetValidateResKey(v string) *HotelOrderCreateRequest
	GetValidateResKey() *string
}

type HotelOrderCreateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-10-20
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-10-20
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// demo
	ContractEmail *string `json:"contract_email,omitempty" xml:"contract_email,omitempty"`
	ContractName  *string `json:"contract_name,omitempty" xml:"contract_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19281772123
	ContractPhone *string `json:"contract_phone,omitempty" xml:"contract_phone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	CorpPayPrice *int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId  *string                             `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	Extra       *string                             `json:"extra,omitempty" xml:"extra,omitempty"`
	InvoiceInfo *HotelOrderCreateRequestInvoiceInfo `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 671570615157
	ItemId *int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fb5e1abf33924b6c912bd6d80deec0eb-1
	ItineraryNo *string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	// This parameter is required.
	OccupantInfoList []*HotelOrderCreateRequestOccupantInfoList `json:"occupant_info_list,omitempty" xml:"occupant_info_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	PersonPayPrice *int64                                `json:"person_pay_price,omitempty" xml:"person_pay_price,omitempty"`
	PromotionInfo  *HotelOrderCreateRequestPromotionInfo `json:"promotion_info,omitempty" xml:"promotion_info,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1399417428510
	RatePlanId *int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 187211
	RoomId *int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomNum *int32 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2088441675613762
	SellerId *int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2198781
	Shid *int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalOrderPrice *int64 `json:"total_order_price,omitempty" xml:"total_order_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nonUltron_1673575241156_d91ea8ad16735752359161037bf6cf_c54d3768312a4b249b719f126377bf82
	ValidateResKey *string `json:"validate_res_key,omitempty" xml:"validate_res_key,omitempty"`
}

func (s HotelOrderCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderCreateRequest) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderCreateRequest) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelOrderCreateRequest) GetContractEmail() *string {
	return s.ContractEmail
}

func (s *HotelOrderCreateRequest) GetContractName() *string {
	return s.ContractName
}

func (s *HotelOrderCreateRequest) GetContractPhone() *string {
	return s.ContractPhone
}

func (s *HotelOrderCreateRequest) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *HotelOrderCreateRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderCreateRequest) GetExtra() *string {
	return s.Extra
}

func (s *HotelOrderCreateRequest) GetInvoiceInfo() *HotelOrderCreateRequestInvoiceInfo {
	return s.InvoiceInfo
}

func (s *HotelOrderCreateRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *HotelOrderCreateRequest) GetItineraryNo() *string {
	return s.ItineraryNo
}

func (s *HotelOrderCreateRequest) GetOccupantInfoList() []*HotelOrderCreateRequestOccupantInfoList {
	return s.OccupantInfoList
}

func (s *HotelOrderCreateRequest) GetPersonPayPrice() *int64 {
	return s.PersonPayPrice
}

func (s *HotelOrderCreateRequest) GetPromotionInfo() *HotelOrderCreateRequestPromotionInfo {
	return s.PromotionInfo
}

func (s *HotelOrderCreateRequest) GetRatePlanId() *int64 {
	return s.RatePlanId
}

func (s *HotelOrderCreateRequest) GetRoomId() *int64 {
	return s.RoomId
}

func (s *HotelOrderCreateRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderCreateRequest) GetSellerId() *int64 {
	return s.SellerId
}

func (s *HotelOrderCreateRequest) GetShid() *int64 {
	return s.Shid
}

func (s *HotelOrderCreateRequest) GetTotalOrderPrice() *int64 {
	return s.TotalOrderPrice
}

func (s *HotelOrderCreateRequest) GetValidateResKey() *string {
	return s.ValidateResKey
}

func (s *HotelOrderCreateRequest) SetBtripUserId(v string) *HotelOrderCreateRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderCreateRequest) SetCheckIn(v string) *HotelOrderCreateRequest {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderCreateRequest) SetCheckOut(v string) *HotelOrderCreateRequest {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderCreateRequest) SetContractEmail(v string) *HotelOrderCreateRequest {
	s.ContractEmail = &v
	return s
}

func (s *HotelOrderCreateRequest) SetContractName(v string) *HotelOrderCreateRequest {
	s.ContractName = &v
	return s
}

func (s *HotelOrderCreateRequest) SetContractPhone(v string) *HotelOrderCreateRequest {
	s.ContractPhone = &v
	return s
}

func (s *HotelOrderCreateRequest) SetCorpPayPrice(v int64) *HotelOrderCreateRequest {
	s.CorpPayPrice = &v
	return s
}

func (s *HotelOrderCreateRequest) SetDisOrderId(v string) *HotelOrderCreateRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderCreateRequest) SetExtra(v string) *HotelOrderCreateRequest {
	s.Extra = &v
	return s
}

func (s *HotelOrderCreateRequest) SetInvoiceInfo(v *HotelOrderCreateRequestInvoiceInfo) *HotelOrderCreateRequest {
	s.InvoiceInfo = v
	return s
}

func (s *HotelOrderCreateRequest) SetItemId(v int64) *HotelOrderCreateRequest {
	s.ItemId = &v
	return s
}

func (s *HotelOrderCreateRequest) SetItineraryNo(v string) *HotelOrderCreateRequest {
	s.ItineraryNo = &v
	return s
}

func (s *HotelOrderCreateRequest) SetOccupantInfoList(v []*HotelOrderCreateRequestOccupantInfoList) *HotelOrderCreateRequest {
	s.OccupantInfoList = v
	return s
}

func (s *HotelOrderCreateRequest) SetPersonPayPrice(v int64) *HotelOrderCreateRequest {
	s.PersonPayPrice = &v
	return s
}

func (s *HotelOrderCreateRequest) SetPromotionInfo(v *HotelOrderCreateRequestPromotionInfo) *HotelOrderCreateRequest {
	s.PromotionInfo = v
	return s
}

func (s *HotelOrderCreateRequest) SetRatePlanId(v int64) *HotelOrderCreateRequest {
	s.RatePlanId = &v
	return s
}

func (s *HotelOrderCreateRequest) SetRoomId(v int64) *HotelOrderCreateRequest {
	s.RoomId = &v
	return s
}

func (s *HotelOrderCreateRequest) SetRoomNum(v int32) *HotelOrderCreateRequest {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderCreateRequest) SetSellerId(v int64) *HotelOrderCreateRequest {
	s.SellerId = &v
	return s
}

func (s *HotelOrderCreateRequest) SetShid(v int64) *HotelOrderCreateRequest {
	s.Shid = &v
	return s
}

func (s *HotelOrderCreateRequest) SetTotalOrderPrice(v int64) *HotelOrderCreateRequest {
	s.TotalOrderPrice = &v
	return s
}

func (s *HotelOrderCreateRequest) SetValidateResKey(v string) *HotelOrderCreateRequest {
	s.ValidateResKey = &v
	return s
}

func (s *HotelOrderCreateRequest) Validate() error {
	return dara.Validate(s)
}

type HotelOrderCreateRequestInvoiceInfo struct {
	// example:
	//
	// demo
	BuyerAdd *string `json:"buyer_add,omitempty" xml:"buyer_add,omitempty"`
	// example:
	//
	// demo
	BuyerBankAcc *string `json:"buyer_bank_acc,omitempty" xml:"buyer_bank_acc,omitempty"`
	// example:
	//
	// demo
	BuyerBankAdd *string `json:"buyer_bank_add,omitempty" xml:"buyer_bank_add,omitempty"`
	// example:
	//
	// 0571-82321777
	BuyerPhone *string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// example:
	//
	// 1
	BuyerTaxNum      *string `json:"buyer_tax_num,omitempty" xml:"buyer_tax_num,omitempty"`
	DeliveryAddress  *string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	DeliveryArea     *string `json:"delivery_area,omitempty" xml:"delivery_area,omitempty"`
	DeliveryCity     *string `json:"delivery_city,omitempty" xml:"delivery_city,omitempty"`
	DeliveryProvince *string `json:"delivery_province,omitempty" xml:"delivery_province,omitempty"`
	DeliveryStreet   *string `json:"delivery_street,omitempty" xml:"delivery_street,omitempty"`
	// example:
	//
	// demo
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1
	InvoiceMaterial *int32 `json:"invoice_material,omitempty" xml:"invoice_material,omitempty"`
	// example:
	//
	// demo
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// 1
	InvoiceType  *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	ReceiverName *string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// example:
	//
	// 0571-82321777
	ReceiverPhone *string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// example:
	//
	// demo
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s HotelOrderCreateRequestInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateRequestInvoiceInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetBuyerAdd() *string {
	return s.BuyerAdd
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetBuyerBankAcc() *string {
	return s.BuyerBankAcc
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetBuyerBankAdd() *string {
	return s.BuyerBankAdd
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetBuyerPhone() *string {
	return s.BuyerPhone
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetBuyerTaxNum() *string {
	return s.BuyerTaxNum
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetDeliveryAddress() *string {
	return s.DeliveryAddress
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetDeliveryArea() *string {
	return s.DeliveryArea
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetDeliveryCity() *string {
	return s.DeliveryCity
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetDeliveryProvince() *string {
	return s.DeliveryProvince
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetDeliveryStreet() *string {
	return s.DeliveryStreet
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetEmail() *string {
	return s.Email
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetInvoiceMaterial() *int32 {
	return s.InvoiceMaterial
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetReceiverPhone() *string {
	return s.ReceiverPhone
}

func (s *HotelOrderCreateRequestInvoiceInfo) GetRemark() *string {
	return s.Remark
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetBuyerAdd(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.BuyerAdd = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetBuyerBankAcc(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.BuyerBankAcc = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetBuyerBankAdd(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.BuyerBankAdd = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetBuyerPhone(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.BuyerPhone = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetBuyerTaxNum(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.BuyerTaxNum = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetDeliveryAddress(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.DeliveryAddress = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetDeliveryArea(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.DeliveryArea = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetDeliveryCity(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.DeliveryCity = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetDeliveryProvince(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.DeliveryProvince = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetDeliveryStreet(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.DeliveryStreet = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetEmail(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.Email = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetInvoiceMaterial(v int32) *HotelOrderCreateRequestInvoiceInfo {
	s.InvoiceMaterial = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetInvoiceTitle(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.InvoiceTitle = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetInvoiceType(v int32) *HotelOrderCreateRequestInvoiceInfo {
	s.InvoiceType = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetReceiverName(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.ReceiverName = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetReceiverPhone(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.ReceiverPhone = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) SetRemark(v string) *HotelOrderCreateRequestInvoiceInfo {
	s.Remark = &v
	return s
}

func (s *HotelOrderCreateRequestInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderCreateRequestOccupantInfoList struct {
	// example:
	//
	// 124231213
	CardNo *string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// example:
	//
	// 1
	CardType *int32 `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// example:
	//
	// 1
	CustomerType *int32 `json:"customer_type,omitempty" xml:"customer_type,omitempty"`
	// example:
	//
	// 123112
	DepartmentId   *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DepartmentName *string `json:"department_name,omitempty" xml:"department_name,omitempty"`
	// example:
	//
	// demo
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1
	EmployeeType *int32 `json:"employee_type,omitempty" xml:"employee_type,omitempty"`
	// example:
	//
	// san
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// zhang
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18277123451
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RoomNo *int32  `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// example:
	//
	// 87817182
	StaffNo *string `json:"staff_no,omitempty" xml:"staff_no,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s HotelOrderCreateRequestOccupantInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateRequestOccupantInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetCardNo() *string {
	return s.CardNo
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetCardType() *int32 {
	return s.CardType
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetCustomerType() *int32 {
	return s.CustomerType
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetEmail() *string {
	return s.Email
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetEmployeeType() *int32 {
	return s.EmployeeType
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetFirstName() *string {
	return s.FirstName
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetLastName() *string {
	return s.LastName
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetName() *string {
	return s.Name
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetPhone() *string {
	return s.Phone
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetRoomNo() *int32 {
	return s.RoomNo
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetStaffNo() *string {
	return s.StaffNo
}

func (s *HotelOrderCreateRequestOccupantInfoList) GetUserType() *int32 {
	return s.UserType
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetCardNo(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.CardNo = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetCardType(v int32) *HotelOrderCreateRequestOccupantInfoList {
	s.CardType = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetCustomerType(v int32) *HotelOrderCreateRequestOccupantInfoList {
	s.CustomerType = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetDepartmentId(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.DepartmentId = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetDepartmentName(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.DepartmentName = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetEmail(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.Email = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetEmployeeType(v int32) *HotelOrderCreateRequestOccupantInfoList {
	s.EmployeeType = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetFirstName(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.FirstName = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetLastName(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.LastName = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetName(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.Name = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetPhone(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.Phone = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetRoomNo(v int32) *HotelOrderCreateRequestOccupantInfoList {
	s.RoomNo = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetStaffNo(v string) *HotelOrderCreateRequestOccupantInfoList {
	s.StaffNo = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) SetUserType(v int32) *HotelOrderCreateRequestOccupantInfoList {
	s.UserType = &v
	return s
}

func (s *HotelOrderCreateRequestOccupantInfoList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderCreateRequestPromotionInfo struct {
	PromotionDetailInfoList []*HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList `json:"promotion_detail_info_list,omitempty" xml:"promotion_detail_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	PromotionTotalPrice *int64 `json:"promotion_total_price,omitempty" xml:"promotion_total_price,omitempty"`
}

func (s HotelOrderCreateRequestPromotionInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateRequestPromotionInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateRequestPromotionInfo) GetPromotionDetailInfoList() []*HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	return s.PromotionDetailInfoList
}

func (s *HotelOrderCreateRequestPromotionInfo) GetPromotionTotalPrice() *int64 {
	return s.PromotionTotalPrice
}

func (s *HotelOrderCreateRequestPromotionInfo) SetPromotionDetailInfoList(v []*HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) *HotelOrderCreateRequestPromotionInfo {
	s.PromotionDetailInfoList = v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfo) SetPromotionTotalPrice(v int64) *HotelOrderCreateRequestPromotionInfo {
	s.PromotionTotalPrice = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList struct {
	// example:
	//
	// true
	CheckStatus *bool `json:"check_status,omitempty" xml:"check_status,omitempty"`
	// example:
	//
	// true
	NeedCheck     *bool   `json:"need_check,omitempty" xml:"need_check,omitempty"`
	PromotionCode *string `json:"promotion_code,omitempty" xml:"promotion_code,omitempty"`
	// example:
	//
	// 23778127
	PromotionId   *string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	PromotionName *string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// example:
	//
	// 100
	PromotionPrice *int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// example:
	//
	// 1
	PromotionType *string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
}

func (s HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetCheckStatus() *bool {
	return s.CheckStatus
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetNeedCheck() *bool {
	return s.NeedCheck
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetPromotionId() *string {
	return s.PromotionId
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetPromotionName() *string {
	return s.PromotionName
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetPromotionPrice() *int64 {
	return s.PromotionPrice
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) GetPromotionType() *string {
	return s.PromotionType
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetCheckStatus(v bool) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.CheckStatus = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetNeedCheck(v bool) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.NeedCheck = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetPromotionCode(v string) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.PromotionCode = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetPromotionId(v string) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.PromotionId = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetPromotionName(v string) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.PromotionName = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetPromotionPrice(v int64) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.PromotionPrice = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) SetPromotionType(v string) *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList {
	s.PromotionType = &v
	return s
}

func (s *HotelOrderCreateRequestPromotionInfoPromotionDetailInfoList) Validate() error {
	return dara.Validate(s)
}
