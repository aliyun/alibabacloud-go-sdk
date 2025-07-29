// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveNodePoolNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveNodePoolNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveNodePoolNodesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveNodePoolNodesResponseBody) *RemoveNodePoolNodesResponse
	GetBody() *RemoveNodePoolNodesResponseBody
}

type RemoveNodePoolNodesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveNodePoolNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveNodePoolNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveNodePoolNodesResponse) GoString() string {
	return s.String()
}

func (s *RemoveNodePoolNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveNodePoolNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveNodePoolNodesResponse) GetBody() *RemoveNodePoolNodesResponseBody {
	return s.Body
}

func (s *RemoveNodePoolNodesResponse) SetHeaders(v map[string]*string) *RemoveNodePoolNodesResponse {
	s.Headers = v
	return s
}

func (s *RemoveNodePoolNodesResponse) SetStatusCode(v int32) *RemoveNodePoolNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveNodePoolNodesResponse) SetBody(v *RemoveNodePoolNodesResponseBody) *RemoveNodePoolNodesResponse {
	s.Body = v
	return s
}

func (s *RemoveNodePoolNodesResponse) Validate() error {
	return dara.Validate(s)
}
