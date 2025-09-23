// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApplicationVariablesResponseBody
	GetCode() *string
	SetData(v []*GetApplicationVariablesResponseBodyData) *GetApplicationVariablesResponseBody
	GetData() []*GetApplicationVariablesResponseBodyData
	SetMessage(v string) *GetApplicationVariablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationVariablesResponseBody
	GetRequestId() *string
}

type GetApplicationVariablesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetApplicationVariablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7036DDBE-0ABA-52D7-A39D-75E511970F07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApplicationVariablesResponseBody) GetData() []*GetApplicationVariablesResponseBodyData {
	return s.Data
}

func (s *GetApplicationVariablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationVariablesResponseBody) SetCode(v string) *GetApplicationVariablesResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationVariablesResponseBody) SetData(v []*GetApplicationVariablesResponseBodyData) *GetApplicationVariablesResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationVariablesResponseBody) SetMessage(v string) *GetApplicationVariablesResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationVariablesResponseBody) SetRequestId(v string) *GetApplicationVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetApplicationVariablesResponseBodyData struct {
	// example:
	//
	// ob5epf79uv****
	InstanceId   *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VariableList []*GetApplicationVariablesResponseBodyDataVariableList `json:"VariableList,omitempty" xml:"VariableList,omitempty" type:"Repeated"`
}

func (s GetApplicationVariablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationVariablesResponseBodyData) GetVariableList() []*GetApplicationVariablesResponseBodyDataVariableList {
	return s.VariableList
}

func (s *GetApplicationVariablesResponseBodyData) SetInstanceId(v string) *GetApplicationVariablesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationVariablesResponseBodyData) SetVariableList(v []*GetApplicationVariablesResponseBodyDataVariableList) *GetApplicationVariablesResponseBodyData {
	s.VariableList = v
	return s
}

func (s *GetApplicationVariablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetApplicationVariablesResponseBodyDataVariableList struct {
	// example:
	//
	// ecs.c6.4xlarge
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// ${instance_type}
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
}

func (s GetApplicationVariablesResponseBodyDataVariableList) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariablesResponseBodyDataVariableList) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) GetValue() *string {
	return s.Value
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) GetVariable() *string {
	return s.Variable
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) SetValue(v string) *GetApplicationVariablesResponseBodyDataVariableList {
	s.Value = &v
	return s
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) SetVariable(v string) *GetApplicationVariablesResponseBodyDataVariableList {
	s.Variable = &v
	return s
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) Validate() error {
	return dara.Validate(s)
}
