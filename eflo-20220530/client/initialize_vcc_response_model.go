// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeVccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeVccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeVccResponse
	GetStatusCode() *int32
	SetBody(v *InitializeVccResponseBody) *InitializeVccResponse
	GetBody() *InitializeVccResponseBody
}

type InitializeVccResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeVccResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeVccResponse) GoString() string {
	return s.String()
}

func (s *InitializeVccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeVccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeVccResponse) GetBody() *InitializeVccResponseBody {
	return s.Body
}

func (s *InitializeVccResponse) SetHeaders(v map[string]*string) *InitializeVccResponse {
	s.Headers = v
	return s
}

func (s *InitializeVccResponse) SetStatusCode(v int32) *InitializeVccResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeVccResponse) SetBody(v *InitializeVccResponseBody) *InitializeVccResponse {
	s.Body = v
	return s
}

func (s *InitializeVccResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
