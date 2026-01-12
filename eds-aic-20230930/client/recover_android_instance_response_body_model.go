// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *RecoverAndroidInstanceResponseBody
	GetCount() *int64
	SetData(v []*RecoverAndroidInstanceResponseBodyData) *RecoverAndroidInstanceResponseBody
	GetData() []*RecoverAndroidInstanceResponseBodyData
	SetRequestId(v string) *RecoverAndroidInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RecoverAndroidInstanceResponseBody
	GetTaskId() *string
}

type RecoverAndroidInstanceResponseBody struct {
	// example:
	//
	// 1
	Count *int64                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*RecoverAndroidInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoverAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverAndroidInstanceResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *RecoverAndroidInstanceResponseBody) GetData() []*RecoverAndroidInstanceResponseBodyData {
	return s.Data
}

func (s *RecoverAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverAndroidInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RecoverAndroidInstanceResponseBody) SetCount(v int64) *RecoverAndroidInstanceResponseBody {
	s.Count = &v
	return s
}

func (s *RecoverAndroidInstanceResponseBody) SetData(v []*RecoverAndroidInstanceResponseBodyData) *RecoverAndroidInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RecoverAndroidInstanceResponseBody) SetRequestId(v string) *RecoverAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverAndroidInstanceResponseBody) SetTaskId(v string) *RecoverAndroidInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RecoverAndroidInstanceResponseBody) Validate() error {
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

type RecoverAndroidInstanceResponseBodyData struct {
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoverAndroidInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecoverAndroidInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecoverAndroidInstanceResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *RecoverAndroidInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RecoverAndroidInstanceResponseBodyData) SetAndroidInstanceId(v string) *RecoverAndroidInstanceResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *RecoverAndroidInstanceResponseBodyData) SetTaskId(v string) *RecoverAndroidInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RecoverAndroidInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
