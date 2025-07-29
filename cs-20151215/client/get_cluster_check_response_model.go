// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterCheckResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterCheckResponseBody) *GetClusterCheckResponse
	GetBody() *GetClusterCheckResponseBody
}

type GetClusterCheckResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckResponse) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterCheckResponse) GetBody() *GetClusterCheckResponseBody {
	return s.Body
}

func (s *GetClusterCheckResponse) SetHeaders(v map[string]*string) *GetClusterCheckResponse {
	s.Headers = v
	return s
}

func (s *GetClusterCheckResponse) SetStatusCode(v int32) *GetClusterCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterCheckResponse) SetBody(v *GetClusterCheckResponseBody) *GetClusterCheckResponse {
	s.Body = v
	return s
}

func (s *GetClusterCheckResponse) Validate() error {
	return dara.Validate(s)
}
