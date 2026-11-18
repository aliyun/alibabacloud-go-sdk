// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DeleteProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *DeleteProjectResponseBody
	GetRequestId() *string
}

type DeleteProjectResponseBody struct {
	// The project ID.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectResponseBody) SetProjectId(v string) *DeleteProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
