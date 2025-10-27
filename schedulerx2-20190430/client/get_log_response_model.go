// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogResponse
	GetStatusCode() *int32
	SetBody(v *GetLogResponseBody) *GetLogResponse
	GetBody() *GetLogResponseBody
}

type GetLogResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogResponse) GoString() string {
	return s.String()
}

func (s *GetLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogResponse) GetBody() *GetLogResponseBody {
	return s.Body
}

func (s *GetLogResponse) SetHeaders(v map[string]*string) *GetLogResponse {
	s.Headers = v
	return s
}

func (s *GetLogResponse) SetStatusCode(v int32) *GetLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogResponse) SetBody(v *GetLogResponseBody) *GetLogResponse {
	s.Body = v
	return s
}

func (s *GetLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
