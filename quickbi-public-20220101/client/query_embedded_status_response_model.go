// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmbeddedStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEmbeddedStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEmbeddedStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryEmbeddedStatusResponseBody) *QueryEmbeddedStatusResponse
	GetBody() *QueryEmbeddedStatusResponseBody
}

type QueryEmbeddedStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmbeddedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEmbeddedStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEmbeddedStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEmbeddedStatusResponse) GetBody() *QueryEmbeddedStatusResponseBody {
	return s.Body
}

func (s *QueryEmbeddedStatusResponse) SetHeaders(v map[string]*string) *QueryEmbeddedStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryEmbeddedStatusResponse) SetStatusCode(v int32) *QueryEmbeddedStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmbeddedStatusResponse) SetBody(v *QueryEmbeddedStatusResponseBody) *QueryEmbeddedStatusResponse {
	s.Body = v
	return s
}

func (s *QueryEmbeddedStatusResponse) Validate() error {
	return dara.Validate(s)
}
