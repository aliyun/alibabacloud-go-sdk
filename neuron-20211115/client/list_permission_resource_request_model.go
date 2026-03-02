// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ListPermissionResourceRequest
	GetAction() *string
	SetEnterpriseId(v int64) *ListPermissionResourceRequest
	GetEnterpriseId() *int64
	SetOperatorId(v string) *ListPermissionResourceRequest
	GetOperatorId() *string
	SetOperatorType(v string) *ListPermissionResourceRequest
	GetOperatorType() *string
	SetResourcePrefix(v string) *ListPermissionResourceRequest
	GetResourcePrefix() *string
}

type ListPermissionResourceRequest struct {
	Action         *string `json:"action,omitempty" xml:"action,omitempty"`
	EnterpriseId   *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	OperatorId     *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OperatorType   *string `json:"operatorType,omitempty" xml:"operatorType,omitempty"`
	ResourcePrefix *string `json:"resourcePrefix,omitempty" xml:"resourcePrefix,omitempty"`
}

func (s ListPermissionResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourceRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionResourceRequest) GetAction() *string {
	return s.Action
}

func (s *ListPermissionResourceRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListPermissionResourceRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListPermissionResourceRequest) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListPermissionResourceRequest) GetResourcePrefix() *string {
	return s.ResourcePrefix
}

func (s *ListPermissionResourceRequest) SetAction(v string) *ListPermissionResourceRequest {
	s.Action = &v
	return s
}

func (s *ListPermissionResourceRequest) SetEnterpriseId(v int64) *ListPermissionResourceRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListPermissionResourceRequest) SetOperatorId(v string) *ListPermissionResourceRequest {
	s.OperatorId = &v
	return s
}

func (s *ListPermissionResourceRequest) SetOperatorType(v string) *ListPermissionResourceRequest {
	s.OperatorType = &v
	return s
}

func (s *ListPermissionResourceRequest) SetResourcePrefix(v string) *ListPermissionResourceRequest {
	s.ResourcePrefix = &v
	return s
}

func (s *ListPermissionResourceRequest) Validate() error {
	return dara.Validate(s)
}
