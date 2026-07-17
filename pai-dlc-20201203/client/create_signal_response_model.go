// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSignalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSignalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSignalResponse
	GetStatusCode() *int32
	SetBody(v *CreateSignalResponseBody) *CreateSignalResponse
	GetBody() *CreateSignalResponseBody
}

type CreateSignalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSignalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSignalResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSignalResponse) GoString() string {
	return s.String()
}

func (s *CreateSignalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSignalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSignalResponse) GetBody() *CreateSignalResponseBody {
	return s.Body
}

func (s *CreateSignalResponse) SetHeaders(v map[string]*string) *CreateSignalResponse {
	s.Headers = v
	return s
}

func (s *CreateSignalResponse) SetStatusCode(v int32) *CreateSignalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSignalResponse) SetBody(v *CreateSignalResponseBody) *CreateSignalResponse {
	s.Body = v
	return s
}

func (s *CreateSignalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
