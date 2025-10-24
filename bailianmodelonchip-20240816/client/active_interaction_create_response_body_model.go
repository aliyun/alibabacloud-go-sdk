// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveInteractionCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActiveInteractionCreateResponseBody
	GetCode() *string
	SetData(v *ActiveInteractionCreateResponseBodyData) *ActiveInteractionCreateResponseBody
	GetData() *ActiveInteractionCreateResponseBodyData
	SetHttpStatusCode(v int32) *ActiveInteractionCreateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ActiveInteractionCreateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActiveInteractionCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ActiveInteractionCreateResponseBody
	GetSuccess() *bool
}

type ActiveInteractionCreateResponseBody struct {
	// example:
	//
	// success
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *ActiveInteractionCreateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AF54F772-60FF-56FD-A3EA-11620EF1229A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ActiveInteractionCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveInteractionCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActiveInteractionCreateResponseBody) GetData() *ActiveInteractionCreateResponseBodyData {
	return s.Data
}

func (s *ActiveInteractionCreateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ActiveInteractionCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActiveInteractionCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveInteractionCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ActiveInteractionCreateResponseBody) SetCode(v string) *ActiveInteractionCreateResponseBody {
	s.Code = &v
	return s
}

func (s *ActiveInteractionCreateResponseBody) SetData(v *ActiveInteractionCreateResponseBodyData) *ActiveInteractionCreateResponseBody {
	s.Data = v
	return s
}

func (s *ActiveInteractionCreateResponseBody) SetHttpStatusCode(v int32) *ActiveInteractionCreateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActiveInteractionCreateResponseBody) SetMessage(v string) *ActiveInteractionCreateResponseBody {
	s.Message = &v
	return s
}

func (s *ActiveInteractionCreateResponseBody) SetRequestId(v string) *ActiveInteractionCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveInteractionCreateResponseBody) SetSuccess(v bool) *ActiveInteractionCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveInteractionCreateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActiveInteractionCreateResponseBodyData struct {
	// example:
	//
	// point
	Gesture *string `json:"gesture,omitempty" xml:"gesture,omitempty"`
	// example:
	//
	// A man in dark clothing stands on a rocky hilltop, facing away from the camera, gazing at the expansive view with a contemplative posture.
	Person *string `json:"person,omitempty" xml:"person,omitempty"`
	// example:
	//
	// Mountainous landscape with layered ridges receding into haze, under a vast blue sky with wispy clouds and soft golden light near the horizon.
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s ActiveInteractionCreateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionCreateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ActiveInteractionCreateResponseBodyData) GetGesture() *string {
	return s.Gesture
}

func (s *ActiveInteractionCreateResponseBodyData) GetPerson() *string {
	return s.Person
}

func (s *ActiveInteractionCreateResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *ActiveInteractionCreateResponseBodyData) SetGesture(v string) *ActiveInteractionCreateResponseBodyData {
	s.Gesture = &v
	return s
}

func (s *ActiveInteractionCreateResponseBodyData) SetPerson(v string) *ActiveInteractionCreateResponseBodyData {
	s.Person = &v
	return s
}

func (s *ActiveInteractionCreateResponseBodyData) SetScene(v string) *ActiveInteractionCreateResponseBodyData {
	s.Scene = &v
	return s
}

func (s *ActiveInteractionCreateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
