// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachRCInstancesResponseBody
	GetRequestId() *string
	SetResponses(v []*AttachRCInstancesResponseBodyResponses) *AttachRCInstancesResponseBody
	GetResponses() []*AttachRCInstancesResponseBodyResponses
	SetTaskId(v string) *AttachRCInstancesResponseBody
	GetTaskId() *string
}

type AttachRCInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 847BA085-B377-4BFA-8267-F82345ECE1D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	Responses []*AttachRCInstancesResponseBodyResponses `json:"Responses,omitempty" xml:"Responses,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// 178968983
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AttachRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachRCInstancesResponseBody) GetResponses() []*AttachRCInstancesResponseBodyResponses {
	return s.Responses
}

func (s *AttachRCInstancesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *AttachRCInstancesResponseBody) SetRequestId(v string) *AttachRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachRCInstancesResponseBody) SetResponses(v []*AttachRCInstancesResponseBodyResponses) *AttachRCInstancesResponseBody {
	s.Responses = v
	return s
}

func (s *AttachRCInstancesResponseBody) SetTaskId(v string) *AttachRCInstancesResponseBody {
	s.TaskId = &v
	return s
}

func (s *AttachRCInstancesResponseBody) Validate() error {
	if s.Responses != nil {
		for _, item := range s.Responses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachRCInstancesResponseBodyResponses struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The node ID.
	//
	// example:
	//
	// rc-e2g521l55k038cr8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The message returned.
	//
	// >  If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachRCInstancesResponseBodyResponses) String() string {
	return dara.Prettify(s)
}

func (s AttachRCInstancesResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *AttachRCInstancesResponseBodyResponses) GetCode() *string {
	return s.Code
}

func (s *AttachRCInstancesResponseBodyResponses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachRCInstancesResponseBodyResponses) GetMessage() *string {
	return s.Message
}

func (s *AttachRCInstancesResponseBodyResponses) SetCode(v string) *AttachRCInstancesResponseBodyResponses {
	s.Code = &v
	return s
}

func (s *AttachRCInstancesResponseBodyResponses) SetInstanceId(v string) *AttachRCInstancesResponseBodyResponses {
	s.InstanceId = &v
	return s
}

func (s *AttachRCInstancesResponseBodyResponses) SetMessage(v string) *AttachRCInstancesResponseBodyResponses {
	s.Message = &v
	return s
}

func (s *AttachRCInstancesResponseBodyResponses) Validate() error {
	return dara.Validate(s)
}
