// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivyComputeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivyComputeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivyComputeTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivyComputeTokenResponseBody) *CreateLivyComputeTokenResponse
	GetBody() *CreateLivyComputeTokenResponseBody
}

type CreateLivyComputeTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivyComputeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivyComputeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivyComputeTokenResponse) GetBody() *CreateLivyComputeTokenResponseBody {
	return s.Body
}

func (s *CreateLivyComputeTokenResponse) SetHeaders(v map[string]*string) *CreateLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateLivyComputeTokenResponse) SetStatusCode(v int32) *CreateLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivyComputeTokenResponse) SetBody(v *CreateLivyComputeTokenResponseBody) *CreateLivyComputeTokenResponse {
	s.Body = v
	return s
}

func (s *CreateLivyComputeTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
