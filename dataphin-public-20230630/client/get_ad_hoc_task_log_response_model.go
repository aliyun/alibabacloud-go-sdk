// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocTaskLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdHocTaskLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdHocTaskLogResponse
	GetStatusCode() *int32
	SetBody(v *GetAdHocTaskLogResponseBody) *GetAdHocTaskLogResponse
	GetBody() *GetAdHocTaskLogResponseBody
}

type GetAdHocTaskLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdHocTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdHocTaskLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskLogResponse) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdHocTaskLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdHocTaskLogResponse) GetBody() *GetAdHocTaskLogResponseBody {
	return s.Body
}

func (s *GetAdHocTaskLogResponse) SetHeaders(v map[string]*string) *GetAdHocTaskLogResponse {
	s.Headers = v
	return s
}

func (s *GetAdHocTaskLogResponse) SetStatusCode(v int32) *GetAdHocTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdHocTaskLogResponse) SetBody(v *GetAdHocTaskLogResponseBody) *GetAdHocTaskLogResponse {
	s.Body = v
	return s
}

func (s *GetAdHocTaskLogResponse) Validate() error {
	return dara.Validate(s)
}
