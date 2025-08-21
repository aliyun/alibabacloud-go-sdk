// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnatEntryId(v string) *DescribeSnatAttributeRequest
	GetSnatEntryId() *string
}

type DescribeSnatAttributeRequest struct {
	// The ID of the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-5tc08qfj5ecblfdn2rqr9****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
}

func (s DescribeSnatAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnatAttributeRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatAttributeRequest) SetSnatEntryId(v string) *DescribeSnatAttributeRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatAttributeRequest) Validate() error {
	return dara.Validate(s)
}
