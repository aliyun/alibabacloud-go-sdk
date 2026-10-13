// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumbers(v []*CreateNumberRequestNumbers) *CreateNumberRequest
	GetNumbers() []*CreateNumberRequestNumbers
	SetTrunkId(v string) *CreateNumberRequest
	GetTrunkId() *string
}

type CreateNumberRequest struct {
	// The list of phone numbers to create and associate. A maximum of 100 phone numbers can be specified in a single request.
	//
	// This parameter is required.
	Numbers []*CreateNumberRequestNumbers `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// The unique identifier of the SIP trunk.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6bf8a1d2-3c4e-4f5a-9b8c-7d6e5f4a3b2c
	TrunkId *string `json:"TrunkId,omitempty" xml:"TrunkId,omitempty"`
}

func (s CreateNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNumberRequest) GoString() string {
	return s.String()
}

func (s *CreateNumberRequest) GetNumbers() []*CreateNumberRequestNumbers {
	return s.Numbers
}

func (s *CreateNumberRequest) GetTrunkId() *string {
	return s.TrunkId
}

func (s *CreateNumberRequest) SetNumbers(v []*CreateNumberRequestNumbers) *CreateNumberRequest {
	s.Numbers = v
	return s
}

func (s *CreateNumberRequest) SetTrunkId(v string) *CreateNumberRequest {
	s.TrunkId = &v
	return s
}

func (s *CreateNumberRequest) Validate() error {
	if s.Numbers != nil {
		for _, item := range s.Numbers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNumberRequestNumbers struct {
	// The phone number.
	//
	// example:
	//
	// 057188886666
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The maximum number of requests allowed within the throttling period.
	//
	// example:
	//
	// 100
	RateLimitCount *int32 `json:"RateLimitCount,omitempty" xml:"RateLimitCount,omitempty"`
	// The throttling period, in seconds.
	//
	// example:
	//
	// 60
	RateLimitPeriod *int32 `json:"RateLimitPeriod,omitempty" xml:"RateLimitPeriod,omitempty"`
}

func (s CreateNumberRequestNumbers) String() string {
	return dara.Prettify(s)
}

func (s CreateNumberRequestNumbers) GoString() string {
	return s.String()
}

func (s *CreateNumberRequestNumbers) GetNumber() *string {
	return s.Number
}

func (s *CreateNumberRequestNumbers) GetRateLimitCount() *int32 {
	return s.RateLimitCount
}

func (s *CreateNumberRequestNumbers) GetRateLimitPeriod() *int32 {
	return s.RateLimitPeriod
}

func (s *CreateNumberRequestNumbers) SetNumber(v string) *CreateNumberRequestNumbers {
	s.Number = &v
	return s
}

func (s *CreateNumberRequestNumbers) SetRateLimitCount(v int32) *CreateNumberRequestNumbers {
	s.RateLimitCount = &v
	return s
}

func (s *CreateNumberRequestNumbers) SetRateLimitPeriod(v int32) *CreateNumberRequestNumbers {
	s.RateLimitPeriod = &v
	return s
}

func (s *CreateNumberRequestNumbers) Validate() error {
	return dara.Validate(s)
}
