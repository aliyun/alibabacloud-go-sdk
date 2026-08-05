// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakebaseS3AccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLakebaseS3AccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLakebaseS3AccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLakebaseS3AccountResponseBody) *DeleteLakebaseS3AccountResponse
	GetBody() *DeleteLakebaseS3AccountResponseBody
}

type DeleteLakebaseS3AccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLakebaseS3AccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLakebaseS3AccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakebaseS3AccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteLakebaseS3AccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLakebaseS3AccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLakebaseS3AccountResponse) GetBody() *DeleteLakebaseS3AccountResponseBody {
	return s.Body
}

func (s *DeleteLakebaseS3AccountResponse) SetHeaders(v map[string]*string) *DeleteLakebaseS3AccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteLakebaseS3AccountResponse) SetStatusCode(v int32) *DeleteLakebaseS3AccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLakebaseS3AccountResponse) SetBody(v *DeleteLakebaseS3AccountResponseBody) *DeleteLakebaseS3AccountResponse {
	s.Body = v
	return s
}

func (s *DeleteLakebaseS3AccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
