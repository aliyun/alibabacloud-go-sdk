// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDIJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateDIJobResponseBody) *CreateDIJobResponse
	GetBody() *CreateDIJobResponseBody
}

type CreateDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDIJobResponse) GetBody() *CreateDIJobResponseBody {
	return s.Body
}

func (s *CreateDIJobResponse) SetHeaders(v map[string]*string) *CreateDIJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDIJobResponse) SetStatusCode(v int32) *CreateDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDIJobResponse) SetBody(v *CreateDIJobResponseBody) *CreateDIJobResponse {
	s.Body = v
	return s
}

func (s *CreateDIJobResponse) Validate() error {
	return dara.Validate(s)
}
