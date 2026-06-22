// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStrictEventNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventNameList(v []*string) *UpdateStrictEventNameRequest
	GetEventNameList() []*string
	SetLang(v string) *UpdateStrictEventNameRequest
	GetLang() *string
	SetOperator(v string) *UpdateStrictEventNameRequest
	GetOperator() *string
}

type UpdateStrictEventNameRequest struct {
	// The list of strict mode alerts to operate on. This is a full list. Strict mode alerts not included in this list will have the opposite action applied.
	//
	// > Call [DescribeStrictEventName](~~DescribeStrictEventName~~) to obtain the list of all strict mode alerts.
	//
	// > -.
	EventNameList []*string `json:"EventNameList,omitempty" xml:"EventNameList,omitempty" type:"Repeated"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operator for the rule action. Valid values:
	//
	// - *on*: enables alerting
	//
	// - *off*: disables alerting.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s UpdateStrictEventNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStrictEventNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateStrictEventNameRequest) GetEventNameList() []*string {
	return s.EventNameList
}

func (s *UpdateStrictEventNameRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateStrictEventNameRequest) GetOperator() *string {
	return s.Operator
}

func (s *UpdateStrictEventNameRequest) SetEventNameList(v []*string) *UpdateStrictEventNameRequest {
	s.EventNameList = v
	return s
}

func (s *UpdateStrictEventNameRequest) SetLang(v string) *UpdateStrictEventNameRequest {
	s.Lang = &v
	return s
}

func (s *UpdateStrictEventNameRequest) SetOperator(v string) *UpdateStrictEventNameRequest {
	s.Operator = &v
	return s
}

func (s *UpdateStrictEventNameRequest) Validate() error {
	return dara.Validate(s)
}
