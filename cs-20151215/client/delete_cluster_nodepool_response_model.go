// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodepoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClusterNodepoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClusterNodepoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClusterNodepoolResponseBody) *DeleteClusterNodepoolResponse
	GetBody() *DeleteClusterNodepoolResponseBody
}

type DeleteClusterNodepoolResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterNodepoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterNodepoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodepoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodepoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClusterNodepoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClusterNodepoolResponse) GetBody() *DeleteClusterNodepoolResponseBody {
	return s.Body
}

func (s *DeleteClusterNodepoolResponse) SetHeaders(v map[string]*string) *DeleteClusterNodepoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodepoolResponse) SetStatusCode(v int32) *DeleteClusterNodepoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterNodepoolResponse) SetBody(v *DeleteClusterNodepoolResponseBody) *DeleteClusterNodepoolResponse {
	s.Body = v
	return s
}

func (s *DeleteClusterNodepoolResponse) Validate() error {
	return dara.Validate(s)
}
