// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSessionClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSessionClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetSessionClusterResponseBody) *GetSessionClusterResponse
	GetBody() *GetSessionClusterResponseBody
}

type GetSessionClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSessionClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSessionClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSessionClusterResponse) GetBody() *GetSessionClusterResponseBody {
	return s.Body
}

func (s *GetSessionClusterResponse) SetHeaders(v map[string]*string) *GetSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *GetSessionClusterResponse) SetStatusCode(v int32) *GetSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSessionClusterResponse) SetBody(v *GetSessionClusterResponseBody) *GetSessionClusterResponse {
	s.Body = v
	return s
}

func (s *GetSessionClusterResponse) Validate() error {
	return dara.Validate(s)
}
