// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSkillGroupResponseBody
	GetCode() *string
	SetData(v *CreateSkillGroupResponseBodyData) *CreateSkillGroupResponseBody
	GetData() *CreateSkillGroupResponseBodyData
	SetHttpStatusCode(v int32) *CreateSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSkillGroupResponseBody
	GetRequestId() *string
}

type CreateSkillGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSkillGroupResponseBody) GetData() *CreateSkillGroupResponseBodyData {
	return s.Data
}

func (s *CreateSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillGroupResponseBody) SetCode(v string) *CreateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetData(v *CreateSkillGroupResponseBodyData) *CreateSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateSkillGroupResponseBody) SetHttpStatusCode(v int32) *CreateSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetMessage(v string) *CreateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetRequestId(v string) *CreateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSkillGroupResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// skillgroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s CreateSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSkillGroupResponseBodyData) GetMediaType() *string {
	return s.MediaType
}

func (s *CreateSkillGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateSkillGroupResponseBodyData) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *CreateSkillGroupResponseBodyData) SetDescription(v string) *CreateSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetInstanceId(v string) *CreateSkillGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetMediaType(v string) *CreateSkillGroupResponseBodyData {
	s.MediaType = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetName(v string) *CreateSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetSkillGroupId(v string) *CreateSkillGroupResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
