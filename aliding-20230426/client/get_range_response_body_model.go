// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundColors(v [][]*GetRangeResponseBodyBackgroundColors) *GetRangeResponseBody
	GetBackgroundColors() [][]*GetRangeResponseBodyBackgroundColors
	SetDisplayValues(v [][]*string) *GetRangeResponseBody
	GetDisplayValues() [][]*string
	SetFormulas(v [][]*string) *GetRangeResponseBody
	GetFormulas() [][]*string
	SetHyperlinks(v [][]*GetRangeResponseBodyHyperlinks) *GetRangeResponseBody
	GetHyperlinks() [][]*GetRangeResponseBodyHyperlinks
	SetRequestId(v string) *GetRangeResponseBody
	GetRequestId() *string
	SetValues(v [][]interface{}) *GetRangeResponseBody
	GetValues() [][]interface{}
}

type GetRangeResponseBody struct {
	// example:
	//
	// []
	BackgroundColors [][]*GetRangeResponseBodyBackgroundColors `json:"backgroundColors,omitempty" xml:"backgroundColors,omitempty" type:"Repeated"`
	// example:
	//
	// []
	DisplayValues [][]*string `json:"displayValues,omitempty" xml:"displayValues,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Formulas [][]*string `json:"formulas,omitempty" xml:"formulas,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Hyperlinks [][]*GetRangeResponseBodyHyperlinks `json:"hyperlinks,omitempty" xml:"hyperlinks,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// []
	Values [][]interface{} `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRangeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRangeResponseBody) GetBackgroundColors() [][]*GetRangeResponseBodyBackgroundColors {
	return s.BackgroundColors
}

func (s *GetRangeResponseBody) GetDisplayValues() [][]*string {
	return s.DisplayValues
}

func (s *GetRangeResponseBody) GetFormulas() [][]*string {
	return s.Formulas
}

func (s *GetRangeResponseBody) GetHyperlinks() [][]*GetRangeResponseBodyHyperlinks {
	return s.Hyperlinks
}

func (s *GetRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRangeResponseBody) GetValues() [][]interface{} {
	return s.Values
}

func (s *GetRangeResponseBody) SetBackgroundColors(v [][]*GetRangeResponseBodyBackgroundColors) *GetRangeResponseBody {
	s.BackgroundColors = v
	return s
}

func (s *GetRangeResponseBody) SetDisplayValues(v [][]*string) *GetRangeResponseBody {
	s.DisplayValues = v
	return s
}

func (s *GetRangeResponseBody) SetFormulas(v [][]*string) *GetRangeResponseBody {
	s.Formulas = v
	return s
}

func (s *GetRangeResponseBody) SetHyperlinks(v [][]*GetRangeResponseBodyHyperlinks) *GetRangeResponseBody {
	s.Hyperlinks = v
	return s
}

func (s *GetRangeResponseBody) SetRequestId(v string) *GetRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRangeResponseBody) SetValues(v [][]interface{}) *GetRangeResponseBody {
	s.Values = v
	return s
}

func (s *GetRangeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRangeResponseBodyBackgroundColors struct {
	// red
	//
	// example:
	//
	// 0
	Red *int32 `json:"Red,omitempty" xml:"Red,omitempty"`
	// green
	//
	// example:
	//
	// 0
	Green *int32 `json:"Green,omitempty" xml:"Green,omitempty"`
	// blue
	//
	// example:
	//
	// 0
	Blue *int32 `json:"Blue,omitempty" xml:"Blue,omitempty"`
	// hexString
	//
	// example:
	//
	// #000000
	HexString *string `json:"HexString,omitempty" xml:"HexString,omitempty"`
}

func (s GetRangeResponseBodyBackgroundColors) String() string {
	return dara.Prettify(s)
}

func (s GetRangeResponseBodyBackgroundColors) GoString() string {
	return s.String()
}

func (s *GetRangeResponseBodyBackgroundColors) GetRed() *int32 {
	return s.Red
}

func (s *GetRangeResponseBodyBackgroundColors) GetGreen() *int32 {
	return s.Green
}

func (s *GetRangeResponseBodyBackgroundColors) GetBlue() *int32 {
	return s.Blue
}

func (s *GetRangeResponseBodyBackgroundColors) GetHexString() *string {
	return s.HexString
}

func (s *GetRangeResponseBodyBackgroundColors) SetRed(v int32) *GetRangeResponseBodyBackgroundColors {
	s.Red = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetGreen(v int32) *GetRangeResponseBodyBackgroundColors {
	s.Green = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetBlue(v int32) *GetRangeResponseBodyBackgroundColors {
	s.Blue = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetHexString(v string) *GetRangeResponseBodyBackgroundColors {
	s.HexString = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) Validate() error {
	return dara.Validate(s)
}

type GetRangeResponseBodyHyperlinks struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetRangeResponseBodyHyperlinks) String() string {
	return dara.Prettify(s)
}

func (s GetRangeResponseBodyHyperlinks) GoString() string {
	return s.String()
}

func (s *GetRangeResponseBodyHyperlinks) GetType() *string {
	return s.Type
}

func (s *GetRangeResponseBodyHyperlinks) GetLink() *string {
	return s.Link
}

func (s *GetRangeResponseBodyHyperlinks) GetText() *string {
	return s.Text
}

func (s *GetRangeResponseBodyHyperlinks) SetType(v string) *GetRangeResponseBodyHyperlinks {
	s.Type = &v
	return s
}

func (s *GetRangeResponseBodyHyperlinks) SetLink(v string) *GetRangeResponseBodyHyperlinks {
	s.Link = &v
	return s
}

func (s *GetRangeResponseBodyHyperlinks) SetText(v string) *GetRangeResponseBodyHyperlinks {
	s.Text = &v
	return s
}

func (s *GetRangeResponseBodyHyperlinks) Validate() error {
	return dara.Validate(s)
}
