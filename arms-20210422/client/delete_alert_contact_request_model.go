// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *DeleteAlertContactRequest
	GetContactId() *int64
	SetRegionId(v string) *DeleteAlertContactRequest
	GetRegionId() *string
}

type DeleteAlertContactRequest struct {
	// This parameter is required.
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *DeleteAlertContactRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAlertContactRequest) SetContactId(v int64) *DeleteAlertContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteAlertContactRequest) SetRegionId(v string) *DeleteAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlertContactRequest) Validate() error {
	return dara.Validate(s)
}
