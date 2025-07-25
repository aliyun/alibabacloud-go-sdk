// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportLinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordLines(v *DescribeSupportLinesResponseBodyRecordLines) *DescribeSupportLinesResponseBody
	GetRecordLines() *DescribeSupportLinesResponseBodyRecordLines
	SetRequestId(v string) *DescribeSupportLinesResponseBody
	GetRequestId() *string
}

type DescribeSupportLinesResponseBody struct {
	// The Alibaba Cloud DNS lines.
	RecordLines *DescribeSupportLinesResponseBodyRecordLines `json:"RecordLines,omitempty" xml:"RecordLines,omitempty" type:"Struct"`
	// example:
	//
	// CFDA0830-7D6E-4C13-8632-B57C7EDCF079
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupportLinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponseBody) GetRecordLines() *DescribeSupportLinesResponseBodyRecordLines {
	return s.RecordLines
}

func (s *DescribeSupportLinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportLinesResponseBody) SetRecordLines(v *DescribeSupportLinesResponseBodyRecordLines) *DescribeSupportLinesResponseBody {
	s.RecordLines = v
	return s
}

func (s *DescribeSupportLinesResponseBody) SetRequestId(v string) *DescribeSupportLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportLinesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSupportLinesResponseBodyRecordLines struct {
	RecordLine []*DescribeSupportLinesResponseBodyRecordLinesRecordLine `json:"RecordLine,omitempty" xml:"RecordLine,omitempty" type:"Repeated"`
}

func (s DescribeSupportLinesResponseBodyRecordLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportLinesResponseBodyRecordLines) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponseBodyRecordLines) GetRecordLine() []*DescribeSupportLinesResponseBodyRecordLinesRecordLine {
	return s.RecordLine
}

func (s *DescribeSupportLinesResponseBodyRecordLines) SetRecordLine(v []*DescribeSupportLinesResponseBodyRecordLinesRecordLine) *DescribeSupportLinesResponseBodyRecordLines {
	s.RecordLine = v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLines) Validate() error {
	return dara.Validate(s)
}

type DescribeSupportLinesResponseBodyRecordLinesRecordLine struct {
	// The code of the parent line. Currently, no data is returned.
	//
	// example:
	//
	// unicom
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	// The code of the child line.
	//
	// example:
	//
	// cn_unicom_shanxi
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The display name of the line.
	//
	// example:
	//
	// China Unicom
	LineDisplayName *string `json:"LineDisplayName,omitempty" xml:"LineDisplayName,omitempty"`
	// The name of the child line.
	//
	// example:
	//
	// China Unicom_Shanxi
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeSupportLinesResponseBodyRecordLinesRecordLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportLinesResponseBodyRecordLinesRecordLine) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) GetFatherCode() *string {
	return s.FatherCode
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) GetLineDisplayName() *string {
	return s.LineDisplayName
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) SetFatherCode(v string) *DescribeSupportLinesResponseBodyRecordLinesRecordLine {
	s.FatherCode = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) SetLineCode(v string) *DescribeSupportLinesResponseBodyRecordLinesRecordLine {
	s.LineCode = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) SetLineDisplayName(v string) *DescribeSupportLinesResponseBodyRecordLinesRecordLine {
	s.LineDisplayName = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) SetLineName(v string) *DescribeSupportLinesResponseBodyRecordLinesRecordLine {
	s.LineName = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLinesRecordLine) Validate() error {
	return dara.Validate(s)
}
