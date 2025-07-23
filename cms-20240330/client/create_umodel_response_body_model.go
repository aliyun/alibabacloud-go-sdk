// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUmodelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUmodelResponseBody
	GetRequestId() *string
	SetWorkspace(v string) *CreateUmodelResponseBody
	GetWorkspace() *string
}

type CreateUmodelResponseBody struct {
	// example:
	//
	// 123-0F43-23423-AC43-34234
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateUmodelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUmodelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUmodelResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateUmodelResponseBody) SetRequestId(v string) *CreateUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUmodelResponseBody) SetWorkspace(v string) *CreateUmodelResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreateUmodelResponseBody) Validate() error {
	return dara.Validate(s)
}
