// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectRollbackFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCorrectRollbackFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCorrectRollbackFileResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCorrectRollbackFileResponseBody) *GetDataCorrectRollbackFileResponse
	GetBody() *GetDataCorrectRollbackFileResponseBody
}

type GetDataCorrectRollbackFileResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCorrectRollbackFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCorrectRollbackFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectRollbackFileResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectRollbackFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCorrectRollbackFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCorrectRollbackFileResponse) GetBody() *GetDataCorrectRollbackFileResponseBody {
	return s.Body
}

func (s *GetDataCorrectRollbackFileResponse) SetHeaders(v map[string]*string) *GetDataCorrectRollbackFileResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectRollbackFileResponse) SetStatusCode(v int32) *GetDataCorrectRollbackFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCorrectRollbackFileResponse) SetBody(v *GetDataCorrectRollbackFileResponseBody) *GetDataCorrectRollbackFileResponse {
	s.Body = v
	return s
}

func (s *GetDataCorrectRollbackFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
