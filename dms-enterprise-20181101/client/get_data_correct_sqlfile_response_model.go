// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectSQLFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCorrectSQLFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCorrectSQLFileResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCorrectSQLFileResponseBody) *GetDataCorrectSQLFileResponse
	GetBody() *GetDataCorrectSQLFileResponseBody
}

type GetDataCorrectSQLFileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCorrectSQLFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCorrectSQLFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectSQLFileResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectSQLFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCorrectSQLFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCorrectSQLFileResponse) GetBody() *GetDataCorrectSQLFileResponseBody {
	return s.Body
}

func (s *GetDataCorrectSQLFileResponse) SetHeaders(v map[string]*string) *GetDataCorrectSQLFileResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectSQLFileResponse) SetStatusCode(v int32) *GetDataCorrectSQLFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCorrectSQLFileResponse) SetBody(v *GetDataCorrectSQLFileResponseBody) *GetDataCorrectSQLFileResponse {
	s.Body = v
	return s
}

func (s *GetDataCorrectSQLFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
