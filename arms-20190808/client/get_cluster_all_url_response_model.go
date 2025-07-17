// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAllUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterAllUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterAllUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterAllUrlResponseBody) *GetClusterAllUrlResponse
	GetBody() *GetClusterAllUrlResponseBody
}

type GetClusterAllUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterAllUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterAllUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAllUrlResponse) GoString() string {
	return s.String()
}

func (s *GetClusterAllUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterAllUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterAllUrlResponse) GetBody() *GetClusterAllUrlResponseBody {
	return s.Body
}

func (s *GetClusterAllUrlResponse) SetHeaders(v map[string]*string) *GetClusterAllUrlResponse {
	s.Headers = v
	return s
}

func (s *GetClusterAllUrlResponse) SetStatusCode(v int32) *GetClusterAllUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterAllUrlResponse) SetBody(v *GetClusterAllUrlResponseBody) *GetClusterAllUrlResponse {
	s.Body = v
	return s
}

func (s *GetClusterAllUrlResponse) Validate() error {
	return dara.Validate(s)
}
