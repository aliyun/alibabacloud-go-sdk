// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoFieldUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoFieldUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoFieldUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoFieldUrlResponseBody) *GetVideoFieldUrlResponse
	GetBody() *GetVideoFieldUrlResponseBody
}

type GetVideoFieldUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoFieldUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoFieldUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoFieldUrlResponse) GoString() string {
	return s.String()
}

func (s *GetVideoFieldUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoFieldUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoFieldUrlResponse) GetBody() *GetVideoFieldUrlResponseBody {
	return s.Body
}

func (s *GetVideoFieldUrlResponse) SetHeaders(v map[string]*string) *GetVideoFieldUrlResponse {
	s.Headers = v
	return s
}

func (s *GetVideoFieldUrlResponse) SetStatusCode(v int32) *GetVideoFieldUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoFieldUrlResponse) SetBody(v *GetVideoFieldUrlResponseBody) *GetVideoFieldUrlResponse {
	s.Body = v
	return s
}

func (s *GetVideoFieldUrlResponse) Validate() error {
	return dara.Validate(s)
}
