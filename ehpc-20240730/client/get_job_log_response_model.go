// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobLogResponse
	GetStatusCode() *int32
	SetBody(v *GetJobLogResponseBody) *GetJobLogResponse
	GetBody() *GetJobLogResponseBody
}

type GetJobLogResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobLogResponse) GetBody() *GetJobLogResponseBody {
	return s.Body
}

func (s *GetJobLogResponse) SetHeaders(v map[string]*string) *GetJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetJobLogResponse) SetStatusCode(v int32) *GetJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobLogResponse) SetBody(v *GetJobLogResponseBody) *GetJobLogResponse {
	s.Body = v
	return s
}

func (s *GetJobLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
