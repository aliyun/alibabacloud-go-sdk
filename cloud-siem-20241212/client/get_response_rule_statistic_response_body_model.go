// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResponseRuleStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResponseRuleStatisticResponseBody
	GetRequestId() *string
	SetResponseStatistic(v *GetResponseRuleStatisticResponseBodyResponseStatistic) *GetResponseRuleStatisticResponseBody
	GetResponseStatistic() *GetResponseRuleStatisticResponseBodyResponseStatistic
}

type GetResponseRuleStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of automated response rules.
	ResponseStatistic *GetResponseRuleStatisticResponseBodyResponseStatistic `json:"ResponseStatistic,omitempty" xml:"ResponseStatistic,omitempty" type:"Struct"`
}

func (s GetResponseRuleStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResponseRuleStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetResponseRuleStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResponseRuleStatisticResponseBody) GetResponseStatistic() *GetResponseRuleStatisticResponseBodyResponseStatistic {
	return s.ResponseStatistic
}

func (s *GetResponseRuleStatisticResponseBody) SetRequestId(v string) *GetResponseRuleStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResponseRuleStatisticResponseBody) SetResponseStatistic(v *GetResponseRuleStatisticResponseBodyResponseStatistic) *GetResponseRuleStatisticResponseBody {
	s.ResponseStatistic = v
	return s
}

func (s *GetResponseRuleStatisticResponseBody) Validate() error {
	if s.ResponseStatistic != nil {
		if err := s.ResponseStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResponseRuleStatisticResponseBodyResponseStatistic struct {
	// The total number of automated response rules.
	//
	// example:
	//
	// 2
	ResponseRuleAllCount *int32 `json:"ResponseRuleAllCount,omitempty" xml:"ResponseRuleAllCount,omitempty"`
	// The number of online automated response rules.
	//
	// example:
	//
	// 1
	ResponseRuleOnlineCount *int32 `json:"ResponseRuleOnlineCount,omitempty" xml:"ResponseRuleOnlineCount,omitempty"`
}

func (s GetResponseRuleStatisticResponseBodyResponseStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetResponseRuleStatisticResponseBodyResponseStatistic) GoString() string {
	return s.String()
}

func (s *GetResponseRuleStatisticResponseBodyResponseStatistic) GetResponseRuleAllCount() *int32 {
	return s.ResponseRuleAllCount
}

func (s *GetResponseRuleStatisticResponseBodyResponseStatistic) GetResponseRuleOnlineCount() *int32 {
	return s.ResponseRuleOnlineCount
}

func (s *GetResponseRuleStatisticResponseBodyResponseStatistic) SetResponseRuleAllCount(v int32) *GetResponseRuleStatisticResponseBodyResponseStatistic {
	s.ResponseRuleAllCount = &v
	return s
}

func (s *GetResponseRuleStatisticResponseBodyResponseStatistic) SetResponseRuleOnlineCount(v int32) *GetResponseRuleStatisticResponseBodyResponseStatistic {
	s.ResponseRuleOnlineCount = &v
	return s
}

func (s *GetResponseRuleStatisticResponseBodyResponseStatistic) Validate() error {
	return dara.Validate(s)
}
