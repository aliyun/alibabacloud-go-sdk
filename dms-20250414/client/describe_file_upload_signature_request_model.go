// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *DescribeFileUploadSignatureRequest
	GetCallFrom() *string
	SetDmsUnit(v string) *DescribeFileUploadSignatureRequest
	GetDmsUnit() *string
}

type DescribeFileUploadSignatureRequest struct {
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
}

func (s DescribeFileUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignatureRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *DescribeFileUploadSignatureRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *DescribeFileUploadSignatureRequest) SetCallFrom(v string) *DescribeFileUploadSignatureRequest {
	s.CallFrom = &v
	return s
}

func (s *DescribeFileUploadSignatureRequest) SetDmsUnit(v string) *DescribeFileUploadSignatureRequest {
	s.DmsUnit = &v
	return s
}

func (s *DescribeFileUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
