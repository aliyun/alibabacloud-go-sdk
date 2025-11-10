// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRequestLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRequestLogResponse
	GetStatusCode() *int32
	SetBody(v *GetRequestLogResponseBody) *GetRequestLogResponse
	GetBody() *GetRequestLogResponseBody
}

type GetRequestLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRequestLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRequestLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponse) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRequestLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRequestLogResponse) GetBody() *GetRequestLogResponseBody {
	return s.Body
}

func (s *GetRequestLogResponse) SetHeaders(v map[string]*string) *GetRequestLogResponse {
	s.Headers = v
	return s
}

func (s *GetRequestLogResponse) SetStatusCode(v int32) *GetRequestLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRequestLogResponse) SetBody(v *GetRequestLogResponseBody) *GetRequestLogResponse {
	s.Body = v
	return s
}

func (s *GetRequestLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
