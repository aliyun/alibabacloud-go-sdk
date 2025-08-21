// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnatEntryId(v string) *DeleteSnatEntryRequest
	GetSnatEntryId() *string
}

type DeleteSnatEntryRequest struct {
	// The ID of the SNAT entry that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-5tfi6f8gds82mjmlofeym****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
}

func (s DeleteSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DeleteSnatEntryRequest) SetSnatEntryId(v string) *DeleteSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DeleteSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
