// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportInterveneFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportInterveneFileResponseBody
	GetCode() *string
	SetData(v *ImportInterveneFileResponseBodyData) *ImportInterveneFileResponseBody
	GetData() *ImportInterveneFileResponseBodyData
	SetHttpStatusCode(v int32) *ImportInterveneFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportInterveneFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportInterveneFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportInterveneFileResponseBody
	GetSuccess() *bool
}

type ImportInterveneFileResponseBody struct {
	// example:
	//
	// 0
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ImportInterveneFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
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

func (s ImportInterveneFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileResponseBody) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportInterveneFileResponseBody) GetData() *ImportInterveneFileResponseBodyData {
	return s.Data
}

func (s *ImportInterveneFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportInterveneFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportInterveneFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportInterveneFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportInterveneFileResponseBody) SetCode(v string) *ImportInterveneFileResponseBody {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetData(v *ImportInterveneFileResponseBodyData) *ImportInterveneFileResponseBody {
	s.Data = v
	return s
}

func (s *ImportInterveneFileResponseBody) SetHttpStatusCode(v int32) *ImportInterveneFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetMessage(v string) *ImportInterveneFileResponseBody {
	s.Message = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetRequestId(v string) *ImportInterveneFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetSuccess(v bool) *ImportInterveneFileResponseBody {
	s.Success = &v
	return s
}

func (s *ImportInterveneFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportInterveneFileResponseBodyData struct {
	Code       *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportInterveneFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ImportInterveneFileResponseBodyData) GetFailIdList() []*string {
	return s.FailIdList
}

func (s *ImportInterveneFileResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImportInterveneFileResponseBodyData) SetCode(v int32) *ImportInterveneFileResponseBodyData {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileResponseBodyData) SetFailIdList(v []*string) *ImportInterveneFileResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ImportInterveneFileResponseBodyData) SetTaskId(v string) *ImportInterveneFileResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImportInterveneFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
