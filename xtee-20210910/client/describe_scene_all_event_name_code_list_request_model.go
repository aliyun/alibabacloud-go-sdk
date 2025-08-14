// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneAllEventNameCodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSceneAllEventNameCodeListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeSceneAllEventNameCodeListRequest
	GetCreateType() *string
	SetRegId(v string) *DescribeSceneAllEventNameCodeListRequest
	GetRegId() *string
}

type DescribeSceneAllEventNameCodeListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSceneAllEventNameCodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneAllEventNameCodeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneAllEventNameCodeListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSceneAllEventNameCodeListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeSceneAllEventNameCodeListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSceneAllEventNameCodeListRequest) SetLang(v string) *DescribeSceneAllEventNameCodeListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListRequest) SetCreateType(v string) *DescribeSceneAllEventNameCodeListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListRequest) SetRegId(v string) *DescribeSceneAllEventNameCodeListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListRequest) Validate() error {
	return dara.Validate(s)
}
