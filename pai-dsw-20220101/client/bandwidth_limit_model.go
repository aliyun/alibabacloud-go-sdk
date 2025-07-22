// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthLimit interface {
	dara.Model
	String() string
	GoString() string
	SetEgressRate(v string) *BandwidthLimit
	GetEgressRate() *string
	SetEgressWhitelists(v []*string) *BandwidthLimit
	GetEgressWhitelists() []*string
	SetIngressRate(v string) *BandwidthLimit
	GetIngressRate() *string
	SetIngressWhitelists(v []*string) *BandwidthLimit
	GetIngressWhitelists() []*string
}

type BandwidthLimit struct {
	EgressRate        *string   `json:"EgressRate,omitempty" xml:"EgressRate,omitempty"`
	EgressWhitelists  []*string `json:"EgressWhitelists,omitempty" xml:"EgressWhitelists,omitempty" type:"Repeated"`
	IngressRate       *string   `json:"IngressRate,omitempty" xml:"IngressRate,omitempty"`
	IngressWhitelists []*string `json:"IngressWhitelists,omitempty" xml:"IngressWhitelists,omitempty" type:"Repeated"`
}

func (s BandwidthLimit) String() string {
	return dara.Prettify(s)
}

func (s BandwidthLimit) GoString() string {
	return s.String()
}

func (s *BandwidthLimit) GetEgressRate() *string {
	return s.EgressRate
}

func (s *BandwidthLimit) GetEgressWhitelists() []*string {
	return s.EgressWhitelists
}

func (s *BandwidthLimit) GetIngressRate() *string {
	return s.IngressRate
}

func (s *BandwidthLimit) GetIngressWhitelists() []*string {
	return s.IngressWhitelists
}

func (s *BandwidthLimit) SetEgressRate(v string) *BandwidthLimit {
	s.EgressRate = &v
	return s
}

func (s *BandwidthLimit) SetEgressWhitelists(v []*string) *BandwidthLimit {
	s.EgressWhitelists = v
	return s
}

func (s *BandwidthLimit) SetIngressRate(v string) *BandwidthLimit {
	s.IngressRate = &v
	return s
}

func (s *BandwidthLimit) SetIngressWhitelists(v []*string) *BandwidthLimit {
	s.IngressWhitelists = v
	return s
}

func (s *BandwidthLimit) Validate() error {
	return dara.Validate(s)
}
