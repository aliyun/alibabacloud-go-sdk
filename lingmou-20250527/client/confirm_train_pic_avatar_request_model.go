// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmTrainPicAvatarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarId(v string) *ConfirmTrainPicAvatarRequest
	GetAvatarId() *string
	SetStatus(v string) *ConfirmTrainPicAvatarRequest
	GetStatus() *string
}

type ConfirmTrainPicAvatarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// M1_eTNYgO5lOys5g7ObvC_nw
	AvatarId *string `json:"avatarId,omitempty" xml:"avatarId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMER_CONFIRMED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ConfirmTrainPicAvatarRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTrainPicAvatarRequest) GoString() string {
	return s.String()
}

func (s *ConfirmTrainPicAvatarRequest) GetAvatarId() *string {
	return s.AvatarId
}

func (s *ConfirmTrainPicAvatarRequest) GetStatus() *string {
	return s.Status
}

func (s *ConfirmTrainPicAvatarRequest) SetAvatarId(v string) *ConfirmTrainPicAvatarRequest {
	s.AvatarId = &v
	return s
}

func (s *ConfirmTrainPicAvatarRequest) SetStatus(v string) *ConfirmTrainPicAvatarRequest {
	s.Status = &v
	return s
}

func (s *ConfirmTrainPicAvatarRequest) Validate() error {
	return dara.Validate(s)
}
