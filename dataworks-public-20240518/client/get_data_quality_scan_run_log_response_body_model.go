// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRunLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogSegment(v *GetDataQualityScanRunLogResponseBodyLogSegment) *GetDataQualityScanRunLogResponseBody
	GetLogSegment() *GetDataQualityScanRunLogResponseBodyLogSegment
	SetRequestId(v string) *GetDataQualityScanRunLogResponseBody
	GetRequestId() *string
}

type GetDataQualityScanRunLogResponseBody struct {
	LogSegment *GetDataQualityScanRunLogResponseBodyLogSegment `json:"LogSegment,omitempty" xml:"LogSegment,omitempty" type:"Struct"`
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityScanRunLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunLogResponseBody) GetLogSegment() *GetDataQualityScanRunLogResponseBodyLogSegment {
	return s.LogSegment
}

func (s *GetDataQualityScanRunLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityScanRunLogResponseBody) SetLogSegment(v *GetDataQualityScanRunLogResponseBodyLogSegment) *GetDataQualityScanRunLogResponseBody {
	s.LogSegment = v
	return s
}

func (s *GetDataQualityScanRunLogResponseBody) SetRequestId(v string) *GetDataQualityScanRunLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityScanRunLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunLogResponseBodyLogSegment struct {
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// example:
	//
	// 512000
	NextOffset *int64 `json:"NextOffset,omitempty" xml:"NextOffset,omitempty"`
}

func (s GetDataQualityScanRunLogResponseBodyLogSegment) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunLogResponseBodyLogSegment) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunLogResponseBodyLogSegment) GetLog() *string {
	return s.Log
}

func (s *GetDataQualityScanRunLogResponseBodyLogSegment) GetNextOffset() *int64 {
	return s.NextOffset
}

func (s *GetDataQualityScanRunLogResponseBodyLogSegment) SetLog(v string) *GetDataQualityScanRunLogResponseBodyLogSegment {
	s.Log = &v
	return s
}

func (s *GetDataQualityScanRunLogResponseBodyLogSegment) SetNextOffset(v int64) *GetDataQualityScanRunLogResponseBodyLogSegment {
	s.NextOffset = &v
	return s
}

func (s *GetDataQualityScanRunLogResponseBodyLogSegment) Validate() error {
	return dara.Validate(s)
}
