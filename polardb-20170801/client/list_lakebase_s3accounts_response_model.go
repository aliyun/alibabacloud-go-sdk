// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLakebaseS3AccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLakebaseS3AccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLakebaseS3AccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListLakebaseS3AccountsResponseBody) *ListLakebaseS3AccountsResponse
	GetBody() *ListLakebaseS3AccountsResponseBody
}

type ListLakebaseS3AccountsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLakebaseS3AccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLakebaseS3AccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLakebaseS3AccountsResponse) GoString() string {
	return s.String()
}

func (s *ListLakebaseS3AccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLakebaseS3AccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLakebaseS3AccountsResponse) GetBody() *ListLakebaseS3AccountsResponseBody {
	return s.Body
}

func (s *ListLakebaseS3AccountsResponse) SetHeaders(v map[string]*string) *ListLakebaseS3AccountsResponse {
	s.Headers = v
	return s
}

func (s *ListLakebaseS3AccountsResponse) SetStatusCode(v int32) *ListLakebaseS3AccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLakebaseS3AccountsResponse) SetBody(v *ListLakebaseS3AccountsResponseBody) *ListLakebaseS3AccountsResponse {
	s.Body = v
	return s
}

func (s *ListLakebaseS3AccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
