// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafRatelimitCharacteristics interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v []*WafRatelimitCharacteristicsCriteria) *WafRatelimitCharacteristics
	GetCriteria() []*WafRatelimitCharacteristicsCriteria
	SetLogic(v string) *WafRatelimitCharacteristics
	GetLogic() *string
	SetMatchType(v string) *WafRatelimitCharacteristics
	GetMatchType() *string
}

type WafRatelimitCharacteristics struct {
	Criteria  []*WafRatelimitCharacteristicsCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic     *string                                `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchType *string                                `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
}

func (s WafRatelimitCharacteristics) String() string {
	return dara.Prettify(s)
}

func (s WafRatelimitCharacteristics) GoString() string {
	return s.String()
}

func (s *WafRatelimitCharacteristics) GetCriteria() []*WafRatelimitCharacteristicsCriteria {
	return s.Criteria
}

func (s *WafRatelimitCharacteristics) GetLogic() *string {
	return s.Logic
}

func (s *WafRatelimitCharacteristics) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRatelimitCharacteristics) SetCriteria(v []*WafRatelimitCharacteristicsCriteria) *WafRatelimitCharacteristics {
	s.Criteria = v
	return s
}

func (s *WafRatelimitCharacteristics) SetLogic(v string) *WafRatelimitCharacteristics {
	s.Logic = &v
	return s
}

func (s *WafRatelimitCharacteristics) SetMatchType(v string) *WafRatelimitCharacteristics {
	s.MatchType = &v
	return s
}

func (s *WafRatelimitCharacteristics) Validate() error {
	if s.Criteria != nil {
		for _, item := range s.Criteria {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafRatelimitCharacteristicsCriteria struct {
	Criteria  []*WafRatelimitCharacteristicsCriteriaCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic     *string                                        `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchType *string                                        `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
}

func (s WafRatelimitCharacteristicsCriteria) String() string {
	return dara.Prettify(s)
}

func (s WafRatelimitCharacteristicsCriteria) GoString() string {
	return s.String()
}

func (s *WafRatelimitCharacteristicsCriteria) GetCriteria() []*WafRatelimitCharacteristicsCriteriaCriteria {
	return s.Criteria
}

func (s *WafRatelimitCharacteristicsCriteria) GetLogic() *string {
	return s.Logic
}

func (s *WafRatelimitCharacteristicsCriteria) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRatelimitCharacteristicsCriteria) SetCriteria(v []*WafRatelimitCharacteristicsCriteriaCriteria) *WafRatelimitCharacteristicsCriteria {
	s.Criteria = v
	return s
}

func (s *WafRatelimitCharacteristicsCriteria) SetLogic(v string) *WafRatelimitCharacteristicsCriteria {
	s.Logic = &v
	return s
}

func (s *WafRatelimitCharacteristicsCriteria) SetMatchType(v string) *WafRatelimitCharacteristicsCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRatelimitCharacteristicsCriteria) Validate() error {
	if s.Criteria != nil {
		for _, item := range s.Criteria {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafRatelimitCharacteristicsCriteriaCriteria struct {
	Criteria  []*WafRatelimitCharacteristicsCriteriaCriteriaCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic     *string                                                `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchType *string                                                `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
}

func (s WafRatelimitCharacteristicsCriteriaCriteria) String() string {
	return dara.Prettify(s)
}

func (s WafRatelimitCharacteristicsCriteriaCriteria) GoString() string {
	return s.String()
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) GetCriteria() []*WafRatelimitCharacteristicsCriteriaCriteriaCriteria {
	return s.Criteria
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) GetLogic() *string {
	return s.Logic
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) SetCriteria(v []*WafRatelimitCharacteristicsCriteriaCriteriaCriteria) *WafRatelimitCharacteristicsCriteriaCriteria {
	s.Criteria = v
	return s
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) SetLogic(v string) *WafRatelimitCharacteristicsCriteriaCriteria {
	s.Logic = &v
	return s
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) SetMatchType(v string) *WafRatelimitCharacteristicsCriteriaCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRatelimitCharacteristicsCriteriaCriteria) Validate() error {
	if s.Criteria != nil {
		for _, item := range s.Criteria {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafRatelimitCharacteristicsCriteriaCriteriaCriteria struct {
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
}

func (s WafRatelimitCharacteristicsCriteriaCriteriaCriteria) String() string {
	return dara.Prettify(s)
}

func (s WafRatelimitCharacteristicsCriteriaCriteriaCriteria) GoString() string {
	return s.String()
}

func (s *WafRatelimitCharacteristicsCriteriaCriteriaCriteria) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRatelimitCharacteristicsCriteriaCriteriaCriteria) SetMatchType(v string) *WafRatelimitCharacteristicsCriteriaCriteriaCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRatelimitCharacteristicsCriteriaCriteriaCriteria) Validate() error {
	return dara.Validate(s)
}
