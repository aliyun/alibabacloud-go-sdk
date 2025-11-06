// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListProjectResponseBody
	GetCode() *int32
	SetData(v *ListProjectResponseBodyData) *ListProjectResponseBody
	GetData() *ListProjectResponseBodyData
	SetMessage(v string) *ListProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectResponseBody
	GetSuccess() *bool
}

type ListProjectResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListProjectResponseBody) GetData() *ListProjectResponseBodyData {
	return s.Data
}

func (s *ListProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectResponseBody) SetCode(v int32) *ListProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectResponseBody) SetData(v *ListProjectResponseBodyData) *ListProjectResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectResponseBody) SetMessage(v string) *ListProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectResponseBody) SetSuccess(v bool) *ListProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectResponseBodyData struct {
	Projects []*string `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s ListProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyData) GetProjects() []*string {
	return s.Projects
}

func (s *ListProjectResponseBodyData) SetProjects(v []*string) *ListProjectResponseBodyData {
	s.Projects = v
	return s
}

func (s *ListProjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
