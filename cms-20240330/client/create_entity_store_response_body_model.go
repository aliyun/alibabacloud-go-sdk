// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEntityStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEntityStoreResponseBody
	GetRequestId() *string
	SetWorkspaceName(v string) *CreateEntityStoreResponseBody
	GetWorkspaceName() *string
}

type CreateEntityStoreResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s CreateEntityStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEntityStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEntityStoreResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateEntityStoreResponseBody) SetRequestId(v string) *CreateEntityStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEntityStoreResponseBody) SetWorkspaceName(v string) *CreateEntityStoreResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *CreateEntityStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
