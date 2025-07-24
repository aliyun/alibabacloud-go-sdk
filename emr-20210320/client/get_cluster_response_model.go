// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterResponseBody) *GetClusterResponse
	GetBody() *GetClusterResponseBody
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterResponse) GetBody() *GetClusterResponseBody {
	return s.Body
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

func (s *GetClusterResponse) Validate() error {
	return dara.Validate(s)
}
