// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRobotNodeResponseBody
	GetCode() *string
	SetData(v []*ListRobotNodeResponseBodyData) *ListRobotNodeResponseBody
	GetData() []*ListRobotNodeResponseBodyData
	SetMessage(v string) *ListRobotNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRobotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRobotNodeResponseBody
	GetSuccess() *bool
}

type ListRobotNodeResponseBody struct {
	// example:
	//
	// Ok
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListRobotNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRobotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRobotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRobotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRobotNodeResponseBody) GetData() []*ListRobotNodeResponseBodyData {
	return s.Data
}

func (s *ListRobotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRobotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRobotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRobotNodeResponseBody) SetCode(v string) *ListRobotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ListRobotNodeResponseBody) SetData(v []*ListRobotNodeResponseBodyData) *ListRobotNodeResponseBody {
	s.Data = v
	return s
}

func (s *ListRobotNodeResponseBody) SetMessage(v string) *ListRobotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *ListRobotNodeResponseBody) SetRequestId(v string) *ListRobotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRobotNodeResponseBody) SetSuccess(v bool) *ListRobotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *ListRobotNodeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRobotNodeResponseBodyData struct {
	// example:
	//
	// 1
	IsOutput  *int32  `json:"IsOutput,omitempty" xml:"IsOutput,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 123456
	NodeIdentifier *string `json:"NodeIdentifier,omitempty" xml:"NodeIdentifier,omitempty"`
	NodeName       *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	ProcessName    *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
}

func (s ListRobotNodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRobotNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRobotNodeResponseBodyData) GetIsOutput() *int32 {
	return s.IsOutput
}

func (s *ListRobotNodeResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *ListRobotNodeResponseBodyData) GetNodeIdentifier() *string {
	return s.NodeIdentifier
}

func (s *ListRobotNodeResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *ListRobotNodeResponseBodyData) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListRobotNodeResponseBodyData) SetIsOutput(v int32) *ListRobotNodeResponseBodyData {
	s.IsOutput = &v
	return s
}

func (s *ListRobotNodeResponseBodyData) SetModelName(v string) *ListRobotNodeResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *ListRobotNodeResponseBodyData) SetNodeIdentifier(v string) *ListRobotNodeResponseBodyData {
	s.NodeIdentifier = &v
	return s
}

func (s *ListRobotNodeResponseBodyData) SetNodeName(v string) *ListRobotNodeResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *ListRobotNodeResponseBodyData) SetProcessName(v string) *ListRobotNodeResponseBodyData {
	s.ProcessName = &v
	return s
}

func (s *ListRobotNodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
