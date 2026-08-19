// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateProjectResponseBody
	GetId() *int64
	SetProjectId(v int64) *CreateProjectResponseBody
	GetProjectId() *int64
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
}

type CreateProjectResponseBody struct {
	// The ID of the DataWorks workspace created by this operation.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Deprecated
	//
	// The ID of the DataWorks workspace created by this operation.
	//
	// Note: This field is deprecated. Use the Id field instead.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateProjectResponseBody) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) SetId(v int64) *CreateProjectResponseBody {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBody) SetProjectId(v int64) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
