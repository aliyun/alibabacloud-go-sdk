// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWorkspaceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RemoveWorkspaceUserResponseBody
	GetData() *bool
	SetRequestId(v string) *RemoveWorkspaceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveWorkspaceUserResponseBody
	GetSuccess() *bool
}

type RemoveWorkspaceUserResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E5EE2B9E-2F95-57FA-B284-CB441CEE49D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveWorkspaceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveWorkspaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWorkspaceUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveWorkspaceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveWorkspaceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveWorkspaceUserResponseBody) SetData(v bool) *RemoveWorkspaceUserResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveWorkspaceUserResponseBody) SetRequestId(v string) *RemoveWorkspaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveWorkspaceUserResponseBody) SetSuccess(v bool) *RemoveWorkspaceUserResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveWorkspaceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
