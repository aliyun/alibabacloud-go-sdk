// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableVersionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExpressionVariableVersionDetailRequest
	GetLang() *string
	SetObjectCode(v string) *DescribeExpressionVariableVersionDetailRequest
	GetObjectCode() *string
	SetObjectId(v int64) *DescribeExpressionVariableVersionDetailRequest
	GetObjectId() *int64
	SetRegId(v string) *DescribeExpressionVariableVersionDetailRequest
	GetRegId() *string
	SetType(v string) *DescribeExpressionVariableVersionDetailRequest
	GetType() *string
	SetVersion(v int64) *DescribeExpressionVariableVersionDetailRequest
	GetVersion() *int64
}

type DescribeExpressionVariableVersionDetailRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Associated variable code.
	//
	// example:
	//
	// ex_0kWIfZ27c525
	ObjectCode *string `json:"objectCode,omitempty" xml:"objectCode,omitempty"`
	// Associated variable primary key ID.
	//
	// example:
	//
	// 397625
	ObjectId *int64 `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Variable type.
	//
	// example:
	//
	// EXPRESSION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Version number.
	//
	// example:
	//
	// 2
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeExpressionVariableVersionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableVersionDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableVersionDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExpressionVariableVersionDetailRequest) GetObjectCode() *string {
	return s.ObjectCode
}

func (s *DescribeExpressionVariableVersionDetailRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *DescribeExpressionVariableVersionDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeExpressionVariableVersionDetailRequest) GetType() *string {
	return s.Type
}

func (s *DescribeExpressionVariableVersionDetailRequest) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeExpressionVariableVersionDetailRequest) SetLang(v string) *DescribeExpressionVariableVersionDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailRequest) SetObjectCode(v string) *DescribeExpressionVariableVersionDetailRequest {
	s.ObjectCode = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailRequest) SetObjectId(v int64) *DescribeExpressionVariableVersionDetailRequest {
	s.ObjectId = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailRequest) SetRegId(v string) *DescribeExpressionVariableVersionDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailRequest) SetType(v string) *DescribeExpressionVariableVersionDetailRequest {
	s.Type = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailRequest) SetVersion(v int64) *DescribeExpressionVariableVersionDetailRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailRequest) Validate() error {
	return dara.Validate(s)
}
