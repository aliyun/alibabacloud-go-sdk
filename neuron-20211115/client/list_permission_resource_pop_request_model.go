// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourcePopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ListPermissionResourcePopRequest
	GetAction() *string
	SetOperatorId(v string) *ListPermissionResourcePopRequest
	GetOperatorId() *string
	SetOperatorType(v string) *ListPermissionResourcePopRequest
	GetOperatorType() *string
	SetResourcePrefix(v string) *ListPermissionResourcePopRequest
	GetResourcePrefix() *string
}

type ListPermissionResourcePopRequest struct {
	Action         *string `json:"action,omitempty" xml:"action,omitempty"`
	OperatorId     *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OperatorType   *string `json:"operatorType,omitempty" xml:"operatorType,omitempty"`
	ResourcePrefix *string `json:"resourcePrefix,omitempty" xml:"resourcePrefix,omitempty"`
}

func (s ListPermissionResourcePopRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourcePopRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionResourcePopRequest) GetAction() *string {
	return s.Action
}

func (s *ListPermissionResourcePopRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListPermissionResourcePopRequest) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListPermissionResourcePopRequest) GetResourcePrefix() *string {
	return s.ResourcePrefix
}

func (s *ListPermissionResourcePopRequest) SetAction(v string) *ListPermissionResourcePopRequest {
	s.Action = &v
	return s
}

func (s *ListPermissionResourcePopRequest) SetOperatorId(v string) *ListPermissionResourcePopRequest {
	s.OperatorId = &v
	return s
}

func (s *ListPermissionResourcePopRequest) SetOperatorType(v string) *ListPermissionResourcePopRequest {
	s.OperatorType = &v
	return s
}

func (s *ListPermissionResourcePopRequest) SetResourcePrefix(v string) *ListPermissionResourcePopRequest {
	s.ResourcePrefix = &v
	return s
}

func (s *ListPermissionResourcePopRequest) Validate() error {
	return dara.Validate(s)
}
