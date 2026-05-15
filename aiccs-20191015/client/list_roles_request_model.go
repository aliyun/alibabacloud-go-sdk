// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListRolesRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListRolesRequest
	GetInstanceId() *string
}

type ListRolesRequest struct {
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListRolesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRolesRequest) SetClientToken(v string) *ListRolesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRolesRequest) SetInstanceId(v string) *ListRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
