// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveClusterNodesResponse
	GetStatusCode() *int32
}

type RemoveClusterNodesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveClusterNodesResponse) SetHeaders(v map[string]*string) *RemoveClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterNodesResponse) SetStatusCode(v int32) *RemoveClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterNodesResponse) Validate() error {
	return dara.Validate(s)
}
