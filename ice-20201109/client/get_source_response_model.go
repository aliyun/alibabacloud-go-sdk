// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetSourceResponseBody) *GetSourceResponse
	GetBody() *GetSourceResponseBody
}

type GetSourceResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponse) GoString() string {
	return s.String()
}

func (s *GetSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSourceResponse) GetBody() *GetSourceResponseBody {
	return s.Body
}

func (s *GetSourceResponse) SetHeaders(v map[string]*string) *GetSourceResponse {
	s.Headers = v
	return s
}

func (s *GetSourceResponse) SetStatusCode(v int32) *GetSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourceResponse) SetBody(v *GetSourceResponseBody) *GetSourceResponse {
	s.Body = v
	return s
}

func (s *GetSourceResponse) Validate() error {
	return dara.Validate(s)
}
