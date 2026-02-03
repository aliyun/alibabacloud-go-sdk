// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisTrafficRankingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNisTrafficRankingId(v string) *StartNisTrafficRankingResponseBody
	GetNisTrafficRankingId() *string
	SetRequestId(v string) *StartNisTrafficRankingResponseBody
	GetRequestId() *string
}

type StartNisTrafficRankingResponseBody struct {
	// example:
	//
	// task-6462a7b4c4a54b****
	NisTrafficRankingId *string `json:"NisTrafficRankingId,omitempty" xml:"NisTrafficRankingId,omitempty"`
	// example:
	//
	// 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartNisTrafficRankingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingResponseBody) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingResponseBody) GetNisTrafficRankingId() *string {
	return s.NisTrafficRankingId
}

func (s *StartNisTrafficRankingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartNisTrafficRankingResponseBody) SetNisTrafficRankingId(v string) *StartNisTrafficRankingResponseBody {
	s.NisTrafficRankingId = &v
	return s
}

func (s *StartNisTrafficRankingResponseBody) SetRequestId(v string) *StartNisTrafficRankingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartNisTrafficRankingResponseBody) Validate() error {
	return dara.Validate(s)
}
