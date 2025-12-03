// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *ListCustomAgentToolsRequest
	GetApiId() *string
}

type ListCustomAgentToolsRequest struct {
	// example:
	//
	// app-iBuGU1VxEY42zrQRQfNA****
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
}

func (s ListCustomAgentToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentToolsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentToolsRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ListCustomAgentToolsRequest) SetApiId(v string) *ListCustomAgentToolsRequest {
	s.ApiId = &v
	return s
}

func (s *ListCustomAgentToolsRequest) Validate() error {
	return dara.Validate(s)
}
