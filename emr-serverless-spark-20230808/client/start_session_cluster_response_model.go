// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSessionClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSessionClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSessionClusterResponse
	GetStatusCode() *int32
	SetBody(v *StartSessionClusterResponseBody) *StartSessionClusterResponse
	GetBody() *StartSessionClusterResponseBody
}

type StartSessionClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSessionClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *StartSessionClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSessionClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSessionClusterResponse) GetBody() *StartSessionClusterResponseBody {
	return s.Body
}

func (s *StartSessionClusterResponse) SetHeaders(v map[string]*string) *StartSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *StartSessionClusterResponse) SetStatusCode(v int32) *StartSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSessionClusterResponse) SetBody(v *StartSessionClusterResponseBody) *StartSessionClusterResponse {
	s.Body = v
	return s
}

func (s *StartSessionClusterResponse) Validate() error {
	return dara.Validate(s)
}
