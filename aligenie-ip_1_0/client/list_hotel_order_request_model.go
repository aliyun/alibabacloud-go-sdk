// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *ListHotelOrderRequestPayload) *ListHotelOrderRequest
	GetPayload() *ListHotelOrderRequestPayload
	SetUserInfo(v *ListHotelOrderRequestUserInfo) *ListHotelOrderRequest
	GetUserInfo() *ListHotelOrderRequestUserInfo
}

type ListHotelOrderRequest struct {
	// This parameter is required.
	Payload *ListHotelOrderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListHotelOrderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderRequest) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequest) GetPayload() *ListHotelOrderRequestPayload {
	return s.Payload
}

func (s *ListHotelOrderRequest) GetUserInfo() *ListHotelOrderRequestUserInfo {
	return s.UserInfo
}

func (s *ListHotelOrderRequest) SetPayload(v *ListHotelOrderRequestPayload) *ListHotelOrderRequest {
	s.Payload = v
	return s
}

func (s *ListHotelOrderRequest) SetUserInfo(v *ListHotelOrderRequestUserInfo) *ListHotelOrderRequest {
	s.UserInfo = v
	return s
}

func (s *ListHotelOrderRequest) Validate() error {
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

type ListHotelOrderRequestPayload struct {
	// This parameter is required.
	Page *ListHotelOrderRequestPayloadPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListHotelOrderRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestPayload) GetPage() *ListHotelOrderRequestPayloadPage {
	return s.Page
}

func (s *ListHotelOrderRequestPayload) SetPage(v *ListHotelOrderRequestPayloadPage) *ListHotelOrderRequestPayload {
	s.Page = v
	return s
}

func (s *ListHotelOrderRequestPayload) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelOrderRequestPayloadPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelOrderRequestPayloadPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderRequestPayloadPage) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestPayloadPage) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListHotelOrderRequestPayloadPage) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListHotelOrderRequestPayloadPage) SetPageNumber(v int64) *ListHotelOrderRequestPayloadPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelOrderRequestPayloadPage) SetPageSize(v int64) *ListHotelOrderRequestPayloadPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelOrderRequestPayloadPage) Validate() error {
	return dara.Validate(s)
}

type ListHotelOrderRequestUserInfo struct {
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

func (s ListHotelOrderRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListHotelOrderRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListHotelOrderRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListHotelOrderRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListHotelOrderRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListHotelOrderRequestUserInfo) SetEncodeKey(v string) *ListHotelOrderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetEncodeType(v string) *ListHotelOrderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetId(v string) *ListHotelOrderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetIdType(v string) *ListHotelOrderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetOrganizationId(v string) *ListHotelOrderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
