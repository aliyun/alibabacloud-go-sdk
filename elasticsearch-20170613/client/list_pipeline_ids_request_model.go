// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *ListPipelineIdsRequest
	GetBody() *string
}

type ListPipelineIdsRequest struct {
	// example:
	//
	// {     "userName":"elastic",     "password":"xxxxxx" }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineIdsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsRequest) GetBody() *string {
	return s.Body
}

func (s *ListPipelineIdsRequest) SetBody(v string) *ListPipelineIdsRequest {
	s.Body = &v
	return s
}

func (s *ListPipelineIdsRequest) Validate() error {
	return dara.Validate(s)
}
