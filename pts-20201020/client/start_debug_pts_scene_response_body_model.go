// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebugPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartDebugPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StartDebugPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartDebugPtsSceneResponseBody
	GetMessage() *string
	SetPlanId(v string) *StartDebugPtsSceneResponseBody
	GetPlanId() *string
	SetRequestId(v string) *StartDebugPtsSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDebugPtsSceneResponseBody
	GetSuccess() *bool
}

type StartDebugPtsSceneResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the stress testing task.
	//
	// example:
	//
	// NJJBH8B
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C1905194-EE28-4F78-AD81-85A40D52D1BC
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
}

func (s StartDebugPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDebugPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartDebugPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartDebugPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartDebugPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartDebugPtsSceneResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *StartDebugPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDebugPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDebugPtsSceneResponseBody) SetCode(v string) *StartDebugPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetHttpStatusCode(v int32) *StartDebugPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetMessage(v string) *StartDebugPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetPlanId(v string) *StartDebugPtsSceneResponseBody {
	s.PlanId = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetRequestId(v string) *StartDebugPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetSuccess(v bool) *StartDebugPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
