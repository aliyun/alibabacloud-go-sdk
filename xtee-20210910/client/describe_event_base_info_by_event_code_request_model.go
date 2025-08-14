// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventBaseInfoByEventCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventBaseInfoByEventCodeRequest
	GetLang() *string
	SetCreateType(v string) *DescribeEventBaseInfoByEventCodeRequest
	GetCreateType() *string
	SetEventCode(v string) *DescribeEventBaseInfoByEventCodeRequest
	GetEventCode() *string
	SetRegId(v string) *DescribeEventBaseInfoByEventCodeRequest
	GetRegId() *string
}

type DescribeEventBaseInfoByEventCodeRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
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
	// de_arcehq4370
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeEventBaseInfoByEventCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventBaseInfoByEventCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventBaseInfoByEventCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventBaseInfoByEventCodeRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeEventBaseInfoByEventCodeRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventBaseInfoByEventCodeRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventBaseInfoByEventCodeRequest) SetLang(v string) *DescribeEventBaseInfoByEventCodeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeRequest) SetCreateType(v string) *DescribeEventBaseInfoByEventCodeRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeRequest) SetEventCode(v string) *DescribeEventBaseInfoByEventCodeRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeRequest) SetRegId(v string) *DescribeEventBaseInfoByEventCodeRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeRequest) Validate() error {
	return dara.Validate(s)
}
