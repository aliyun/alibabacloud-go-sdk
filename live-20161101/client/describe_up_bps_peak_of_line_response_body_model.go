// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpBpsPeakOfLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeUpBpsPeakOfLines(v *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) *DescribeUpBpsPeakOfLineResponseBody
	GetDescribeUpBpsPeakOfLines() *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines
	SetRequestId(v string) *DescribeUpBpsPeakOfLineResponseBody
	GetRequestId() *string
}

type DescribeUpBpsPeakOfLineResponseBody struct {
	DescribeUpBpsPeakOfLines *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines `json:"DescribeUpBpsPeakOfLines,omitempty" xml:"DescribeUpBpsPeakOfLines,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6CFDE7AB-571A-14EA-B072-989FF753****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpBpsPeakOfLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponseBody) GetDescribeUpBpsPeakOfLines() *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines {
	return s.DescribeUpBpsPeakOfLines
}

func (s *DescribeUpBpsPeakOfLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpBpsPeakOfLineResponseBody) SetDescribeUpBpsPeakOfLines(v *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) *DescribeUpBpsPeakOfLineResponseBody {
	s.DescribeUpBpsPeakOfLines = v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBody) SetRequestId(v string) *DescribeUpBpsPeakOfLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBody) Validate() error {
	if s.DescribeUpBpsPeakOfLines != nil {
		if err := s.DescribeUpBpsPeakOfLines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines struct {
	DescribeUpBpsPeakOfLine []*DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine `json:"DescribeUpBpsPeakOfLine,omitempty" xml:"DescribeUpBpsPeakOfLine,omitempty" type:"Repeated"`
}

func (s DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) GetDescribeUpBpsPeakOfLine() []*DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	return s.DescribeUpBpsPeakOfLine
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) SetDescribeUpBpsPeakOfLine(v []*DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines {
	s.DescribeUpBpsPeakOfLine = v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLines) Validate() error {
	if s.DescribeUpBpsPeakOfLine != nil {
		for _, item := range s.DescribeUpBpsPeakOfLine {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine struct {
	BandWidth *float32 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	PeakTime  *string  `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	QueryTime *string  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	StatName  *string  `json:"StatName,omitempty" xml:"StatName,omitempty"`
}

func (s DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) GetBandWidth() *float32 {
	return s.BandWidth
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) GetPeakTime() *string {
	return s.PeakTime
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) GetStatName() *string {
	return s.StatName
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetBandWidth(v float32) *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.BandWidth = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetPeakTime(v string) *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.PeakTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetQueryTime(v string) *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.QueryTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetStatName(v string) *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.StatName = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseBodyDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) Validate() error {
	return dara.Validate(s)
}
