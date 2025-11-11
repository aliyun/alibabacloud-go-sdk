// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneGlobalReplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsertInterveneGlobalReplyResponseBody
	GetCode() *string
	SetData(v *InsertInterveneGlobalReplyResponseBodyData) *InsertInterveneGlobalReplyResponseBody
	GetData() *InsertInterveneGlobalReplyResponseBodyData
	SetHttpStatusCode(v int32) *InsertInterveneGlobalReplyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InsertInterveneGlobalReplyResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertInterveneGlobalReplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsertInterveneGlobalReplyResponseBody
	GetSuccess() *bool
}

type InsertInterveneGlobalReplyResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InsertInterveneGlobalReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertInterveneGlobalReplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsertInterveneGlobalReplyResponseBody) GetData() *InsertInterveneGlobalReplyResponseBodyData {
	return s.Data
}

func (s *InsertInterveneGlobalReplyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InsertInterveneGlobalReplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertInterveneGlobalReplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertInterveneGlobalReplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsertInterveneGlobalReplyResponseBody) SetCode(v string) *InsertInterveneGlobalReplyResponseBody {
	s.Code = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetData(v *InsertInterveneGlobalReplyResponseBodyData) *InsertInterveneGlobalReplyResponseBody {
	s.Data = v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetHttpStatusCode(v int32) *InsertInterveneGlobalReplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetMessage(v string) *InsertInterveneGlobalReplyResponseBody {
	s.Message = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetRequestId(v string) *InsertInterveneGlobalReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetSuccess(v bool) *InsertInterveneGlobalReplyResponseBody {
	s.Success = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertInterveneGlobalReplyResponseBodyData struct {
	Code       *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 4829
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InsertInterveneGlobalReplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *InsertInterveneGlobalReplyResponseBodyData) GetFailIdList() []*string {
	return s.FailIdList
}

func (s *InsertInterveneGlobalReplyResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetCode(v int32) *InsertInterveneGlobalReplyResponseBodyData {
	s.Code = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetFailIdList(v []*string) *InsertInterveneGlobalReplyResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetTaskId(v string) *InsertInterveneGlobalReplyResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
