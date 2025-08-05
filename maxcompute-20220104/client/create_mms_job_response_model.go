// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMmsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMmsJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateMmsJobResponseBody) *CreateMmsJobResponse
	GetBody() *CreateMmsJobResponseBody
}

type CreateMmsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMmsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMmsJobResponse) GetBody() *CreateMmsJobResponseBody {
	return s.Body
}

func (s *CreateMmsJobResponse) SetHeaders(v map[string]*string) *CreateMmsJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsJobResponse) SetStatusCode(v int32) *CreateMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsJobResponse) SetBody(v *CreateMmsJobResponseBody) *CreateMmsJobResponse {
	s.Body = v
	return s
}

func (s *CreateMmsJobResponse) Validate() error {
	return dara.Validate(s)
}
