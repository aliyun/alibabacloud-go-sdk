// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FailoverDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FailoverDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *FailoverDBClusterResponseBody) *FailoverDBClusterResponse
	GetBody() *FailoverDBClusterResponseBody
}

type FailoverDBClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterResponse) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FailoverDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FailoverDBClusterResponse) GetBody() *FailoverDBClusterResponseBody {
	return s.Body
}

func (s *FailoverDBClusterResponse) SetHeaders(v map[string]*string) *FailoverDBClusterResponse {
	s.Headers = v
	return s
}

func (s *FailoverDBClusterResponse) SetStatusCode(v int32) *FailoverDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDBClusterResponse) SetBody(v *FailoverDBClusterResponseBody) *FailoverDBClusterResponse {
	s.Body = v
	return s
}

func (s *FailoverDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
