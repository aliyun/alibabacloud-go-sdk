// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCopyRuleVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckCopyRuleVariableResponseBody
	GetRequestId() *string
	SetResultObject(v *CheckCopyRuleVariableResponseBodyResultObject) *CheckCopyRuleVariableResponseBody
	GetResultObject() *CheckCopyRuleVariableResponseBodyResultObject
}

type CheckCopyRuleVariableResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *CheckCopyRuleVariableResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CheckCopyRuleVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCopyRuleVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCopyRuleVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCopyRuleVariableResponseBody) GetResultObject() *CheckCopyRuleVariableResponseBodyResultObject {
	return s.ResultObject
}

func (s *CheckCopyRuleVariableResponseBody) SetRequestId(v string) *CheckCopyRuleVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCopyRuleVariableResponseBody) SetResultObject(v *CheckCopyRuleVariableResponseBodyResultObject) *CheckCopyRuleVariableResponseBody {
	s.ResultObject = v
	return s
}

func (s *CheckCopyRuleVariableResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckCopyRuleVariableResponseBodyResultObject struct {
	// Information.
	Message []*CheckCopyRuleVariableResponseBodyResultObjectMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CheckCopyRuleVariableResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CheckCopyRuleVariableResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CheckCopyRuleVariableResponseBodyResultObject) GetMessage() []*CheckCopyRuleVariableResponseBodyResultObjectMessage {
	return s.Message
}

func (s *CheckCopyRuleVariableResponseBodyResultObject) SetMessage(v []*CheckCopyRuleVariableResponseBodyResultObjectMessage) *CheckCopyRuleVariableResponseBodyResultObject {
	s.Message = v
	return s
}

func (s *CheckCopyRuleVariableResponseBodyResultObject) Validate() error {
	if s.Message != nil {
		for _, item := range s.Message {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckCopyRuleVariableResponseBodyResultObjectMessage struct {
	// Primary key ID of the variable
	//
	// example:
	//
	// 213
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name of the variable
	//
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Title of the variable
	//
	// example:
	//
	// 年龄
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Type of the variable
	//
	// example:
	//
	// NATIVE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckCopyRuleVariableResponseBodyResultObjectMessage) String() string {
	return dara.Prettify(s)
}

func (s CheckCopyRuleVariableResponseBodyResultObjectMessage) GoString() string {
	return s.String()
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) GetId() *int64 {
	return s.Id
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) GetName() *string {
	return s.Name
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) GetTitle() *string {
	return s.Title
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) GetType() *string {
	return s.Type
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) SetId(v int64) *CheckCopyRuleVariableResponseBodyResultObjectMessage {
	s.Id = &v
	return s
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) SetName(v string) *CheckCopyRuleVariableResponseBodyResultObjectMessage {
	s.Name = &v
	return s
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) SetTitle(v string) *CheckCopyRuleVariableResponseBodyResultObjectMessage {
	s.Title = &v
	return s
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) SetType(v string) *CheckCopyRuleVariableResponseBodyResultObjectMessage {
	s.Type = &v
	return s
}

func (s *CheckCopyRuleVariableResponseBodyResultObjectMessage) Validate() error {
	return dara.Validate(s)
}
