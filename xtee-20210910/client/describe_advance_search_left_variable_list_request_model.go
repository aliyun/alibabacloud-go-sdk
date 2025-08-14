// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvanceSearchLeftVariableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAdvanceSearchLeftVariableListRequest
	GetLang() *string
	SetEventCodes(v string) *DescribeAdvanceSearchLeftVariableListRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeAdvanceSearchLeftVariableListRequest
	GetRegId() *string
	SetScene(v string) *DescribeAdvanceSearchLeftVariableListRequest
	GetScene() *string
}

type DescribeAdvanceSearchLeftVariableListRequest struct {
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
	// Event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ahqhsw7665,de_agbzfi5134
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Scene
	//
	// This parameter is required.
	//
	// example:
	//
	// VELOCITY
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s DescribeAdvanceSearchLeftVariableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchLeftVariableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) SetLang(v string) *DescribeAdvanceSearchLeftVariableListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) SetEventCodes(v string) *DescribeAdvanceSearchLeftVariableListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) SetRegId(v string) *DescribeAdvanceSearchLeftVariableListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) SetScene(v string) *DescribeAdvanceSearchLeftVariableListRequest {
	s.Scene = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListRequest) Validate() error {
	return dara.Validate(s)
}
