// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDdosCredit(v *DescribeDdosCreditResponseBodyDdosCredit) *DescribeDdosCreditResponseBody
	GetDdosCredit() *DescribeDdosCreditResponseBodyDdosCredit
	SetRequestId(v string) *DescribeDdosCreditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDdosCreditResponseBody
	GetSuccess() *bool
}

type DescribeDdosCreditResponseBody struct {
	// The details of the security credit score of the current Alibaba Cloud account in the specified region.
	DdosCredit *DescribeDdosCreditResponseBodyDdosCredit `json:"DdosCredit,omitempty" xml:"DdosCredit,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E1F7BD73-8E9D-58D9-8658-CFC97112C641
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDdosCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCreditResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBody) GetDdosCredit() *DescribeDdosCreditResponseBodyDdosCredit {
	return s.DdosCredit
}

func (s *DescribeDdosCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosCreditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDdosCreditResponseBody) SetDdosCredit(v *DescribeDdosCreditResponseBodyDdosCredit) *DescribeDdosCreditResponseBody {
	s.DdosCredit = v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetRequestId(v string) *DescribeDdosCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetSuccess(v bool) *DescribeDdosCreditResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDdosCreditResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosCreditResponseBodyDdosCredit struct {
	// The time period after which blackhole filtering is automatically deactivated in the specified region. Unit: minutes.
	//
	// example:
	//
	// 150
	BlackholeTime *int32 `json:"BlackholeTime,omitempty" xml:"BlackholeTime,omitempty"`
	// The security credit score. The full score is **1000**.
	//
	// example:
	//
	// 550
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The security credit level. Valid values:
	//
	// 	- **A**: outstanding
	//
	// 	- **B**: excellent
	//
	// 	- **C**: good
	//
	// 	- **D**: average
	//
	// 	- **E**: poor
	//
	// 	- **F**: poorer
	//
	// example:
	//
	// D
	ScoreLevel *string `json:"ScoreLevel,omitempty" xml:"ScoreLevel,omitempty"`
}

func (s DescribeDdosCreditResponseBodyDdosCredit) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCreditResponseBodyDdosCredit) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) GetBlackholeTime() *int32 {
	return s.BlackholeTime
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) GetScore() *int32 {
	return s.Score
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) GetScoreLevel() *string {
	return s.ScoreLevel
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetBlackholeTime(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.BlackholeTime = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScore(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.Score = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScoreLevel(v string) *DescribeDdosCreditResponseBodyDdosCredit {
	s.ScoreLevel = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) Validate() error {
	return dara.Validate(s)
}
