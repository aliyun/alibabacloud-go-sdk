// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupResponseBodyData) *GetSkillGroupResponseBody
	GetData() *GetSkillGroupResponseBodyData
	SetHttpStatusCode(v int32) *GetSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupResponseBody
	GetRequestId() *string
}

type GetSkillGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CF1C21B9-2D49-4B54-880F-FBE248C16903
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupResponseBody) GetData() *GetSkillGroupResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupResponseBody) SetCode(v string) *GetSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupResponseBody) SetData(v *GetSkillGroupResponseBodyData) *GetSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupResponseBody) SetHttpStatusCode(v int32) *GetSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSkillGroupResponseBody) SetMessage(v string) *GetSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupResponseBody) SetRequestId(v string) *GetSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSkillGroupResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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

func (s GetSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetSkillGroupResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSkillGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupResponseBodyData) GetMediaType() *string {
	return s.MediaType
}

func (s *GetSkillGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSkillGroupResponseBodyData) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *GetSkillGroupResponseBodyData) SetDescription(v string) *GetSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSkillGroupResponseBodyData) SetDisplayName(v string) *GetSkillGroupResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetSkillGroupResponseBodyData) SetInstanceId(v string) *GetSkillGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupResponseBodyData) SetMediaType(v string) *GetSkillGroupResponseBodyData {
	s.MediaType = &v
	return s
}

func (s *GetSkillGroupResponseBodyData) SetName(v string) *GetSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillGroupResponseBodyData) SetSkillGroupId(v string) *GetSkillGroupResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *GetSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
