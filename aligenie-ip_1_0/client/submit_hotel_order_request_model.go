// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *SubmitHotelOrderRequestPayload) *SubmitHotelOrderRequest
	GetPayload() *SubmitHotelOrderRequestPayload
	SetUserInfo(v *SubmitHotelOrderRequestUserInfo) *SubmitHotelOrderRequest
	GetUserInfo() *SubmitHotelOrderRequestUserInfo
}

type SubmitHotelOrderRequest struct {
	// This parameter is required.
	Payload *SubmitHotelOrderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *SubmitHotelOrderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SubmitHotelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequest) GetPayload() *SubmitHotelOrderRequestPayload {
	return s.Payload
}

func (s *SubmitHotelOrderRequest) GetUserInfo() *SubmitHotelOrderRequestUserInfo {
	return s.UserInfo
}

func (s *SubmitHotelOrderRequest) SetPayload(v *SubmitHotelOrderRequestPayload) *SubmitHotelOrderRequest {
	s.Payload = v
	return s
}

func (s *SubmitHotelOrderRequest) SetUserInfo(v *SubmitHotelOrderRequestUserInfo) *SubmitHotelOrderRequest {
	s.UserInfo = v
	return s
}

func (s *SubmitHotelOrderRequest) Validate() error {
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitHotelOrderRequestPayload struct {
	// This parameter is required.
	ItemList []*SubmitHotelOrderRequestPayloadItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitHotelOrderRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderRequestPayload) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestPayload) GetItemList() []*SubmitHotelOrderRequestPayloadItemList {
	return s.ItemList
}

func (s *SubmitHotelOrderRequestPayload) GetType() *string {
	return s.Type
}

func (s *SubmitHotelOrderRequestPayload) SetItemList(v []*SubmitHotelOrderRequestPayloadItemList) *SubmitHotelOrderRequestPayload {
	s.ItemList = v
	return s
}

func (s *SubmitHotelOrderRequestPayload) SetType(v string) *SubmitHotelOrderRequestPayload {
	s.Type = &v
	return s
}

func (s *SubmitHotelOrderRequestPayload) Validate() error {
	if s.ItemList != nil {
		for _, item := range s.ItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitHotelOrderRequestPayloadItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// 152860
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Quantity *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s SubmitHotelOrderRequestPayloadItemList) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderRequestPayloadItemList) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestPayloadItemList) GetItemId() *int64 {
	return s.ItemId
}

func (s *SubmitHotelOrderRequestPayloadItemList) GetQuantity() *int64 {
	return s.Quantity
}

func (s *SubmitHotelOrderRequestPayloadItemList) GetRemark() *string {
	return s.Remark
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetItemId(v int64) *SubmitHotelOrderRequestPayloadItemList {
	s.ItemId = &v
	return s
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetQuantity(v int64) *SubmitHotelOrderRequestPayloadItemList {
	s.Quantity = &v
	return s
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetRemark(v string) *SubmitHotelOrderRequestPayloadItemList {
	s.Remark = &v
	return s
}

func (s *SubmitHotelOrderRequestPayloadItemList) Validate() error {
	return dara.Validate(s)
}

type SubmitHotelOrderRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SubmitHotelOrderRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SubmitHotelOrderRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SubmitHotelOrderRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *SubmitHotelOrderRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *SubmitHotelOrderRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SubmitHotelOrderRequestUserInfo) SetEncodeKey(v string) *SubmitHotelOrderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetEncodeType(v string) *SubmitHotelOrderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetId(v string) *SubmitHotelOrderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetIdType(v string) *SubmitHotelOrderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetOrganizationId(v string) *SubmitHotelOrderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
