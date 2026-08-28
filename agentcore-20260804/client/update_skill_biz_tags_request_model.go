// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillBizTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateSkillBizTagsRequestBody) *UpdateSkillBizTagsRequest
	GetBody() *UpdateSkillBizTagsRequestBody
}

type UpdateSkillBizTagsRequest struct {
	// The request body.
	Body *UpdateSkillBizTagsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateSkillBizTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillBizTagsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillBizTagsRequest) GetBody() *UpdateSkillBizTagsRequestBody {
	return s.Body
}

func (s *UpdateSkillBizTagsRequest) SetBody(v *UpdateSkillBizTagsRequestBody) *UpdateSkillBizTagsRequest {
	s.Body = v
	return s
}

func (s *UpdateSkillBizTagsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSkillBizTagsRequestBody struct {
	// The business tags as a JSON array string.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cs","qa","support"]
	BizTags *string `json:"bizTags,omitempty" xml:"bizTags,omitempty"`
}

func (s UpdateSkillBizTagsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillBizTagsRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillBizTagsRequestBody) GetBizTags() *string {
	return s.BizTags
}

func (s *UpdateSkillBizTagsRequestBody) SetBizTags(v string) *UpdateSkillBizTagsRequestBody {
	s.BizTags = &v
	return s
}

func (s *UpdateSkillBizTagsRequestBody) Validate() error {
	return dara.Validate(s)
}
