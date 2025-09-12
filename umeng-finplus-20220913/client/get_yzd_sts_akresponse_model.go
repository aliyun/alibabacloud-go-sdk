// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYzdStsAKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYzdStsAKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYzdStsAKResponse
	GetStatusCode() *int32
	SetBody(v *GetYzdStsAKResponseBody) *GetYzdStsAKResponse
	GetBody() *GetYzdStsAKResponseBody
}

type GetYzdStsAKResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYzdStsAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYzdStsAKResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYzdStsAKResponse) GoString() string {
	return s.String()
}

func (s *GetYzdStsAKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYzdStsAKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYzdStsAKResponse) GetBody() *GetYzdStsAKResponseBody {
	return s.Body
}

func (s *GetYzdStsAKResponse) SetHeaders(v map[string]*string) *GetYzdStsAKResponse {
	s.Headers = v
	return s
}

func (s *GetYzdStsAKResponse) SetStatusCode(v int32) *GetYzdStsAKResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYzdStsAKResponse) SetBody(v *GetYzdStsAKResponseBody) *GetYzdStsAKResponse {
	s.Body = v
	return s
}

func (s *GetYzdStsAKResponse) Validate() error {
	return dara.Validate(s)
}
