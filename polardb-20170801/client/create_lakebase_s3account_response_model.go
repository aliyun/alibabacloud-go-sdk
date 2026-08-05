// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakebaseS3AccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLakebaseS3AccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLakebaseS3AccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateLakebaseS3AccountResponseBody) *CreateLakebaseS3AccountResponse
	GetBody() *CreateLakebaseS3AccountResponseBody
}

type CreateLakebaseS3AccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLakebaseS3AccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLakebaseS3AccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLakebaseS3AccountResponse) GoString() string {
	return s.String()
}

func (s *CreateLakebaseS3AccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLakebaseS3AccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLakebaseS3AccountResponse) GetBody() *CreateLakebaseS3AccountResponseBody {
	return s.Body
}

func (s *CreateLakebaseS3AccountResponse) SetHeaders(v map[string]*string) *CreateLakebaseS3AccountResponse {
	s.Headers = v
	return s
}

func (s *CreateLakebaseS3AccountResponse) SetStatusCode(v int32) *CreateLakebaseS3AccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLakebaseS3AccountResponse) SetBody(v *CreateLakebaseS3AccountResponseBody) *CreateLakebaseS3AccountResponse {
	s.Body = v
	return s
}

func (s *CreateLakebaseS3AccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
