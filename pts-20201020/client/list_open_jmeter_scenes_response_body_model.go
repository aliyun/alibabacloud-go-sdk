// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenJMeterScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOpenJMeterScenesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListOpenJMeterScenesResponseBody
	GetHttpStatusCode() *int32
	SetJMeterScene(v []*ListOpenJMeterScenesResponseBodyJMeterScene) *ListOpenJMeterScenesResponseBody
	GetJMeterScene() []*ListOpenJMeterScenesResponseBodyJMeterScene
	SetMessage(v string) *ListOpenJMeterScenesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListOpenJMeterScenesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOpenJMeterScenesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListOpenJMeterScenesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOpenJMeterScenesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListOpenJMeterScenesResponseBody
	GetTotalCount() *int64
}

type ListOpenJMeterScenesResponseBody struct {
	// The system status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The scenarios.
	JMeterScene []*ListOpenJMeterScenesResponseBodyJMeterScene `json:"JMeterScene,omitempty" xml:"JMeterScene,omitempty" type:"Repeated"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of returned scenarios.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned scenarios.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpenJMeterScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpenJMeterScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOpenJMeterScenesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOpenJMeterScenesResponseBody) GetJMeterScene() []*ListOpenJMeterScenesResponseBodyJMeterScene {
	return s.JMeterScene
}

func (s *ListOpenJMeterScenesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOpenJMeterScenesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOpenJMeterScenesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOpenJMeterScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpenJMeterScenesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOpenJMeterScenesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOpenJMeterScenesResponseBody) SetCode(v string) *ListOpenJMeterScenesResponseBody {
	s.Code = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetHttpStatusCode(v int32) *ListOpenJMeterScenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetJMeterScene(v []*ListOpenJMeterScenesResponseBodyJMeterScene) *ListOpenJMeterScenesResponseBody {
	s.JMeterScene = v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetMessage(v string) *ListOpenJMeterScenesResponseBody {
	s.Message = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetPageNumber(v int32) *ListOpenJMeterScenesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetPageSize(v int32) *ListOpenJMeterScenesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetRequestId(v string) *ListOpenJMeterScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetSuccess(v bool) *ListOpenJMeterScenesResponseBody {
	s.Success = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetTotalCount(v int64) *ListOpenJMeterScenesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) Validate() error {
	if s.JMeterScene != nil {
		for _, item := range s.JMeterScene {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOpenJMeterScenesResponseBodyJMeterScene struct {
	// The stress testing duration.
	//
	// example:
	//
	// 10分钟
	DurationStr *string `json:"DurationStr,omitempty" xml:"DurationStr,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the scenario. Valid values:
	//
	// 	- PREPARING: The scenario is being prepared.
	//
	// 	- PREPARED: The scenario has been prepared.
	//
	// 	- STARTING: The scenario is starting.
	//
	// 	- RUNNING: The scenario is running.
	//
	// 	- STOPPING: The scenario is being stopped.
	//
	// 	- STOPPED: The scenario waits for startup
	//
	// example:
	//
	// STOPPED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOpenJMeterScenesResponseBodyJMeterScene) String() string {
	return dara.Prettify(s)
}

func (s ListOpenJMeterScenesResponseBodyJMeterScene) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) GetDurationStr() *string {
	return s.DurationStr
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) GetSceneId() *string {
	return s.SceneId
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) GetSceneName() *string {
	return s.SceneName
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) GetStatus() *string {
	return s.Status
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetDurationStr(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.DurationStr = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetSceneId(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.SceneId = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetSceneName(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.SceneName = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetStatus(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.Status = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) Validate() error {
	return dara.Validate(s)
}
