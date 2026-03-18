// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnableMultiAzPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryEnableMultiAzPriceRequest
	GetInstanceId() *string
	SetObservers(v []*QueryEnableMultiAzPriceRequestObservers) *QueryEnableMultiAzPriceRequest
	GetObservers() []*QueryEnableMultiAzPriceRequestObservers
	SetPromotionOptionNo(v string) *QueryEnableMultiAzPriceRequest
	GetPromotionOptionNo() *string
}

type QueryEnableMultiAzPriceRequest struct {
	// example:
	//
	// c-8dsy12g*****
	InstanceId *string                                    `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Observers  []*QueryEnableMultiAzPriceRequestObservers `json:"observers,omitempty" xml:"observers,omitempty" type:"Repeated"`
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"promotionOptionNo,omitempty" xml:"promotionOptionNo,omitempty"`
}

func (s QueryEnableMultiAzPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryEnableMultiAzPriceRequest) GetObservers() []*QueryEnableMultiAzPriceRequestObservers {
	return s.Observers
}

func (s *QueryEnableMultiAzPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryEnableMultiAzPriceRequest) SetInstanceId(v string) *QueryEnableMultiAzPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryEnableMultiAzPriceRequest) SetObservers(v []*QueryEnableMultiAzPriceRequestObservers) *QueryEnableMultiAzPriceRequest {
	s.Observers = v
	return s
}

func (s *QueryEnableMultiAzPriceRequest) SetPromotionOptionNo(v string) *QueryEnableMultiAzPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryEnableMultiAzPriceRequest) Validate() error {
	if s.Observers != nil {
		for _, item := range s.Observers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryEnableMultiAzPriceRequestObservers struct {
	// example:
	//
	// vsw-9sdur12t27s
	VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
	// example:
	//
	// cn-shanghai-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s QueryEnableMultiAzPriceRequestObservers) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceRequestObservers) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceRequestObservers) GetVswId() *string {
	return s.VswId
}

func (s *QueryEnableMultiAzPriceRequestObservers) GetZoneId() *string {
	return s.ZoneId
}

func (s *QueryEnableMultiAzPriceRequestObservers) SetVswId(v string) *QueryEnableMultiAzPriceRequestObservers {
	s.VswId = &v
	return s
}

func (s *QueryEnableMultiAzPriceRequestObservers) SetZoneId(v string) *QueryEnableMultiAzPriceRequestObservers {
	s.ZoneId = &v
	return s
}

func (s *QueryEnableMultiAzPriceRequestObservers) Validate() error {
	return dara.Validate(s)
}
