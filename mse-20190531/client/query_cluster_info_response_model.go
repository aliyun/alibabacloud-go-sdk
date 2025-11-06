// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryClusterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryClusterInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryClusterInfoResponseBody) *QueryClusterInfoResponse
	GetBody() *QueryClusterInfoResponseBody
}

type QueryClusterInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryClusterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryClusterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryClusterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryClusterInfoResponse) GetBody() *QueryClusterInfoResponseBody {
	return s.Body
}

func (s *QueryClusterInfoResponse) SetHeaders(v map[string]*string) *QueryClusterInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterInfoResponse) SetStatusCode(v int32) *QueryClusterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClusterInfoResponse) SetBody(v *QueryClusterInfoResponseBody) *QueryClusterInfoResponse {
	s.Body = v
	return s
}

func (s *QueryClusterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
