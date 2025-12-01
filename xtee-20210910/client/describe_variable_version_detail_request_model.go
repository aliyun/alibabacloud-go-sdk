// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableVersionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableVersionDetailRequest
	GetLang() *string
	SetObjectCode(v string) *DescribeVariableVersionDetailRequest
	GetObjectCode() *string
	SetObjectId(v int64) *DescribeVariableVersionDetailRequest
	GetObjectId() *int64
	SetRegId(v string) *DescribeVariableVersionDetailRequest
	GetRegId() *string
	SetType(v string) *DescribeVariableVersionDetailRequest
	GetType() *string
	SetVersion(v int64) *DescribeVariableVersionDetailRequest
	GetVersion() *int64
}

type DescribeVariableVersionDetailRequest struct {
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
	// Associated variable name.
	//
	// example:
	//
	// dHSi2zf5bb8
	ObjectCode *string `json:"objectCode,omitempty" xml:"objectCode,omitempty"`
	// Associated variable ID.
	//
	// example:
	//
	// 3434
	ObjectId *int64 `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Type.
	//
	// example:
	//
	// VELOCITY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Version.
	//
	// example:
	//
	// 2
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeVariableVersionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableVersionDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableVersionDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableVersionDetailRequest) GetObjectCode() *string {
	return s.ObjectCode
}

func (s *DescribeVariableVersionDetailRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *DescribeVariableVersionDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableVersionDetailRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVariableVersionDetailRequest) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeVariableVersionDetailRequest) SetLang(v string) *DescribeVariableVersionDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableVersionDetailRequest) SetObjectCode(v string) *DescribeVariableVersionDetailRequest {
	s.ObjectCode = &v
	return s
}

func (s *DescribeVariableVersionDetailRequest) SetObjectId(v int64) *DescribeVariableVersionDetailRequest {
	s.ObjectId = &v
	return s
}

func (s *DescribeVariableVersionDetailRequest) SetRegId(v string) *DescribeVariableVersionDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableVersionDetailRequest) SetType(v string) *DescribeVariableVersionDetailRequest {
	s.Type = &v
	return s
}

func (s *DescribeVariableVersionDetailRequest) SetVersion(v int64) *DescribeVariableVersionDetailRequest {
	s.Version = &v
	return s
}

func (s *DescribeVariableVersionDetailRequest) Validate() error {
	return dara.Validate(s)
}
