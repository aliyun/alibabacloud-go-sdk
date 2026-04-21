// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *CreateYikeWorkspaceResponseBody
	GetProductionId() *string
	SetRequestId(v string) *CreateYikeWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *CreateYikeWorkspaceResponseBody
	GetWorkspaceId() *string
}

type CreateYikeWorkspaceResponseBody struct {
	// example:
	//
	// ProductionId
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// llm-zna577pdximvztk5
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateYikeWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYikeWorkspaceResponseBody) GetProductionId() *string {
	return s.ProductionId
}

func (s *CreateYikeWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYikeWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateYikeWorkspaceResponseBody) SetProductionId(v string) *CreateYikeWorkspaceResponseBody {
	s.ProductionId = &v
	return s
}

func (s *CreateYikeWorkspaceResponseBody) SetRequestId(v string) *CreateYikeWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYikeWorkspaceResponseBody) SetWorkspaceId(v string) *CreateYikeWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateYikeWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
