// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveInteractionEuCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActiveInteractionEuCreateResponseBody
	GetCode() *string
	SetData(v *ActiveInteractionEuCreateResponseBodyData) *ActiveInteractionEuCreateResponseBody
	GetData() *ActiveInteractionEuCreateResponseBodyData
	SetHttpStatusCode(v int32) *ActiveInteractionEuCreateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ActiveInteractionEuCreateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActiveInteractionEuCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ActiveInteractionEuCreateResponseBody
	GetSuccess() *bool
}

type ActiveInteractionEuCreateResponseBody struct {
	// example:
	//
	// success
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *ActiveInteractionEuCreateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 84656A01-32F0-5D12-8C72-E3AFAA0C8A29
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ActiveInteractionEuCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionEuCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveInteractionEuCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActiveInteractionEuCreateResponseBody) GetData() *ActiveInteractionEuCreateResponseBodyData {
	return s.Data
}

func (s *ActiveInteractionEuCreateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ActiveInteractionEuCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActiveInteractionEuCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveInteractionEuCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ActiveInteractionEuCreateResponseBody) SetCode(v string) *ActiveInteractionEuCreateResponseBody {
	s.Code = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBody) SetData(v *ActiveInteractionEuCreateResponseBodyData) *ActiveInteractionEuCreateResponseBody {
	s.Data = v
	return s
}

func (s *ActiveInteractionEuCreateResponseBody) SetHttpStatusCode(v int32) *ActiveInteractionEuCreateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBody) SetMessage(v string) *ActiveInteractionEuCreateResponseBody {
	s.Message = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBody) SetRequestId(v string) *ActiveInteractionEuCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBody) SetSuccess(v bool) *ActiveInteractionEuCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActiveInteractionEuCreateResponseBodyData struct {
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

func (s ActiveInteractionEuCreateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionEuCreateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ActiveInteractionEuCreateResponseBodyData) GetGesture() *string {
	return s.Gesture
}

func (s *ActiveInteractionEuCreateResponseBodyData) GetPerson() *string {
	return s.Person
}

func (s *ActiveInteractionEuCreateResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *ActiveInteractionEuCreateResponseBodyData) SetGesture(v string) *ActiveInteractionEuCreateResponseBodyData {
	s.Gesture = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBodyData) SetPerson(v string) *ActiveInteractionEuCreateResponseBodyData {
	s.Person = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBodyData) SetScene(v string) *ActiveInteractionEuCreateResponseBodyData {
	s.Scene = &v
	return s
}

func (s *ActiveInteractionEuCreateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
