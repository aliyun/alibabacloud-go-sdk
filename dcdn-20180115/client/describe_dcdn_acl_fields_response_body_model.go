// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnAclFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeDcdnAclFieldsResponseBodyContent) *DescribeDcdnAclFieldsResponseBody
	GetContent() []*DescribeDcdnAclFieldsResponseBodyContent
	SetRequestId(v string) *DescribeDcdnAclFieldsResponseBody
	GetRequestId() *string
}

type DescribeDcdnAclFieldsResponseBody struct {
	// The details about the rules.
	Content []*DescribeDcdnAclFieldsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30A3A25A-86B3-4C1D-BAA8-12B8607A5CFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnAclFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnAclFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsResponseBody) GetContent() []*DescribeDcdnAclFieldsResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnAclFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnAclFieldsResponseBody) SetContent(v []*DescribeDcdnAclFieldsResponseBodyContent) *DescribeDcdnAclFieldsResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnAclFieldsResponseBody) SetRequestId(v string) *DescribeDcdnAclFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnAclFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnAclFieldsResponseBodyContent struct {
	// The rules and policies that were configured. The JSON string is decoded.
	//
	// example:
	//
	// \\"fieldList\\":[{\\"name\\":\\"alert\\",\\"display\\":\\"observe\\",\\"tip\\":\\"mark the request in the log without blocking it\\"},{\\"name\\":\\"bypass\\",\\"display\\":\\"bypass\\",\\"tip\\":\\"bypass security modules\\"}],\\"module\\":[{\\"name\\":\\"cc\\",\\"display\\":\\"Rate Limit\\",\\"tip\\":\\"bypass Rate Limit\\"},{\\"name\\":\\"bot\\",\\"display\\":\\"Bot Traffic Management\\",\\"tip\\":\\"bypass Bot Manager\\"}]
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
}

func (s DescribeDcdnAclFieldsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnAclFieldsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsResponseBodyContent) GetFields() *string {
	return s.Fields
}

func (s *DescribeDcdnAclFieldsResponseBodyContent) SetFields(v string) *DescribeDcdnAclFieldsResponseBodyContent {
	s.Fields = &v
	return s
}

func (s *DescribeDcdnAclFieldsResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
