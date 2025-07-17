// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetColumnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetColumnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetColumnResponse
	GetStatusCode() *int32
	SetBody(v *GetColumnResponseBody) *GetColumnResponse
	GetBody() *GetColumnResponseBody
}

type GetColumnResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetColumnResponse) String() string {
	return dara.Prettify(s)
}

func (s GetColumnResponse) GoString() string {
	return s.String()
}

func (s *GetColumnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetColumnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetColumnResponse) GetBody() *GetColumnResponseBody {
	return s.Body
}

func (s *GetColumnResponse) SetHeaders(v map[string]*string) *GetColumnResponse {
	s.Headers = v
	return s
}

func (s *GetColumnResponse) SetStatusCode(v int32) *GetColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetColumnResponse) SetBody(v *GetColumnResponseBody) *GetColumnResponse {
	s.Body = v
	return s
}

func (s *GetColumnResponse) Validate() error {
	return dara.Validate(s)
}
