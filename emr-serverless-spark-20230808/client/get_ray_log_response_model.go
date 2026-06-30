// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRayLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRayLogResponse
	GetStatusCode() *int32
	SetBody(v *GetRayLogResponseBody) *GetRayLogResponse
	GetBody() *GetRayLogResponseBody
}

type GetRayLogResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRayLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRayLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRayLogResponse) GoString() string {
	return s.String()
}

func (s *GetRayLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRayLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRayLogResponse) GetBody() *GetRayLogResponseBody {
	return s.Body
}

func (s *GetRayLogResponse) SetHeaders(v map[string]*string) *GetRayLogResponse {
	s.Headers = v
	return s
}

func (s *GetRayLogResponse) SetStatusCode(v int32) *GetRayLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRayLogResponse) SetBody(v *GetRayLogResponseBody) *GetRayLogResponse {
	s.Body = v
	return s
}

func (s *GetRayLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
