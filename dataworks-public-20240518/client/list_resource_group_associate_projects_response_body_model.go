// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupAssociateProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectIdList(v []*int64) *ListResourceGroupAssociateProjectsResponseBody
	GetProjectIdList() []*int64
	SetRequestId(v string) *ListResourceGroupAssociateProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourceGroupAssociateProjectsResponseBody
	GetSuccess() *bool
}

type ListResourceGroupAssociateProjectsResponseBody struct {
	// The list of workspace IDs.
	ProjectIdList []*int64 `json:"ProjectIdList,omitempty" xml:"ProjectIdList,omitempty" type:"Repeated"`
	// The request ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourceGroupAssociateProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupAssociateProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupAssociateProjectsResponseBody) GetProjectIdList() []*int64 {
	return s.ProjectIdList
}

func (s *ListResourceGroupAssociateProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupAssociateProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceGroupAssociateProjectsResponseBody) SetProjectIdList(v []*int64) *ListResourceGroupAssociateProjectsResponseBody {
	s.ProjectIdList = v
	return s
}

func (s *ListResourceGroupAssociateProjectsResponseBody) SetRequestId(v string) *ListResourceGroupAssociateProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupAssociateProjectsResponseBody) SetSuccess(v bool) *ListResourceGroupAssociateProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceGroupAssociateProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
