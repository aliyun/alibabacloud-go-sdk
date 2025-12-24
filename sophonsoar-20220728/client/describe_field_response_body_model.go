// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v string) *DescribeFieldResponseBody
	GetFields() *string
	SetName(v string) *DescribeFieldResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeFieldResponseBody
	GetRequestId() *string
}

type DescribeFieldResponseBody struct {
	// The configuration content.
	//
	// example:
	//
	// ["ip","name","hostinfo","md5"]
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The name of the global configuration.
	//
	// example:
	//
	// soar_filed_tags
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCDE6498-83CC-50A1-8307-3D5A539C42F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldResponseBody) GetFields() *string {
	return s.Fields
}

func (s *DescribeFieldResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFieldResponseBody) SetFields(v string) *DescribeFieldResponseBody {
	s.Fields = &v
	return s
}

func (s *DescribeFieldResponseBody) SetName(v string) *DescribeFieldResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFieldResponseBody) SetRequestId(v string) *DescribeFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
