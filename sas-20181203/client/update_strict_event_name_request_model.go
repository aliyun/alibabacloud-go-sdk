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
	// List of strict alarms to be operated on. This list is a complete list, and any strict alarms not included in this list will have the opposite operation performed.
	//
	// > You can call [DescribeStrictEventName](~~DescribeStrictEventName~~) to get the list of all strict mode alarms.
	//
	// > -
	EventNameList []*string `json:"EventNameList,omitempty" xml:"EventNameList,omitempty" type:"Repeated"`
	// Sets the language type for requests and received messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Operation rule determination operator:
	//
	// - *on*: Turn on the alarm
	//
	// - *off*: Turn off the alarm
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
