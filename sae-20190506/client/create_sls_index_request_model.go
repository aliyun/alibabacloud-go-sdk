// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlsIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *CreateSlsIndexRequest
	GetLogstore() *string
	SetProject(v string) *CreateSlsIndexRequest
	GetProject() *string
}

type CreateSlsIndexRequest struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project  *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s CreateSlsIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSlsIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateSlsIndexRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateSlsIndexRequest) GetProject() *string {
	return s.Project
}

func (s *CreateSlsIndexRequest) SetLogstore(v string) *CreateSlsIndexRequest {
	s.Logstore = &v
	return s
}

func (s *CreateSlsIndexRequest) SetProject(v string) *CreateSlsIndexRequest {
	s.Project = &v
	return s
}

func (s *CreateSlsIndexRequest) Validate() error {
	return dara.Validate(s)
}
