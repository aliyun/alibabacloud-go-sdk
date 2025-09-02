// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectIds(v []*int64) *ListProjectIdsResponseBody
	GetProjectIds() []*int64
	SetRequestId(v string) *ListProjectIdsResponseBody
	GetRequestId() *string
}

type ListProjectIdsResponseBody struct {
	// The IDs of the DataWorks workspaces. The IDs of the workspaces on which the desired Alibaba Cloud account has permissions were returned.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0b57ff7216278945532771749d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectIdsResponseBody) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListProjectIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectIdsResponseBody) SetProjectIds(v []*int64) *ListProjectIdsResponseBody {
	s.ProjectIds = v
	return s
}

func (s *ListProjectIdsResponseBody) SetRequestId(v string) *ListProjectIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectIdsResponseBody) Validate() error {
	return dara.Validate(s)
}
