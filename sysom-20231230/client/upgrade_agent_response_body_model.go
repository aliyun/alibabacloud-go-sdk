// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeAgentResponseBody
	GetRequestId() *string
	SetCode(v string) *UpgradeAgentResponseBody
	GetCode() *string
	SetData(v *UpgradeAgentResponseBodyData) *UpgradeAgentResponseBody
	GetData() *UpgradeAgentResponseBodyData
	SetMessage(v string) *UpgradeAgentResponseBody
	GetMessage() *string
}

type UpgradeAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpgradeAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpgradeAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeAgentResponseBody) GetData() *UpgradeAgentResponseBodyData {
	return s.Data
}

func (s *UpgradeAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeAgentResponseBody) SetRequestId(v string) *UpgradeAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAgentResponseBody) SetCode(v string) *UpgradeAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeAgentResponseBody) SetData(v *UpgradeAgentResponseBodyData) *UpgradeAgentResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeAgentResponseBody) SetMessage(v string) *UpgradeAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeAgentResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeAgentResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeAgentResponseBodyData) SetTaskId(v string) *UpgradeAgentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpgradeAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
