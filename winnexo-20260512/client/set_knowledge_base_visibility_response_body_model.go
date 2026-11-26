// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKnowledgeBaseVisibilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetKnowledgeBaseVisibilityResponseBody
	GetCode() *string
	SetMessage(v string) *SetKnowledgeBaseVisibilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetKnowledgeBaseVisibilityResponseBody
	GetRequestId() *string
	SetTargetId(v string) *SetKnowledgeBaseVisibilityResponseBody
	GetTargetId() *string
	SetTargetType(v string) *SetKnowledgeBaseVisibilityResponseBody
	GetTargetType() *string
	SetUserGroupIds(v []*string) *SetKnowledgeBaseVisibilityResponseBody
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *SetKnowledgeBaseVisibilityResponseBody
	GetUserIds() []*string
	SetVisibilityMode(v string) *SetKnowledgeBaseVisibilityResponseBody
	GetVisibilityMode() *string
}

type SetKnowledgeBaseVisibilityResponseBody struct {
	// The status code.
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
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The target ID, which varies based on the target type.
	//
	// - If the type is period, specify the period ID.
	//
	// - If the type is objective, specify the objective ID.
	//
	// - If the type is kr, specify the keyResult ID.
	//
	// example:
	//
	// kb_001
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// The type of the current permission. Valid values:
	//
	// - period: period type
	//
	// - objective: objective type
	//
	// - kr: keyResult type
	//
	// example:
	//
	// KB
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The collection of user groups.
	//
	// example:
	//
	// ["group-1"]
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The collection of users.
	//
	// example:
	//
	// ["10001"]
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// The visibility scope mode.
	//
	// example:
	//
	// SPECIFIED
	VisibilityMode *string `json:"visibilityMode,omitempty" xml:"visibilityMode,omitempty"`
}

func (s SetKnowledgeBaseVisibilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetKnowledgeBaseVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetTargetId() *string {
	return s.TargetId
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetTargetType() *string {
	return s.TargetType
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetUserIds() []*string {
	return s.UserIds
}

func (s *SetKnowledgeBaseVisibilityResponseBody) GetVisibilityMode() *string {
	return s.VisibilityMode
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetCode(v string) *SetKnowledgeBaseVisibilityResponseBody {
	s.Code = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetMessage(v string) *SetKnowledgeBaseVisibilityResponseBody {
	s.Message = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetRequestId(v string) *SetKnowledgeBaseVisibilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetTargetId(v string) *SetKnowledgeBaseVisibilityResponseBody {
	s.TargetId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetTargetType(v string) *SetKnowledgeBaseVisibilityResponseBody {
	s.TargetType = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetUserGroupIds(v []*string) *SetKnowledgeBaseVisibilityResponseBody {
	s.UserGroupIds = v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetUserIds(v []*string) *SetKnowledgeBaseVisibilityResponseBody {
	s.UserIds = v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) SetVisibilityMode(v string) *SetKnowledgeBaseVisibilityResponseBody {
	s.VisibilityMode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponseBody) Validate() error {
	return dara.Validate(s)
}
