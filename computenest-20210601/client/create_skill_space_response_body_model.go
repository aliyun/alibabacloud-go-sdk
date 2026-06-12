// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSkillSpaceResponseBody
	GetRequestId() *string
	SetSkillSpaceId(v string) *CreateSkillSpaceResponseBody
	GetSkillSpaceId() *string
}

type CreateSkillSpaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the skill space.
	//
	// example:
	//
	// ss-11111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
}

func (s CreateSkillSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillSpaceResponseBody) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *CreateSkillSpaceResponseBody) SetRequestId(v string) *CreateSkillSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillSpaceResponseBody) SetSkillSpaceId(v string) *CreateSkillSpaceResponseBody {
	s.SkillSpaceId = &v
	return s
}

func (s *CreateSkillSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
