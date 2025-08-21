// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRenderingSessionsResponseBody
	GetRequestId() *string
	SetSessions(v []*ListRenderingSessionsResponseBodySessions) *ListRenderingSessionsResponseBody
	GetSessions() []*ListRenderingSessionsResponseBodySessions
	SetTotalCount(v int64) *ListRenderingSessionsResponseBody
	GetTotalCount() *int64
}

type ListRenderingSessionsResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sessions  []*ListRenderingSessionsResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingSessionsResponseBody) GetSessions() []*ListRenderingSessionsResponseBodySessions {
	return s.Sessions
}

func (s *ListRenderingSessionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingSessionsResponseBody) SetRequestId(v string) *ListRenderingSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingSessionsResponseBody) SetSessions(v []*ListRenderingSessionsResponseBodySessions) *ListRenderingSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *ListRenderingSessionsResponseBody) SetTotalCount(v int64) *ListRenderingSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRenderingSessionsResponseBodySessions struct {
	// example:
	//
	// cap-4e1a6a425495458ba78693b8ac6600ea
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// fd6b2134-7954-4754-8915-5fb8b0469622
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId             *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-07-04T01:23:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRenderingSessionsResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingSessionsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *ListRenderingSessionsResponseBodySessions) GetAppId() *string {
	return s.AppId
}

func (s *ListRenderingSessionsResponseBodySessions) GetClientId() *string {
	return s.ClientId
}

func (s *ListRenderingSessionsResponseBodySessions) GetPatchId() *string {
	return s.PatchId
}

func (s *ListRenderingSessionsResponseBodySessions) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingSessionsResponseBodySessions) GetSessionId() *string {
	return s.SessionId
}

func (s *ListRenderingSessionsResponseBodySessions) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingSessionsResponseBodySessions) SetAppId(v string) *ListRenderingSessionsResponseBodySessions {
	s.AppId = &v
	return s
}

func (s *ListRenderingSessionsResponseBodySessions) SetClientId(v string) *ListRenderingSessionsResponseBodySessions {
	s.ClientId = &v
	return s
}

func (s *ListRenderingSessionsResponseBodySessions) SetPatchId(v string) *ListRenderingSessionsResponseBodySessions {
	s.PatchId = &v
	return s
}

func (s *ListRenderingSessionsResponseBodySessions) SetRenderingInstanceId(v string) *ListRenderingSessionsResponseBodySessions {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingSessionsResponseBodySessions) SetSessionId(v string) *ListRenderingSessionsResponseBodySessions {
	s.SessionId = &v
	return s
}

func (s *ListRenderingSessionsResponseBodySessions) SetStartTime(v string) *ListRenderingSessionsResponseBodySessions {
	s.StartTime = &v
	return s
}

func (s *ListRenderingSessionsResponseBodySessions) Validate() error {
	return dara.Validate(s)
}
