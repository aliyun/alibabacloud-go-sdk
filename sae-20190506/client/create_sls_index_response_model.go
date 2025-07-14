// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlsIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSlsIndexResponse
	GetRequestId() *string
	SetLogStore(v string) *CreateSlsIndexResponse
	GetLogStore() *string
	SetProject(v string) *CreateSlsIndexResponse
	GetProject() *string
}

type CreateSlsIndexResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LogStore  *string `json:"logStore,omitempty" xml:"logStore,omitempty"`
	Project   *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s CreateSlsIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlsIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateSlsIndexResponse) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlsIndexResponse) GetLogStore() *string {
	return s.LogStore
}

func (s *CreateSlsIndexResponse) GetProject() *string {
	return s.Project
}

func (s *CreateSlsIndexResponse) SetRequestId(v string) *CreateSlsIndexResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSlsIndexResponse) SetLogStore(v string) *CreateSlsIndexResponse {
	s.LogStore = &v
	return s
}

func (s *CreateSlsIndexResponse) SetProject(v string) *CreateSlsIndexResponse {
	s.Project = &v
	return s
}

func (s *CreateSlsIndexResponse) Validate() error {
	return dara.Validate(s)
}
