// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourceForFrontRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ListPermissionResourceForFrontRequest
	GetAction() *string
	SetEnterpriseId(v string) *ListPermissionResourceForFrontRequest
	GetEnterpriseId() *string
	SetOperatorId(v string) *ListPermissionResourceForFrontRequest
	GetOperatorId() *string
	SetOperatorType(v string) *ListPermissionResourceForFrontRequest
	GetOperatorType() *string
	SetQueryType(v string) *ListPermissionResourceForFrontRequest
	GetQueryType() *string
	SetResourcePrefix(v string) *ListPermissionResourceForFrontRequest
	GetResourcePrefix() *string
}

type ListPermissionResourceForFrontRequest struct {
	// example:
	//
	// VIEW
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// 41
	EnterpriseId *string `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 292758960042264423
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// DEVELOPER
	OperatorType *string `json:"operatorType,omitempty" xml:"operatorType,omitempty"`
	// example:
	//
	// CONTAIN_MATCH
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// example:
	//
	// neuron:catalog:company/1
	ResourcePrefix *string `json:"resourcePrefix,omitempty" xml:"resourcePrefix,omitempty"`
}

func (s ListPermissionResourceForFrontRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourceForFrontRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionResourceForFrontRequest) GetAction() *string {
	return s.Action
}

func (s *ListPermissionResourceForFrontRequest) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *ListPermissionResourceForFrontRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListPermissionResourceForFrontRequest) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListPermissionResourceForFrontRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListPermissionResourceForFrontRequest) GetResourcePrefix() *string {
	return s.ResourcePrefix
}

func (s *ListPermissionResourceForFrontRequest) SetAction(v string) *ListPermissionResourceForFrontRequest {
	s.Action = &v
	return s
}

func (s *ListPermissionResourceForFrontRequest) SetEnterpriseId(v string) *ListPermissionResourceForFrontRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListPermissionResourceForFrontRequest) SetOperatorId(v string) *ListPermissionResourceForFrontRequest {
	s.OperatorId = &v
	return s
}

func (s *ListPermissionResourceForFrontRequest) SetOperatorType(v string) *ListPermissionResourceForFrontRequest {
	s.OperatorType = &v
	return s
}

func (s *ListPermissionResourceForFrontRequest) SetQueryType(v string) *ListPermissionResourceForFrontRequest {
	s.QueryType = &v
	return s
}

func (s *ListPermissionResourceForFrontRequest) SetResourcePrefix(v string) *ListPermissionResourceForFrontRequest {
	s.ResourcePrefix = &v
	return s
}

func (s *ListPermissionResourceForFrontRequest) Validate() error {
	return dara.Validate(s)
}
