// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePunishedDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPunishedDomains(v []*string) *DescribePunishedDomainsResponseBody
	GetPunishedDomains() []*string
	SetRequestId(v string) *DescribePunishedDomainsResponseBody
	GetRequestId() *string
}

type DescribePunishedDomainsResponseBody struct {
	// The domain names that are penalized for failing to obtain an ICP filing.
	PunishedDomains []*string `json:"PunishedDomains,omitempty" xml:"PunishedDomains,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B1F4D802-55A1-5D53-A247-7E79****85E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePunishedDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePunishedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePunishedDomainsResponseBody) GetPunishedDomains() []*string {
	return s.PunishedDomains
}

func (s *DescribePunishedDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePunishedDomainsResponseBody) SetPunishedDomains(v []*string) *DescribePunishedDomainsResponseBody {
	s.PunishedDomains = v
	return s
}

func (s *DescribePunishedDomainsResponseBody) SetRequestId(v string) *DescribePunishedDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePunishedDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
