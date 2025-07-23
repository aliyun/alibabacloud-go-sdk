// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUmodelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUmodelResponseBody
	GetRequestId() *string
	SetWorkspace(v string) *UpdateUmodelResponseBody
	GetWorkspace() *string
}

type UpdateUmodelResponseBody struct {
	// example:
	//
	// 234324-123-123-123-23423
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateUmodelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUmodelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUmodelResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateUmodelResponseBody) SetRequestId(v string) *UpdateUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUmodelResponseBody) SetWorkspace(v string) *UpdateUmodelResponseBody {
	s.Workspace = &v
	return s
}

func (s *UpdateUmodelResponseBody) Validate() error {
	return dara.Validate(s)
}
