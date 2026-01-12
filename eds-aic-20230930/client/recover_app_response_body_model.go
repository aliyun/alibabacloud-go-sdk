// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *RecoverAppResponseBody
	GetCount() *int64
	SetData(v []*RecoverAppResponseBodyData) *RecoverAppResponseBody
	GetData() []*RecoverAppResponseBodyData
	SetRequestId(v string) *RecoverAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RecoverAppResponseBody
	GetTaskId() *string
}

type RecoverAppResponseBody struct {
	// example:
	//
	// 1
	Count *int64                        `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*RecoverAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoverAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverAppResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *RecoverAppResponseBody) GetData() []*RecoverAppResponseBodyData {
	return s.Data
}

func (s *RecoverAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RecoverAppResponseBody) SetCount(v int64) *RecoverAppResponseBody {
	s.Count = &v
	return s
}

func (s *RecoverAppResponseBody) SetData(v []*RecoverAppResponseBodyData) *RecoverAppResponseBody {
	s.Data = v
	return s
}

func (s *RecoverAppResponseBody) SetRequestId(v string) *RecoverAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverAppResponseBody) SetTaskId(v string) *RecoverAppResponseBody {
	s.TaskId = &v
	return s
}

func (s *RecoverAppResponseBody) Validate() error {
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

type RecoverAppResponseBodyData struct {
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// example:
	//
	// t-22ex666a653gq****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoverAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecoverAppResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *RecoverAppResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RecoverAppResponseBodyData) SetAndroidInstanceId(v string) *RecoverAppResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *RecoverAppResponseBodyData) SetTaskId(v string) *RecoverAppResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RecoverAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
