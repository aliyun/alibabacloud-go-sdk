// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSecFuncInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeDcdnSecFuncInfoResponseBodyContent) *DescribeDcdnSecFuncInfoResponseBody
	GetContent() []*DescribeDcdnSecFuncInfoResponseBodyContent
	SetDescription(v string) *DescribeDcdnSecFuncInfoResponseBody
	GetDescription() *string
	SetHttpStatus(v string) *DescribeDcdnSecFuncInfoResponseBody
	GetHttpStatus() *string
	SetRequestId(v string) *DescribeDcdnSecFuncInfoResponseBody
	GetRequestId() *string
	SetRetCode(v string) *DescribeDcdnSecFuncInfoResponseBody
	GetRetCode() *string
}

type DescribeDcdnSecFuncInfoResponseBody struct {
	// The parameters required by the code.
	Content []*DescribeDcdnSecFuncInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The description of HTTP responses.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatus *string `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 30A3A25A-86B3-4C1D-BAA8-12B8607A5CFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return value for HTTP requests. Valid values:
	//
	// 	- 0: OK.
	//
	// 	- Values other than 0: an error.
	//
	// example:
	//
	// 0
	RetCode *string `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DescribeDcdnSecFuncInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoResponseBody) GetContent() []*DescribeDcdnSecFuncInfoResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnSecFuncInfoResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnSecFuncInfoResponseBody) GetHttpStatus() *string {
	return s.HttpStatus
}

func (s *DescribeDcdnSecFuncInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSecFuncInfoResponseBody) GetRetCode() *string {
	return s.RetCode
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetContent(v []*DescribeDcdnSecFuncInfoResponseBodyContent) *DescribeDcdnSecFuncInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetDescription(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetHttpStatus(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.HttpStatus = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetRequestId(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetRetCode(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.RetCode = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSecFuncInfoResponseBodyContent struct {
	// The language (Chinese or English).
	//
	// example:
	//
	// ai_defense
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The options in the drop-down list.
	//
	// example:
	//
	// ai_defense
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnSecFuncInfoResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) GetLabel() *string {
	return s.Label
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) SetLabel(v string) *DescribeDcdnSecFuncInfoResponseBodyContent {
	s.Label = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) SetValue(v string) *DescribeDcdnSecFuncInfoResponseBodyContent {
	s.Value = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
