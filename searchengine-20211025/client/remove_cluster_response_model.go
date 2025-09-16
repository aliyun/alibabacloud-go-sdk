// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveClusterResponse
	GetStatusCode() *int32
	SetBody(v *RemoveClusterResponseBody) *RemoveClusterResponse
	GetBody() *RemoveClusterResponseBody
}

type RemoveClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveClusterResponse) GetBody() *RemoveClusterResponseBody {
	return s.Body
}

func (s *RemoveClusterResponse) SetHeaders(v map[string]*string) *RemoveClusterResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterResponse) SetStatusCode(v int32) *RemoveClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterResponse) SetBody(v *RemoveClusterResponseBody) *RemoveClusterResponse {
	s.Body = v
	return s
}

func (s *RemoveClusterResponse) Validate() error {
	return dara.Validate(s)
}
