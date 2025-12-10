// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmDefsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLatestTimestamp(v string) *GetAlgorithmDefsRequest
	GetLatestTimestamp() *string
	SetRangeEnd(v int32) *GetAlgorithmDefsRequest
	GetRangeEnd() *int32
	SetRangeStart(v int32) *GetAlgorithmDefsRequest
	GetRangeStart() *int32
	SetTimestamp(v string) *GetAlgorithmDefsRequest
	GetTimestamp() *string
}

type GetAlgorithmDefsRequest struct {
	// example:
	//
	// 1709950208
	LatestTimestamp *string `json:"LatestTimestamp,omitempty" xml:"LatestTimestamp,omitempty"`
	// example:
	//
	// 10
	RangeEnd *int32 `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	// example:
	//
	// 1
	RangeStart *int32 `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
	// example:
	//
	// 1709950208
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetAlgorithmDefsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmDefsRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefsRequest) GetLatestTimestamp() *string {
	return s.LatestTimestamp
}

func (s *GetAlgorithmDefsRequest) GetRangeEnd() *int32 {
	return s.RangeEnd
}

func (s *GetAlgorithmDefsRequest) GetRangeStart() *int32 {
	return s.RangeStart
}

func (s *GetAlgorithmDefsRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAlgorithmDefsRequest) SetLatestTimestamp(v string) *GetAlgorithmDefsRequest {
	s.LatestTimestamp = &v
	return s
}

func (s *GetAlgorithmDefsRequest) SetRangeEnd(v int32) *GetAlgorithmDefsRequest {
	s.RangeEnd = &v
	return s
}

func (s *GetAlgorithmDefsRequest) SetRangeStart(v int32) *GetAlgorithmDefsRequest {
	s.RangeStart = &v
	return s
}

func (s *GetAlgorithmDefsRequest) SetTimestamp(v string) *GetAlgorithmDefsRequest {
	s.Timestamp = &v
	return s
}

func (s *GetAlgorithmDefsRequest) Validate() error {
	return dara.Validate(s)
}
