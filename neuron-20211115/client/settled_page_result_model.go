// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSettledPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SettledPageResult
	GetRequestId() *string
	SetSettleds(v []*Settled) *SettledPageResult
	GetSettleds() []*Settled
}

type SettledPageResult struct {
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Settleds  []*Settled `json:"settleds,omitempty" xml:"settleds,omitempty" type:"Repeated"`
}

func (s SettledPageResult) String() string {
	return dara.Prettify(s)
}

func (s SettledPageResult) GoString() string {
	return s.String()
}

func (s *SettledPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *SettledPageResult) GetSettleds() []*Settled {
	return s.Settleds
}

func (s *SettledPageResult) SetRequestId(v string) *SettledPageResult {
	s.RequestId = &v
	return s
}

func (s *SettledPageResult) SetSettleds(v []*Settled) *SettledPageResult {
	s.Settleds = v
	return s
}

func (s *SettledPageResult) Validate() error {
	if s.Settleds != nil {
		for _, item := range s.Settleds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
