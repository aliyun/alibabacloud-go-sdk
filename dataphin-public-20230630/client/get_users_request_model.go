// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetUsersRequest
	GetOpTenantId() *int64
	SetUserIdList(v []*string) *GetUsersRequest
	GetUserIdList() []*string
}

type GetUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s GetUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUsersRequest) GoString() string {
	return s.String()
}

func (s *GetUsersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUsersRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *GetUsersRequest) SetOpTenantId(v int64) *GetUsersRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUsersRequest) SetUserIdList(v []*string) *GetUsersRequest {
	s.UserIdList = v
	return s
}

func (s *GetUsersRequest) Validate() error {
	return dara.Validate(s)
}
