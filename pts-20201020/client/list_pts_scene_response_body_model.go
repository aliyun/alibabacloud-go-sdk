// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPtsSceneResponseBody
	GetRequestId() *string
	SetSceneViewList(v []*ListPtsSceneResponseBodySceneViewList) *ListPtsSceneResponseBody
	GetSceneViewList() []*ListPtsSceneResponseBodySceneViewList
	SetSuccess(v bool) *ListPtsSceneResponseBody
	GetSuccess() *bool
}

type ListPtsSceneResponseBody struct {
	// The system status code. If the request was successful, no data is returned.
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the request was successful, no data is returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, no data is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6F2ED8-E31B-497F-85AB-C4E358A5F667
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned scenarios.
	SceneViewList []*ListPtsSceneResponseBodySceneViewList `json:"SceneViewList,omitempty" xml:"SceneViewList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPtsSceneResponseBody) GetSceneViewList() []*ListPtsSceneResponseBodySceneViewList {
	return s.SceneViewList
}

func (s *ListPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPtsSceneResponseBody) SetCode(v string) *ListPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetHttpStatusCode(v int32) *ListPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetMessage(v string) *ListPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetRequestId(v string) *ListPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetSceneViewList(v []*ListPtsSceneResponseBodySceneViewList) *ListPtsSceneResponseBody {
	s.SceneViewList = v
	return s
}

func (s *ListPtsSceneResponseBody) SetSuccess(v bool) *ListPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListPtsSceneResponseBody) Validate() error {
	if s.SceneViewList != nil {
		for _, item := range s.SceneViewList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPtsSceneResponseBodySceneViewList struct {
	// The time when the PTS scenario was created.
	//
	// example:
	//
	// 2021-02-26 15:28:39
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// DFGVS3S
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scenario name.
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the PTS scenario. Valid values:
	//
	// example:
	//
	// Draft WaitStart Debugging Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPtsSceneResponseBodySceneViewList) String() string {
	return dara.Prettify(s)
}

func (s ListPtsSceneResponseBodySceneViewList) GoString() string {
	return s.String()
}

func (s *ListPtsSceneResponseBodySceneViewList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPtsSceneResponseBodySceneViewList) GetSceneId() *string {
	return s.SceneId
}

func (s *ListPtsSceneResponseBodySceneViewList) GetSceneName() *string {
	return s.SceneName
}

func (s *ListPtsSceneResponseBodySceneViewList) GetStatus() *string {
	return s.Status
}

func (s *ListPtsSceneResponseBodySceneViewList) SetCreateTime(v string) *ListPtsSceneResponseBodySceneViewList {
	s.CreateTime = &v
	return s
}

func (s *ListPtsSceneResponseBodySceneViewList) SetSceneId(v string) *ListPtsSceneResponseBodySceneViewList {
	s.SceneId = &v
	return s
}

func (s *ListPtsSceneResponseBodySceneViewList) SetSceneName(v string) *ListPtsSceneResponseBodySceneViewList {
	s.SceneName = &v
	return s
}

func (s *ListPtsSceneResponseBodySceneViewList) SetStatus(v string) *ListPtsSceneResponseBodySceneViewList {
	s.Status = &v
	return s
}

func (s *ListPtsSceneResponseBodySceneViewList) Validate() error {
	return dara.Validate(s)
}
