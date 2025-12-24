// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPredictiveTaskInfoResponseBody
	GetRequestId() *string
	SetTaskInfos(v *QueryPredictiveTaskInfoResponseBodyTaskInfos) *QueryPredictiveTaskInfoResponseBody
	GetTaskInfos() *QueryPredictiveTaskInfoResponseBodyTaskInfos
}

type QueryPredictiveTaskInfoResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfos *QueryPredictiveTaskInfoResponseBodyTaskInfos `json:"TaskInfos,omitempty" xml:"TaskInfos,omitempty" type:"Struct"`
}

func (s QueryPredictiveTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPredictiveTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPredictiveTaskInfoResponseBody) GetTaskInfos() *QueryPredictiveTaskInfoResponseBodyTaskInfos {
	return s.TaskInfos
}

func (s *QueryPredictiveTaskInfoResponseBody) SetRequestId(v string) *QueryPredictiveTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPredictiveTaskInfoResponseBody) SetTaskInfos(v *QueryPredictiveTaskInfoResponseBodyTaskInfos) *QueryPredictiveTaskInfoResponseBody {
	s.TaskInfos = v
	return s
}

func (s *QueryPredictiveTaskInfoResponseBody) Validate() error {
	if s.TaskInfos != nil {
		if err := s.TaskInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPredictiveTaskInfoResponseBodyTaskInfos struct {
	TaskInfo []*QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Repeated"`
}

func (s QueryPredictiveTaskInfoResponseBodyTaskInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveTaskInfoResponseBodyTaskInfos) GoString() string {
	return s.String()
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfos) GetTaskInfo() []*QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo {
	return s.TaskInfo
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfos) SetTaskInfo(v []*QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) *QueryPredictiveTaskInfoResponseBodyTaskInfos {
	s.TaskInfo = v
	return s
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfos) Validate() error {
	if s.TaskInfo != nil {
		for _, item := range s.TaskInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo struct {
	// example:
	//
	// 10
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// example:
	//
	// 2
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// example:
	//
	// 2025-12-17T10:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) GoString() string {
	return s.String()
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) GetMinValue() *int32 {
	return s.MinValue
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) GetTime() *string {
	return s.Time
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) SetMaxValue(v int32) *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo {
	s.MaxValue = &v
	return s
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) SetMinValue(v int32) *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo {
	s.MinValue = &v
	return s
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) SetTime(v string) *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo {
	s.Time = &v
	return s
}

func (s *QueryPredictiveTaskInfoResponseBodyTaskInfosTaskInfo) Validate() error {
	return dara.Validate(s)
}
