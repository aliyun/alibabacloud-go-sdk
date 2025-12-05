// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StartPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartPtsSceneResponseBody
	GetMessage() *string
	SetPlanId(v string) *StartPtsSceneResponseBody
	GetPlanId() *string
	SetRequestId(v string) *StartPtsSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartPtsSceneResponseBody
	GetSuccess() *bool
}

type StartPtsSceneResponseBody struct {
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
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The performance testing plan ID. If the scenario is successfully started, this parameter is returned.
	//
	// example:
	//
	// SFVAFE
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BD12DCC9-5E48-4E77-9657-8D34D8C0F97B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
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

func (s StartPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartPtsSceneResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *StartPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartPtsSceneResponseBody) SetCode(v string) *StartPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetHttpStatusCode(v int32) *StartPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetMessage(v string) *StartPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetPlanId(v string) *StartPtsSceneResponseBody {
	s.PlanId = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetRequestId(v string) *StartPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetSuccess(v bool) *StartPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StartPtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
