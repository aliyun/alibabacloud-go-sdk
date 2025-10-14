// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGdnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateGdnInstanceResponseBodyData) *CreateGdnInstanceResponseBody
	GetData() *CreateGdnInstanceResponseBodyData
	SetMessage(v string) *CreateGdnInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGdnInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGdnInstanceResponseBody
	GetSuccess() *bool
}

type CreateGdnInstanceResponseBody struct {
	Data *CreateGdnInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0DBA1657-7607-56D6-BB6D-641BF17CCFDB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGdnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGdnInstanceResponseBody) GetData() *CreateGdnInstanceResponseBodyData {
	return s.Data
}

func (s *CreateGdnInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGdnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGdnInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGdnInstanceResponseBody) SetData(v *CreateGdnInstanceResponseBodyData) *CreateGdnInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateGdnInstanceResponseBody) SetMessage(v string) *CreateGdnInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGdnInstanceResponseBody) SetRequestId(v string) *CreateGdnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGdnInstanceResponseBody) SetSuccess(v bool) *CreateGdnInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGdnInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGdnInstanceResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGdnInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGdnInstanceResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateGdnInstanceResponseBodyData) SetTaskId(v int32) *CreateGdnInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateGdnInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
