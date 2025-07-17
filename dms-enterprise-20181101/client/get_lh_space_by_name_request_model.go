// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhSpaceByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceName(v string) *GetLhSpaceByNameRequest
	GetSpaceName() *string
	SetTid(v int64) *GetLhSpaceByNameRequest
	GetTid() *int64
}

type GetLhSpaceByNameRequest struct {
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_space
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3000
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetLhSpaceByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLhSpaceByNameRequest) GoString() string {
	return s.String()
}

func (s *GetLhSpaceByNameRequest) GetSpaceName() *string {
	return s.SpaceName
}

func (s *GetLhSpaceByNameRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetLhSpaceByNameRequest) SetSpaceName(v string) *GetLhSpaceByNameRequest {
	s.SpaceName = &v
	return s
}

func (s *GetLhSpaceByNameRequest) SetTid(v int64) *GetLhSpaceByNameRequest {
	s.Tid = &v
	return s
}

func (s *GetLhSpaceByNameRequest) Validate() error {
	return dara.Validate(s)
}
