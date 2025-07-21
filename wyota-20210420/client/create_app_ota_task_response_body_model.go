// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppOtaTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAppOtaTaskResponseBody
	GetCode() *string
	SetData(v *CreateAppOtaTaskResponseBodyData) *CreateAppOtaTaskResponseBody
	GetData() *CreateAppOtaTaskResponseBodyData
	SetMessage(v string) *CreateAppOtaTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAppOtaTaskResponseBody
	GetRequestId() *string
}

type CreateAppOtaTaskResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAppOtaTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppOtaTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAppOtaTaskResponseBody) GetData() *CreateAppOtaTaskResponseBodyData {
	return s.Data
}

func (s *CreateAppOtaTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAppOtaTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppOtaTaskResponseBody) SetCode(v string) *CreateAppOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppOtaTaskResponseBody) SetData(v *CreateAppOtaTaskResponseBodyData) *CreateAppOtaTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppOtaTaskResponseBody) SetMessage(v string) *CreateAppOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppOtaTaskResponseBody) SetRequestId(v string) *CreateAppOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppOtaTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppOtaTaskResponseBodyData struct {
	TaskUid *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s CreateAppOtaTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskResponseBodyData) GetTaskUid() *string {
	return s.TaskUid
}

func (s *CreateAppOtaTaskResponseBodyData) SetTaskUid(v string) *CreateAppOtaTaskResponseBodyData {
	s.TaskUid = &v
	return s
}

func (s *CreateAppOtaTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
