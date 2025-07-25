// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iACS interface {
	dara.Model
	String() string
	GoString() string
	SetACSQuotaId(v string) *ACS
	GetACSQuotaId() *string
	SetAssociatedProducts(v []*string) *ACS
	GetAssociatedProducts() []*string
}

type ACS struct {
	ACSQuotaId         *string   `json:"ACSQuotaId,omitempty" xml:"ACSQuotaId,omitempty"`
	AssociatedProducts []*string `json:"AssociatedProducts,omitempty" xml:"AssociatedProducts,omitempty" type:"Repeated"`
}

func (s ACS) String() string {
	return dara.Prettify(s)
}

func (s ACS) GoString() string {
	return s.String()
}

func (s *ACS) GetACSQuotaId() *string {
	return s.ACSQuotaId
}

func (s *ACS) GetAssociatedProducts() []*string {
	return s.AssociatedProducts
}

func (s *ACS) SetACSQuotaId(v string) *ACS {
	s.ACSQuotaId = &v
	return s
}

func (s *ACS) SetAssociatedProducts(v []*string) *ACS {
	s.AssociatedProducts = v
	return s
}

func (s *ACS) Validate() error {
	return dara.Validate(s)
}
