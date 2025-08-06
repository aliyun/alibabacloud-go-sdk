// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMeasureReadyFlagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *PutMeasureReadyFlagRequest
	GetBizType() *string
	SetDataType(v string) *PutMeasureReadyFlagRequest
	GetDataType() *string
	SetEndTime(v string) *PutMeasureReadyFlagRequest
	GetEndTime() *string
	SetReadyFlag(v string) *PutMeasureReadyFlagRequest
	GetReadyFlag() *string
	SetStartTime(v string) *PutMeasureReadyFlagRequest
	GetStartTime() *string
}

type PutMeasureReadyFlagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.bookkeeping
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634784240000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReadyFlag *string `json:"ReadyFlag,omitempty" xml:"ReadyFlag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634969692175
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PutMeasureReadyFlagRequest) String() string {
	return dara.Prettify(s)
}

func (s PutMeasureReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagRequest) GetBizType() *string {
	return s.BizType
}

func (s *PutMeasureReadyFlagRequest) GetDataType() *string {
	return s.DataType
}

func (s *PutMeasureReadyFlagRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *PutMeasureReadyFlagRequest) GetReadyFlag() *string {
	return s.ReadyFlag
}

func (s *PutMeasureReadyFlagRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PutMeasureReadyFlagRequest) SetBizType(v string) *PutMeasureReadyFlagRequest {
	s.BizType = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetDataType(v string) *PutMeasureReadyFlagRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetEndTime(v string) *PutMeasureReadyFlagRequest {
	s.EndTime = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetReadyFlag(v string) *PutMeasureReadyFlagRequest {
	s.ReadyFlag = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetStartTime(v string) *PutMeasureReadyFlagRequest {
	s.StartTime = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) Validate() error {
	return dara.Validate(s)
}
