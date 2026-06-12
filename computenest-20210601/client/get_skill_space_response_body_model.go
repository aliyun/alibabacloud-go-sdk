// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetSkillSpaceResponseBody
	GetCreateTime() *string
	SetRequestId(v string) *GetSkillSpaceResponseBody
	GetRequestId() *string
	SetSkillSpaceDescription(v string) *GetSkillSpaceResponseBody
	GetSkillSpaceDescription() *string
	SetSkillSpaceId(v string) *GetSkillSpaceResponseBody
	GetSkillSpaceId() *string
	SetSkillSpaceName(v string) *GetSkillSpaceResponseBody
	GetSkillSpaceName() *string
	SetUpdateTime(v string) *GetSkillSpaceResponseBody
	GetUpdateTime() *string
}

type GetSkillSpaceResponseBody struct {
	// The time when the SkillSpace was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of the SkillSpace.
	//
	// example:
	//
	// 1111111
	SkillSpaceDescription *string `json:"SkillSpaceDescription,omitempty" xml:"SkillSpaceDescription,omitempty"`
	// The ID of the SkillSpace.
	//
	// example:
	//
	// ss-11111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The name of the SkillSpace.
	//
	// example:
	//
	// 1111111
	SkillSpaceName *string `json:"SkillSpaceName,omitempty" xml:"SkillSpaceName,omitempty"`
	// The time when the SkillSpace was last updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSkillSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillSpaceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSkillSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillSpaceResponseBody) GetSkillSpaceDescription() *string {
	return s.SkillSpaceDescription
}

func (s *GetSkillSpaceResponseBody) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *GetSkillSpaceResponseBody) GetSkillSpaceName() *string {
	return s.SkillSpaceName
}

func (s *GetSkillSpaceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSkillSpaceResponseBody) SetCreateTime(v string) *GetSkillSpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSkillSpaceResponseBody) SetRequestId(v string) *GetSkillSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillSpaceResponseBody) SetSkillSpaceDescription(v string) *GetSkillSpaceResponseBody {
	s.SkillSpaceDescription = &v
	return s
}

func (s *GetSkillSpaceResponseBody) SetSkillSpaceId(v string) *GetSkillSpaceResponseBody {
	s.SkillSpaceId = &v
	return s
}

func (s *GetSkillSpaceResponseBody) SetSkillSpaceName(v string) *GetSkillSpaceResponseBody {
	s.SkillSpaceName = &v
	return s
}

func (s *GetSkillSpaceResponseBody) SetUpdateTime(v string) *GetSkillSpaceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
