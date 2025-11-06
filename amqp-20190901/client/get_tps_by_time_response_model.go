// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTpsByTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTpsByTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTpsByTimeResponse
	GetStatusCode() *int32
	SetBody(v *GetTpsByTimeResponseBody) *GetTpsByTimeResponse
	GetBody() *GetTpsByTimeResponseBody
}

type GetTpsByTimeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTpsByTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTpsByTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTpsByTimeResponse) GoString() string {
	return s.String()
}

func (s *GetTpsByTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTpsByTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTpsByTimeResponse) GetBody() *GetTpsByTimeResponseBody {
	return s.Body
}

func (s *GetTpsByTimeResponse) SetHeaders(v map[string]*string) *GetTpsByTimeResponse {
	s.Headers = v
	return s
}

func (s *GetTpsByTimeResponse) SetStatusCode(v int32) *GetTpsByTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTpsByTimeResponse) SetBody(v *GetTpsByTimeResponseBody) *GetTpsByTimeResponse {
	s.Body = v
	return s
}

func (s *GetTpsByTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
