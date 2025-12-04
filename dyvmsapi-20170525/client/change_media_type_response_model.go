// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMediaTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeMediaTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeMediaTypeResponse
	GetStatusCode() *int32
	SetBody(v *ChangeMediaTypeResponseBody) *ChangeMediaTypeResponse
	GetBody() *ChangeMediaTypeResponseBody
}

type ChangeMediaTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMediaTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMediaTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeMediaTypeResponse) GoString() string {
	return s.String()
}

func (s *ChangeMediaTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeMediaTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeMediaTypeResponse) GetBody() *ChangeMediaTypeResponseBody {
	return s.Body
}

func (s *ChangeMediaTypeResponse) SetHeaders(v map[string]*string) *ChangeMediaTypeResponse {
	s.Headers = v
	return s
}

func (s *ChangeMediaTypeResponse) SetStatusCode(v int32) *ChangeMediaTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMediaTypeResponse) SetBody(v *ChangeMediaTypeResponseBody) *ChangeMediaTypeResponse {
	s.Body = v
	return s
}

func (s *ChangeMediaTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
