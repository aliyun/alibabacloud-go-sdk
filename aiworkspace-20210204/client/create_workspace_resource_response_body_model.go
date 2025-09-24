// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkspaceResourceResponseBody
	GetRequestId() *string
	SetResources(v []*CreateWorkspaceResourceResponseBodyResources) *CreateWorkspaceResourceResponseBody
	GetResources() []*CreateWorkspaceResourceResponseBodyResources
	SetTotalCount(v int64) *CreateWorkspaceResourceResponseBody
	GetTotalCount() *int64
}

type CreateWorkspaceResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	Resources []*CreateWorkspaceResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of resources.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateWorkspaceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResourceResponseBody) GetResources() []*CreateWorkspaceResourceResponseBodyResources {
	return s.Resources
}

func (s *CreateWorkspaceResourceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *CreateWorkspaceResourceResponseBody) SetRequestId(v string) *CreateWorkspaceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResourceResponseBody) SetResources(v []*CreateWorkspaceResourceResponseBodyResources) *CreateWorkspaceResourceResponseBody {
	s.Resources = v
	return s
}

func (s *CreateWorkspaceResourceResponseBody) SetTotalCount(v int64) *CreateWorkspaceResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateWorkspaceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateWorkspaceResourceResponseBodyResources struct {
	// The resource ID.
	//
	// example:
	//
	// 1234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateWorkspaceResourceResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceResponseBodyResources) GetId() *string {
	return s.Id
}

func (s *CreateWorkspaceResourceResponseBodyResources) SetId(v string) *CreateWorkspaceResourceResponseBodyResources {
	s.Id = &v
	return s
}

func (s *CreateWorkspaceResourceResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
