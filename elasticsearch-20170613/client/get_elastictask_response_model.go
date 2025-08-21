// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElastictaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetElastictaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetElastictaskResponse
	GetStatusCode() *int32
	SetBody(v *GetElastictaskResponseBody) *GetElastictaskResponse
	GetBody() *GetElastictaskResponseBody
}

type GetElastictaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElastictaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElastictaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetElastictaskResponse) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetElastictaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetElastictaskResponse) GetBody() *GetElastictaskResponseBody {
	return s.Body
}

func (s *GetElastictaskResponse) SetHeaders(v map[string]*string) *GetElastictaskResponse {
	s.Headers = v
	return s
}

func (s *GetElastictaskResponse) SetStatusCode(v int32) *GetElastictaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElastictaskResponse) SetBody(v *GetElastictaskResponseBody) *GetElastictaskResponse {
	s.Body = v
	return s
}

func (s *GetElastictaskResponse) Validate() error {
	return dara.Validate(s)
}
