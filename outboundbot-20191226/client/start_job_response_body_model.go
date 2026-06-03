// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallIds(v []*StartJobResponseBodyCallIds) *StartJobResponseBody
	GetCallIds() []*StartJobResponseBodyCallIds
	SetCode(v string) *StartJobResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StartJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartJobResponseBody
	GetSuccess() *bool
	SetTaskIds(v []*StartJobResponseBodyTaskIds) *StartJobResponseBody
	GetTaskIds() []*StartJobResponseBodyTaskIds
}

type StartJobResponseBody struct {
	CallIds []*StartJobResponseBodyCallIds `json:"CallIds,omitempty" xml:"CallIds,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8a621aa1-d2e7-43f3-b54d-8830af73c468
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskIds []*StartJobResponseBodyTaskIds `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s StartJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobResponseBody) GetCallIds() []*StartJobResponseBodyCallIds {
	return s.CallIds
}

func (s *StartJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartJobResponseBody) GetTaskIds() []*StartJobResponseBodyTaskIds {
	return s.TaskIds
}

func (s *StartJobResponseBody) SetCallIds(v []*StartJobResponseBodyCallIds) *StartJobResponseBody {
	s.CallIds = v
	return s
}

func (s *StartJobResponseBody) SetCode(v string) *StartJobResponseBody {
	s.Code = &v
	return s
}

func (s *StartJobResponseBody) SetHttpStatusCode(v int32) *StartJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartJobResponseBody) SetMessage(v string) *StartJobResponseBody {
	s.Message = &v
	return s
}

func (s *StartJobResponseBody) SetRequestId(v string) *StartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartJobResponseBody) SetSuccess(v bool) *StartJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartJobResponseBody) SetTaskIds(v []*StartJobResponseBodyTaskIds) *StartJobResponseBody {
	s.TaskIds = v
	return s
}

func (s *StartJobResponseBody) Validate() error {
	if s.CallIds != nil {
		for _, item := range s.CallIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskIds != nil {
		for _, item := range s.TaskIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartJobResponseBodyCallIds struct {
	// example:
	//
	// c93cdd1c-f9b5-4758-be43-7a237a7eaa1d
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 62229628-45d8-41bd-a80f-6e4c0a39f79b
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StartJobResponseBodyCallIds) String() string {
	return dara.Prettify(s)
}

func (s StartJobResponseBodyCallIds) GoString() string {
	return s.String()
}

func (s *StartJobResponseBodyCallIds) GetKey() *string {
	return s.Key
}

func (s *StartJobResponseBodyCallIds) GetValue() *string {
	return s.Value
}

func (s *StartJobResponseBodyCallIds) SetKey(v string) *StartJobResponseBodyCallIds {
	s.Key = &v
	return s
}

func (s *StartJobResponseBodyCallIds) SetValue(v string) *StartJobResponseBodyCallIds {
	s.Value = &v
	return s
}

func (s *StartJobResponseBodyCallIds) Validate() error {
	return dara.Validate(s)
}

type StartJobResponseBodyTaskIds struct {
	// example:
	//
	// c93cdd1c-f9b5-4758-be43-7a237a7eaa1d
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 62229628-45d8-41bd-a80f-6e4c0a39f79b
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StartJobResponseBodyTaskIds) String() string {
	return dara.Prettify(s)
}

func (s StartJobResponseBodyTaskIds) GoString() string {
	return s.String()
}

func (s *StartJobResponseBodyTaskIds) GetKey() *string {
	return s.Key
}

func (s *StartJobResponseBodyTaskIds) GetValue() *string {
	return s.Value
}

func (s *StartJobResponseBodyTaskIds) SetKey(v string) *StartJobResponseBodyTaskIds {
	s.Key = &v
	return s
}

func (s *StartJobResponseBodyTaskIds) SetValue(v string) *StartJobResponseBodyTaskIds {
	s.Value = &v
	return s
}

func (s *StartJobResponseBodyTaskIds) Validate() error {
	return dara.Validate(s)
}
