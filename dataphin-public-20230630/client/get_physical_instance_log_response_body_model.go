// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalInstanceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhysicalInstanceLogResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPhysicalInstanceLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPhysicalInstanceLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhysicalInstanceLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalInstanceLogResponseBody
	GetSuccess() *bool
	SetTaskrunLogList(v []*GetPhysicalInstanceLogResponseBodyTaskrunLogList) *GetPhysicalInstanceLogResponseBody
	GetTaskrunLogList() []*GetPhysicalInstanceLogResponseBodyTaskrunLogList
}

type GetPhysicalInstanceLogResponseBody struct {
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskrunLogList []*GetPhysicalInstanceLogResponseBodyTaskrunLogList `json:"TaskrunLogList,omitempty" xml:"TaskrunLogList,omitempty" type:"Repeated"`
}

func (s GetPhysicalInstanceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhysicalInstanceLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPhysicalInstanceLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhysicalInstanceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalInstanceLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalInstanceLogResponseBody) GetTaskrunLogList() []*GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	return s.TaskrunLogList
}

func (s *GetPhysicalInstanceLogResponseBody) SetCode(v string) *GetPhysicalInstanceLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetHttpStatusCode(v int32) *GetPhysicalInstanceLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetMessage(v string) *GetPhysicalInstanceLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetRequestId(v string) *GetPhysicalInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetSuccess(v bool) *GetPhysicalInstanceLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetTaskrunLogList(v []*GetPhysicalInstanceLogResponseBodyTaskrunLogList) *GetPhysicalInstanceLogResponseBody {
	s.TaskrunLogList = v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) Validate() error {
	if s.TaskrunLogList != nil {
		for _, item := range s.TaskrunLogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPhysicalInstanceLogResponseBodyTaskrunLogList struct {
	// example:
	//
	// 60s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-05-30 16:48:13
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xx
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tr_23231
	TaskrunId *string `json:"TaskrunId,omitempty" xml:"TaskrunId,omitempty"`
}

func (s GetPhysicalInstanceLogResponseBodyTaskrunLogList) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceLogResponseBodyTaskrunLogList) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) GetDuration() *string {
	return s.Duration
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) GetLogContent() *string {
	return s.LogContent
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) GetStatus() *string {
	return s.Status
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) GetTaskrunId() *string {
	return s.TaskrunId
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetDuration(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.Duration = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetEndTime(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.EndTime = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetLogContent(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.LogContent = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetStartTime(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.StartTime = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetStatus(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.Status = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetTaskrunId(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.TaskrunId = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) Validate() error {
	return dara.Validate(s)
}
