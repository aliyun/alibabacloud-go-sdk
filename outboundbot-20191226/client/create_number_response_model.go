// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNumberResponse
	GetStatusCode() *int32
	SetBody(v *CreateNumberResponseBody) *CreateNumberResponse
	GetBody() *CreateNumberResponseBody
}

type CreateNumberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNumberResponse) GoString() string {
	return s.String()
}

func (s *CreateNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNumberResponse) GetBody() *CreateNumberResponseBody {
	return s.Body
}

func (s *CreateNumberResponse) SetHeaders(v map[string]*string) *CreateNumberResponse {
	s.Headers = v
	return s
}

func (s *CreateNumberResponse) SetStatusCode(v int32) *CreateNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNumberResponse) SetBody(v *CreateNumberResponseBody) *CreateNumberResponse {
	s.Body = v
	return s
}

func (s *CreateNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
