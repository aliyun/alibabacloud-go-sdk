// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoginProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoginProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoginProfileResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoginProfileResponseBody) *CreateLoginProfileResponse
	GetBody() *CreateLoginProfileResponseBody
}

type CreateLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoginProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoginProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoginProfileResponse) GetBody() *CreateLoginProfileResponseBody {
	return s.Body
}

func (s *CreateLoginProfileResponse) SetHeaders(v map[string]*string) *CreateLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateLoginProfileResponse) SetStatusCode(v int32) *CreateLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoginProfileResponse) SetBody(v *CreateLoginProfileResponseBody) *CreateLoginProfileResponse {
	s.Body = v
	return s
}

func (s *CreateLoginProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
