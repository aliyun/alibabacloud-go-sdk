// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInputFeildCountByEventCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeInputFeildCountByEventCodeRequest
	GetLang() *string
	SetCreateType(v string) *DescribeInputFeildCountByEventCodeRequest
	GetCreateType() *string
	SetEventCode(v string) *DescribeInputFeildCountByEventCodeRequest
	GetEventCode() *string
	SetRegId(v string) *DescribeInputFeildCountByEventCodeRequest
	GetRegId() *string
}

type DescribeInputFeildCountByEventCodeRequest struct {
	// Set the language type for request and response, default value is **zh**. Values:
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
	// Event code
	//
	// example:
	//
	// de_ahqido8038
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeInputFeildCountByEventCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInputFeildCountByEventCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInputFeildCountByEventCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInputFeildCountByEventCodeRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeInputFeildCountByEventCodeRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeInputFeildCountByEventCodeRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeInputFeildCountByEventCodeRequest) SetLang(v string) *DescribeInputFeildCountByEventCodeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeRequest) SetCreateType(v string) *DescribeInputFeildCountByEventCodeRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeRequest) SetEventCode(v string) *DescribeInputFeildCountByEventCodeRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeRequest) SetRegId(v string) *DescribeInputFeildCountByEventCodeRequest {
	s.RegId = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeRequest) Validate() error {
	return dara.Validate(s)
}
