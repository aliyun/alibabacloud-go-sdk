// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveChildAccountAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *RemoveChildAccountAuthRequest
	GetAppKey() *string
	SetChildAccountName(v string) *RemoveChildAccountAuthRequest
	GetChildAccountName() *string
	SetHotelId(v string) *RemoveChildAccountAuthRequest
	GetHotelId() *string
	SetTbOpenId(v string) *RemoveChildAccountAuthRequest
	GetTbOpenId() *string
}

type RemoveChildAccountAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30**53
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tbxxxx
	ChildAccountName *string `json:"ChildAccountName,omitempty" xml:"ChildAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s RemoveChildAccountAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveChildAccountAuthRequest) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *RemoveChildAccountAuthRequest) GetChildAccountName() *string {
	return s.ChildAccountName
}

func (s *RemoveChildAccountAuthRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *RemoveChildAccountAuthRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *RemoveChildAccountAuthRequest) SetAppKey(v string) *RemoveChildAccountAuthRequest {
	s.AppKey = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) SetChildAccountName(v string) *RemoveChildAccountAuthRequest {
	s.ChildAccountName = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) SetHotelId(v string) *RemoveChildAccountAuthRequest {
	s.HotelId = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) SetTbOpenId(v string) *RemoveChildAccountAuthRequest {
	s.TbOpenId = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) Validate() error {
	return dara.Validate(s)
}
