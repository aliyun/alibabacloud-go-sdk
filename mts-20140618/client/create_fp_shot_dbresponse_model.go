// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFpShotDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFpShotDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFpShotDBResponse
	GetStatusCode() *int32
	SetBody(v *CreateFpShotDBResponseBody) *CreateFpShotDBResponse
	GetBody() *CreateFpShotDBResponseBody
}

type CreateFpShotDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFpShotDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFpShotDBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFpShotDBResponse) GoString() string {
	return s.String()
}

func (s *CreateFpShotDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFpShotDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFpShotDBResponse) GetBody() *CreateFpShotDBResponseBody {
	return s.Body
}

func (s *CreateFpShotDBResponse) SetHeaders(v map[string]*string) *CreateFpShotDBResponse {
	s.Headers = v
	return s
}

func (s *CreateFpShotDBResponse) SetStatusCode(v int32) *CreateFpShotDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFpShotDBResponse) SetBody(v *CreateFpShotDBResponseBody) *CreateFpShotDBResponse {
	s.Body = v
	return s
}

func (s *CreateFpShotDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
