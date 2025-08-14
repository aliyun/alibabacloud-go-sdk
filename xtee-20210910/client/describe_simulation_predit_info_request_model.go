// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationPreditInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSimulationPreditInfoRequest
	GetLang() *string
	SetEventCode(v string) *DescribeSimulationPreditInfoRequest
	GetEventCode() *string
	SetRegId(v string) *DescribeSimulationPreditInfoRequest
	GetRegId() *string
	SetRulesStr(v string) *DescribeSimulationPreditInfoRequest
	GetRulesStr() *string
}

type DescribeSimulationPreditInfoRequest struct {
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
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_azywkh7523
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy list
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"100244\\",\\"100245\\"]
	RulesStr *string `json:"rulesStr,omitempty" xml:"rulesStr,omitempty"`
}

func (s DescribeSimulationPreditInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationPreditInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimulationPreditInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSimulationPreditInfoRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSimulationPreditInfoRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSimulationPreditInfoRequest) GetRulesStr() *string {
	return s.RulesStr
}

func (s *DescribeSimulationPreditInfoRequest) SetLang(v string) *DescribeSimulationPreditInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimulationPreditInfoRequest) SetEventCode(v string) *DescribeSimulationPreditInfoRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeSimulationPreditInfoRequest) SetRegId(v string) *DescribeSimulationPreditInfoRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSimulationPreditInfoRequest) SetRulesStr(v string) *DescribeSimulationPreditInfoRequest {
	s.RulesStr = &v
	return s
}

func (s *DescribeSimulationPreditInfoRequest) Validate() error {
	return dara.Validate(s)
}
