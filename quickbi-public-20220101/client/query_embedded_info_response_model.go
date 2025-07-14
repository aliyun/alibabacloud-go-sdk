// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmbeddedInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEmbeddedInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEmbeddedInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryEmbeddedInfoResponseBody) *QueryEmbeddedInfoResponse
	GetBody() *QueryEmbeddedInfoResponseBody
}

type QueryEmbeddedInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmbeddedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEmbeddedInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEmbeddedInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEmbeddedInfoResponse) GetBody() *QueryEmbeddedInfoResponseBody {
	return s.Body
}

func (s *QueryEmbeddedInfoResponse) SetHeaders(v map[string]*string) *QueryEmbeddedInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryEmbeddedInfoResponse) SetStatusCode(v int32) *QueryEmbeddedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmbeddedInfoResponse) SetBody(v *QueryEmbeddedInfoResponseBody) *QueryEmbeddedInfoResponse {
	s.Body = v
	return s
}

func (s *QueryEmbeddedInfoResponse) Validate() error {
	return dara.Validate(s)
}
