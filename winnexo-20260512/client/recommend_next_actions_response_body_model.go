// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendNextActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RecommendNextActionsResponseBody
	GetCode() *string
	SetMessage(v string) *RecommendNextActionsResponseBody
	GetMessage() *string
	SetNextActions(v []*RecommendNextActionsResponseBodyNextActions) *RecommendNextActionsResponseBody
	GetNextActions() []*RecommendNextActionsResponseBodyNextActions
	SetRequestId(v string) *RecommendNextActionsResponseBody
	GetRequestId() *string
	SetTitle(v string) *RecommendNextActionsResponseBody
	GetTitle() *string
}

type RecommendNextActionsResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The next-step recommendations.
	NextActions []*RecommendNextActionsResponseBodyNextActions `json:"nextActions,omitempty" xml:"nextActions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The meeting reservation title.
	//
	// example:
	//
	// Sample session title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecommendNextActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecommendNextActionsResponseBody) GoString() string {
	return s.String()
}

func (s *RecommendNextActionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *RecommendNextActionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecommendNextActionsResponseBody) GetNextActions() []*RecommendNextActionsResponseBodyNextActions {
	return s.NextActions
}

func (s *RecommendNextActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecommendNextActionsResponseBody) GetTitle() *string {
	return s.Title
}

func (s *RecommendNextActionsResponseBody) SetCode(v string) *RecommendNextActionsResponseBody {
	s.Code = &v
	return s
}

func (s *RecommendNextActionsResponseBody) SetMessage(v string) *RecommendNextActionsResponseBody {
	s.Message = &v
	return s
}

func (s *RecommendNextActionsResponseBody) SetNextActions(v []*RecommendNextActionsResponseBodyNextActions) *RecommendNextActionsResponseBody {
	s.NextActions = v
	return s
}

func (s *RecommendNextActionsResponseBody) SetRequestId(v string) *RecommendNextActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecommendNextActionsResponseBody) SetTitle(v string) *RecommendNextActionsResponseBody {
	s.Title = &v
	return s
}

func (s *RecommendNextActionsResponseBody) Validate() error {
	if s.NextActions != nil {
		for _, item := range s.NextActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecommendNextActionsResponseBodyNextActions struct {
	// The recommendation title.
	//
	// example:
	//
	// Continue analyzing this metric
	ActionTitle *string `json:"actionTitle,omitempty" xml:"actionTitle,omitempty"`
	// The skill code.
	//
	// example:
	//
	// exampleSkillCode
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The skill name.
	//
	// example:
	//
	// Sample skill
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The recommendation type.
	//
	// example:
	//
	// recommend_reply
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RecommendNextActionsResponseBodyNextActions) String() string {
	return dara.Prettify(s)
}

func (s RecommendNextActionsResponseBodyNextActions) GoString() string {
	return s.String()
}

func (s *RecommendNextActionsResponseBodyNextActions) GetActionTitle() *string {
	return s.ActionTitle
}

func (s *RecommendNextActionsResponseBodyNextActions) GetSkillCode() *string {
	return s.SkillCode
}

func (s *RecommendNextActionsResponseBodyNextActions) GetSkillName() *string {
	return s.SkillName
}

func (s *RecommendNextActionsResponseBodyNextActions) GetType() *string {
	return s.Type
}

func (s *RecommendNextActionsResponseBodyNextActions) SetActionTitle(v string) *RecommendNextActionsResponseBodyNextActions {
	s.ActionTitle = &v
	return s
}

func (s *RecommendNextActionsResponseBodyNextActions) SetSkillCode(v string) *RecommendNextActionsResponseBodyNextActions {
	s.SkillCode = &v
	return s
}

func (s *RecommendNextActionsResponseBodyNextActions) SetSkillName(v string) *RecommendNextActionsResponseBodyNextActions {
	s.SkillName = &v
	return s
}

func (s *RecommendNextActionsResponseBodyNextActions) SetType(v string) *RecommendNextActionsResponseBodyNextActions {
	s.Type = &v
	return s
}

func (s *RecommendNextActionsResponseBodyNextActions) Validate() error {
	return dara.Validate(s)
}
