// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAllExpirationDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *SetAllExpirationDayRequest
	GetOwnerId() *int64
	SetUnifyExpireDay(v string) *SetAllExpirationDayRequest
	GetUnifyExpireDay() *string
}

type SetAllExpirationDayRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current unified expiration date settings only support unexpired ECS instances. Expired instances cannot be aligned to the unified expiration date. After the settings are complete, you can align the renewal time with the configured unified expiration time.
	//
	// Valid values: 1 to 28, representing the 1st to the 28th day of each month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	UnifyExpireDay *string `json:"UnifyExpireDay,omitempty" xml:"UnifyExpireDay,omitempty"`
}

func (s SetAllExpirationDayRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAllExpirationDayRequest) GoString() string {
	return s.String()
}

func (s *SetAllExpirationDayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetAllExpirationDayRequest) GetUnifyExpireDay() *string {
	return s.UnifyExpireDay
}

func (s *SetAllExpirationDayRequest) SetOwnerId(v int64) *SetAllExpirationDayRequest {
	s.OwnerId = &v
	return s
}

func (s *SetAllExpirationDayRequest) SetUnifyExpireDay(v string) *SetAllExpirationDayRequest {
	s.UnifyExpireDay = &v
	return s
}

func (s *SetAllExpirationDayRequest) Validate() error {
	return dara.Validate(s)
}
