// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSourceLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSourceLocationResponse
	GetStatusCode() *int32
	SetBody(v *GetSourceLocationResponseBody) *GetSourceLocationResponse
	GetBody() *GetSourceLocationResponseBody
}

type GetSourceLocationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSourceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSourceLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSourceLocationResponse) GoString() string {
	return s.String()
}

func (s *GetSourceLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSourceLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSourceLocationResponse) GetBody() *GetSourceLocationResponseBody {
	return s.Body
}

func (s *GetSourceLocationResponse) SetHeaders(v map[string]*string) *GetSourceLocationResponse {
	s.Headers = v
	return s
}

func (s *GetSourceLocationResponse) SetStatusCode(v int32) *GetSourceLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourceLocationResponse) SetBody(v *GetSourceLocationResponseBody) *GetSourceLocationResponse {
	s.Body = v
	return s
}

func (s *GetSourceLocationResponse) Validate() error {
	return dara.Validate(s)
}
