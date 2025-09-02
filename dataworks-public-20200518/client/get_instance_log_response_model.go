// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceLogResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceLogResponseBody) *GetInstanceLogResponse
	GetBody() *GetInstanceLogResponseBody
}

type GetInstanceLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceLogResponse) GetBody() *GetInstanceLogResponseBody {
	return s.Body
}

func (s *GetInstanceLogResponse) SetHeaders(v map[string]*string) *GetInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceLogResponse) SetStatusCode(v int32) *GetInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceLogResponse) SetBody(v *GetInstanceLogResponseBody) *GetInstanceLogResponse {
	s.Body = v
	return s
}

func (s *GetInstanceLogResponse) Validate() error {
	return dara.Validate(s)
}
