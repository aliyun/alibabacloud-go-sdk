// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealTimeConcurrencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxConcurrency(v int64) *GetRealTimeConcurrencyResponseBody
	GetMaxConcurrency() *int64
	SetRealTimeConcurrency(v int64) *GetRealTimeConcurrencyResponseBody
	GetRealTimeConcurrency() *int64
	SetRequestId(v string) *GetRealTimeConcurrencyResponseBody
	GetRequestId() *string
	SetTimestamp(v int64) *GetRealTimeConcurrencyResponseBody
	GetTimestamp() *int64
}

type GetRealTimeConcurrencyResponseBody struct {
	// example:
	//
	// 2
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// 1
	RealTimeConcurrency *int64 `json:"RealTimeConcurrency,omitempty" xml:"RealTimeConcurrency,omitempty"`
	// example:
	//
	// E6E61E1A-D2DC-5ACF-AED4-A115B6691F98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1661584255029
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetRealTimeConcurrencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealTimeConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealTimeConcurrencyResponseBody) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *GetRealTimeConcurrencyResponseBody) GetRealTimeConcurrency() *int64 {
	return s.RealTimeConcurrency
}

func (s *GetRealTimeConcurrencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealTimeConcurrencyResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetRealTimeConcurrencyResponseBody) SetMaxConcurrency(v int64) *GetRealTimeConcurrencyResponseBody {
	s.MaxConcurrency = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) SetRealTimeConcurrency(v int64) *GetRealTimeConcurrencyResponseBody {
	s.RealTimeConcurrency = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) SetRequestId(v string) *GetRealTimeConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) SetTimestamp(v int64) *GetRealTimeConcurrencyResponseBody {
	s.Timestamp = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) Validate() error {
	return dara.Validate(s)
}
