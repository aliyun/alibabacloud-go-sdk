// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMeasureDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *PutMeasureDataRequest
	GetBizType() *string
	SetData(v string) *PutMeasureDataRequest
	GetData() *string
	SetDataType(v string) *PutMeasureDataRequest
	GetDataType() *string
	SetEndTime(v string) *PutMeasureDataRequest
	GetEndTime() *string
	SetStartTime(v string) *PutMeasureDataRequest
	GetStartTime() *string
}

type PutMeasureDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.sp
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 1634019240000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1640400574804
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PutMeasureDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PutMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureDataRequest) GetBizType() *string {
	return s.BizType
}

func (s *PutMeasureDataRequest) GetData() *string {
	return s.Data
}

func (s *PutMeasureDataRequest) GetDataType() *string {
	return s.DataType
}

func (s *PutMeasureDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *PutMeasureDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PutMeasureDataRequest) SetBizType(v string) *PutMeasureDataRequest {
	s.BizType = &v
	return s
}

func (s *PutMeasureDataRequest) SetData(v string) *PutMeasureDataRequest {
	s.Data = &v
	return s
}

func (s *PutMeasureDataRequest) SetDataType(v string) *PutMeasureDataRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataRequest) SetEndTime(v string) *PutMeasureDataRequest {
	s.EndTime = &v
	return s
}

func (s *PutMeasureDataRequest) SetStartTime(v string) *PutMeasureDataRequest {
	s.StartTime = &v
	return s
}

func (s *PutMeasureDataRequest) Validate() error {
	return dara.Validate(s)
}
