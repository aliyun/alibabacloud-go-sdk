// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCorrectPreCheckSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCorrectPreCheckSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCorrectPreCheckSQLResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCorrectPreCheckSQLResponseBody) *ListDataCorrectPreCheckSQLResponse
	GetBody() *ListDataCorrectPreCheckSQLResponseBody
}

type ListDataCorrectPreCheckSQLResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCorrectPreCheckSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCorrectPreCheckSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLResponse) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCorrectPreCheckSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCorrectPreCheckSQLResponse) GetBody() *ListDataCorrectPreCheckSQLResponseBody {
	return s.Body
}

func (s *ListDataCorrectPreCheckSQLResponse) SetHeaders(v map[string]*string) *ListDataCorrectPreCheckSQLResponse {
	s.Headers = v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponse) SetStatusCode(v int32) *ListDataCorrectPreCheckSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponse) SetBody(v *ListDataCorrectPreCheckSQLResponseBody) *ListDataCorrectPreCheckSQLResponse {
	s.Body = v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
