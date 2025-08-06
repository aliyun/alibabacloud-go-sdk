// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTotalStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTotalStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTotalStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetTotalStatisResponseBody) *GetTotalStatisResponse
	GetBody() *GetTotalStatisResponseBody
}

type GetTotalStatisResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTotalStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTotalStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTotalStatisResponse) GoString() string {
	return s.String()
}

func (s *GetTotalStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTotalStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTotalStatisResponse) GetBody() *GetTotalStatisResponseBody {
	return s.Body
}

func (s *GetTotalStatisResponse) SetHeaders(v map[string]*string) *GetTotalStatisResponse {
	s.Headers = v
	return s
}

func (s *GetTotalStatisResponse) SetStatusCode(v int32) *GetTotalStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTotalStatisResponse) SetBody(v *GetTotalStatisResponseBody) *GetTotalStatisResponse {
	s.Body = v
	return s
}

func (s *GetTotalStatisResponse) Validate() error {
	return dara.Validate(s)
}
