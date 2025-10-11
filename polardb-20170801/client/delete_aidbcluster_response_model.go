// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIDBClusterResponseBody) *DeleteAIDBClusterResponse
	GetBody() *DeleteAIDBClusterResponseBody
}

type DeleteAIDBClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIDBClusterResponse) GetBody() *DeleteAIDBClusterResponseBody {
	return s.Body
}

func (s *DeleteAIDBClusterResponse) SetHeaders(v map[string]*string) *DeleteAIDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIDBClusterResponse) SetStatusCode(v int32) *DeleteAIDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIDBClusterResponse) SetBody(v *DeleteAIDBClusterResponseBody) *DeleteAIDBClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteAIDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
