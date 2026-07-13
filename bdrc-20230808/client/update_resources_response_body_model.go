// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateResourcesResponseBodyData) *UpdateResourcesResponseBody
	GetData() *UpdateResourcesResponseBodyData
	SetRequestId(v string) *UpdateResourcesResponseBody
	GetRequestId() *string
}

type UpdateResourcesResponseBody struct {
	// The returned data.
	Data *UpdateResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identity of the request.
	//
	// example:
	//
	// 5B2F09BF-CEBD-5A7E-AC01-E7F86169A5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourcesResponseBody) GetData() *UpdateResourcesResponseBodyData {
	return s.Data
}

func (s *UpdateResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourcesResponseBody) SetData(v *UpdateResourcesResponseBodyData) *UpdateResourcesResponseBody {
	s.Data = v
	return s
}

func (s *UpdateResourcesResponseBody) SetRequestId(v string) *UpdateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateResourcesResponseBodyData struct {
	// The unique identity of the asynchronous task.
	//
	// example:
	//
	// t-bp1ewftyzmeg3bl4dtd2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateResourcesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateResourcesResponseBodyData) SetTaskId(v string) *UpdateResourcesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
