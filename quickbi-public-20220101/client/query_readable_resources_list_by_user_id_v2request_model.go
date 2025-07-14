// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReadableResourcesListByUserIdV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryReadableResourcesListByUserIdV2Request
	GetUserId() *string
	SetWorkType(v string) *QueryReadableResourcesListByUserIdV2Request
	GetWorkType() *string
	SetWorkspaceId(v string) *QueryReadableResourcesListByUserIdV2Request
	GetWorkspaceId() *string
}

type QueryReadableResourcesListByUserIdV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// asdas*********sdsddf
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// DATAFORM
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryReadableResourcesListByUserIdV2Request) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdV2Request) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdV2Request) GetUserId() *string {
	return s.UserId
}

func (s *QueryReadableResourcesListByUserIdV2Request) GetWorkType() *string {
	return s.WorkType
}

func (s *QueryReadableResourcesListByUserIdV2Request) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryReadableResourcesListByUserIdV2Request) SetUserId(v string) *QueryReadableResourcesListByUserIdV2Request {
	s.UserId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2Request) SetWorkType(v string) *QueryReadableResourcesListByUserIdV2Request {
	s.WorkType = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2Request) SetWorkspaceId(v string) *QueryReadableResourcesListByUserIdV2Request {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2Request) Validate() error {
	return dara.Validate(s)
}
