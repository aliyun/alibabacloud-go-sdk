// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListDialogueTemplateRequest
	GetHotelId() *string
}

type ListDialogueTemplateRequest struct {
	// hotelId
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListDialogueTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListDialogueTemplateRequest) SetHotelId(v string) *ListDialogueTemplateRequest {
	s.HotelId = &v
	return s
}

func (s *ListDialogueTemplateRequest) Validate() error {
	return dara.Validate(s)
}
