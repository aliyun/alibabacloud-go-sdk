// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDIJobLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDIJobLogResponse
	GetStatusCode() *int32
	SetBody(v *GetDIJobLogResponseBody) *GetDIJobLogResponse
	GetBody() *GetDIJobLogResponseBody
}

type GetDIJobLogResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIJobLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetDIJobLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDIJobLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDIJobLogResponse) GetBody() *GetDIJobLogResponseBody {
	return s.Body
}

func (s *GetDIJobLogResponse) SetHeaders(v map[string]*string) *GetDIJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetDIJobLogResponse) SetStatusCode(v int32) *GetDIJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIJobLogResponse) SetBody(v *GetDIJobLogResponseBody) *GetDIJobLogResponse {
	s.Body = v
	return s
}

func (s *GetDIJobLogResponse) Validate() error {
	return dara.Validate(s)
}
