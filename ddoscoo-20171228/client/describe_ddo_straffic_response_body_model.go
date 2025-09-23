// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDDoSTrafficPoints(v []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) *DescribeDDoSTrafficResponseBody
	GetDDoSTrafficPoints() []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints
	SetDefenseInBytes(v int64) *DescribeDDoSTrafficResponseBody
	GetDefenseInBytes() *int64
	SetRequestId(v string) *DescribeDDoSTrafficResponseBody
	GetRequestId() *string
	SetSourceInBytes(v int64) *DescribeDDoSTrafficResponseBody
	GetSourceInBytes() *int64
}

type DescribeDDoSTrafficResponseBody struct {
	DDoSTrafficPoints []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints `json:"DDoSTrafficPoints,omitempty" xml:"DDoSTrafficPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 23482234
	DefenseInBytes *int64 `json:"DefenseInBytes,omitempty" xml:"DefenseInBytes,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 19284762
	SourceInBytes *int64 `json:"SourceInBytes,omitempty" xml:"SourceInBytes,omitempty"`
}

func (s DescribeDDoSTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseBody) GetDDoSTrafficPoints() []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	return s.DDoSTrafficPoints
}

func (s *DescribeDDoSTrafficResponseBody) GetDefenseInBytes() *int64 {
	return s.DefenseInBytes
}

func (s *DescribeDDoSTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSTrafficResponseBody) GetSourceInBytes() *int64 {
	return s.SourceInBytes
}

func (s *DescribeDDoSTrafficResponseBody) SetDDoSTrafficPoints(v []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) *DescribeDDoSTrafficResponseBody {
	s.DDoSTrafficPoints = v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) SetDefenseInBytes(v int64) *DescribeDDoSTrafficResponseBody {
	s.DefenseInBytes = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) SetRequestId(v string) *DescribeDDoSTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) SetSourceInBytes(v int64) *DescribeDDoSTrafficResponseBody {
	s.SourceInBytes = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDDoSTrafficResponseBodyDDoSTrafficPoints struct {
	// example:
	//
	// 129867
	DefenseMaxInBps *int64 `json:"DefenseMaxInBps,omitempty" xml:"DefenseMaxInBps,omitempty"`
	// example:
	//
	// 129867
	SourceMaxInBps *int64 `json:"SourceMaxInBps,omitempty" xml:"SourceMaxInBps,omitempty"`
	// example:
	//
	// 234082304
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) GetDefenseMaxInBps() *int64 {
	return s.DefenseMaxInBps
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) GetSourceMaxInBps() *int64 {
	return s.SourceMaxInBps
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetDefenseMaxInBps(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.DefenseMaxInBps = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetSourceMaxInBps(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.SourceMaxInBps = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetTime(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.Time = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) Validate() error {
	return dara.Validate(s)
}
