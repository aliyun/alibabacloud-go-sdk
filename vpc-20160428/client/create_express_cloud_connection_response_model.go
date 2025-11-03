// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressCloudConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressCloudConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressCloudConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressCloudConnectionResponseBody) *CreateExpressCloudConnectionResponse
	GetBody() *CreateExpressCloudConnectionResponseBody
}

type CreateExpressCloudConnectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressCloudConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressCloudConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressCloudConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressCloudConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressCloudConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressCloudConnectionResponse) GetBody() *CreateExpressCloudConnectionResponseBody {
	return s.Body
}

func (s *CreateExpressCloudConnectionResponse) SetHeaders(v map[string]*string) *CreateExpressCloudConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressCloudConnectionResponse) SetStatusCode(v int32) *CreateExpressCloudConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressCloudConnectionResponse) SetBody(v *CreateExpressCloudConnectionResponseBody) *CreateExpressCloudConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateExpressCloudConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
