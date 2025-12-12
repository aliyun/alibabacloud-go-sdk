// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSyncQualityCheckDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSyncQualityCheckDataResponseBody
	GetCode() *string
	SetData(v *UpdateSyncQualityCheckDataResponseBodyData) *UpdateSyncQualityCheckDataResponseBody
	GetData() *UpdateSyncQualityCheckDataResponseBodyData
	SetMessage(v string) *UpdateSyncQualityCheckDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSyncQualityCheckDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSyncQualityCheckDataResponseBody
	GetSuccess() *bool
}

type UpdateSyncQualityCheckDataResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateSyncQualityCheckDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F8***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSyncQualityCheckDataResponseBody) GetData() *UpdateSyncQualityCheckDataResponseBodyData {
	return s.Data
}

func (s *UpdateSyncQualityCheckDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSyncQualityCheckDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSyncQualityCheckDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetCode(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetData(v *UpdateSyncQualityCheckDataResponseBodyData) *UpdateSyncQualityCheckDataResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetMessage(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetRequestId(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetSuccess(v bool) *UpdateSyncQualityCheckDataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSyncQualityCheckDataResponseBodyData struct {
	// example:
	//
	// 123123D8C-5BD9-42A7-B527-1235F8**
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 20210101-1212121***
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) GetTid() *string {
	return s.Tid
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) SetTaskId(v string) *UpdateSyncQualityCheckDataResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) SetTid(v string) *UpdateSyncQualityCheckDataResponseBodyData {
	s.Tid = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
