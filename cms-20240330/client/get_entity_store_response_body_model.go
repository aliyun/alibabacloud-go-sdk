// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetEntityStoreResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetEntityStoreResponseBody
	GetRequestId() *string
	SetWorkspaceName(v string) *GetEntityStoreResponseBody
	GetWorkspaceName() *string
}

type GetEntityStoreResponseBody struct {
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s GetEntityStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityStoreResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEntityStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEntityStoreResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetEntityStoreResponseBody) SetRegionId(v string) *GetEntityStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetEntityStoreResponseBody) SetRequestId(v string) *GetEntityStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityStoreResponseBody) SetWorkspaceName(v string) *GetEntityStoreResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetEntityStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
