// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlsResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetLogStore(v string) *CreateSlsResourceResponse
	GetLogStore() *string
	SetProject(v string) *CreateSlsResourceResponse
	GetProject() *string
	SetRequestId(v string) *CreateSlsResourceResponse
	GetRequestId() *string
}

type CreateSlsResourceResponse struct {
	LogStore  *string `json:"logStore,omitempty" xml:"logStore,omitempty"`
	Project   *string `json:"project,omitempty" xml:"project,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSlsResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlsResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateSlsResourceResponse) GetLogStore() *string {
	return s.LogStore
}

func (s *CreateSlsResourceResponse) GetProject() *string {
	return s.Project
}

func (s *CreateSlsResourceResponse) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlsResourceResponse) SetLogStore(v string) *CreateSlsResourceResponse {
	s.LogStore = &v
	return s
}

func (s *CreateSlsResourceResponse) SetProject(v string) *CreateSlsResourceResponse {
	s.Project = &v
	return s
}

func (s *CreateSlsResourceResponse) SetRequestId(v string) *CreateSlsResourceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSlsResourceResponse) Validate() error {
	return dara.Validate(s)
}
