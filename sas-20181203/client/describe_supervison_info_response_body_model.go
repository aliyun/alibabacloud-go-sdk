// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupervisonInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatestScanTime(v int64) *DescribeSupervisonInfoResponseBody
	GetLatestScanTime() *int64
	SetRequestId(v string) *DescribeSupervisonInfoResponseBody
	GetRequestId() *string
}

type DescribeSupervisonInfoResponseBody struct {
	// The time of the last system vulnerability scan. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1721134553000
	LatestScanTime *int64 `json:"LatestScanTime,omitempty" xml:"LatestScanTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E10BAF1C-A6C5-51E2-866C-76D5922E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupervisonInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupervisonInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupervisonInfoResponseBody) GetLatestScanTime() *int64 {
	return s.LatestScanTime
}

func (s *DescribeSupervisonInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupervisonInfoResponseBody) SetLatestScanTime(v int64) *DescribeSupervisonInfoResponseBody {
	s.LatestScanTime = &v
	return s
}

func (s *DescribeSupervisonInfoResponseBody) SetRequestId(v string) *DescribeSupervisonInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupervisonInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
