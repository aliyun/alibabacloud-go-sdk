// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelDetailsByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *DescribeModelDetailsByIdRequest
	GetModelId() *string
	SetRegId(v string) *DescribeModelDetailsByIdRequest
	GetRegId() *string
}

type DescribeModelDetailsByIdRequest struct {
	// Model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20619
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s DescribeModelDetailsByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelDetailsByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelDetailsByIdRequest) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeModelDetailsByIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeModelDetailsByIdRequest) SetModelId(v string) *DescribeModelDetailsByIdRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeModelDetailsByIdRequest) SetRegId(v string) *DescribeModelDetailsByIdRequest {
	s.RegId = &v
	return s
}

func (s *DescribeModelDetailsByIdRequest) Validate() error {
	return dara.Validate(s)
}
