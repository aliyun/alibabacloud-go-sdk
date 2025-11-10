// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOwnRequestLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOwnRequestLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOwnRequestLogResponse
	GetStatusCode() *int32
	SetBody(v *GetOwnRequestLogResponseBody) *GetOwnRequestLogResponse
	GetBody() *GetOwnRequestLogResponseBody
}

type GetOwnRequestLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOwnRequestLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOwnRequestLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponse) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOwnRequestLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOwnRequestLogResponse) GetBody() *GetOwnRequestLogResponseBody {
	return s.Body
}

func (s *GetOwnRequestLogResponse) SetHeaders(v map[string]*string) *GetOwnRequestLogResponse {
	s.Headers = v
	return s
}

func (s *GetOwnRequestLogResponse) SetStatusCode(v int32) *GetOwnRequestLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOwnRequestLogResponse) SetBody(v *GetOwnRequestLogResponseBody) *GetOwnRequestLogResponse {
	s.Body = v
	return s
}

func (s *GetOwnRequestLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
