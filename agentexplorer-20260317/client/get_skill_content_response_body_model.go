// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetSkillContentResponseBody
	GetContent() *string
	SetRequestId(v string) *GetSkillContentResponseBody
	GetRequestId() *string
}

type GetSkillContentResponseBody struct {
	// example:
	//
	// ---
	//
	// name: alibabacloud-find-skills
	//
	// description: "Search for official Alibaba Cloud Agent Skills based on user requirements"
	//
	// ---
	//
	// Agent Skill Body Content Here
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1764D64D-5262-55DA-BDBF-1F949B1B34F7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSkillContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetSkillContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillContentResponseBody) SetContent(v string) *GetSkillContentResponseBody {
	s.Content = &v
	return s
}

func (s *GetSkillContentResponseBody) SetRequestId(v string) *GetSkillContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillContentResponseBody) Validate() error {
	return dara.Validate(s)
}
