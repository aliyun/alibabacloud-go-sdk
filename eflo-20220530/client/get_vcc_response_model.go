// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVccResponse
	GetStatusCode() *int32
	SetBody(v *GetVccResponseBody) *GetVccResponse
	GetBody() *GetVccResponseBody
}

type GetVccResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVccResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponse) GoString() string {
	return s.String()
}

func (s *GetVccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVccResponse) GetBody() *GetVccResponseBody {
	return s.Body
}

func (s *GetVccResponse) SetHeaders(v map[string]*string) *GetVccResponse {
	s.Headers = v
	return s
}

func (s *GetVccResponse) SetStatusCode(v int32) *GetVccResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVccResponse) SetBody(v *GetVccResponseBody) *GetVccResponse {
	s.Body = v
	return s
}

func (s *GetVccResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
