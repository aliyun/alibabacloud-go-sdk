// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppOtaTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppOtaTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppOtaTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppOtaTaskResponseBody) *CreateAppOtaTaskResponse
	GetBody() *CreateAppOtaTaskResponseBody
}

type CreateAppOtaTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppOtaTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppOtaTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppOtaTaskResponse) GetBody() *CreateAppOtaTaskResponseBody {
	return s.Body
}

func (s *CreateAppOtaTaskResponse) SetHeaders(v map[string]*string) *CreateAppOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAppOtaTaskResponse) SetStatusCode(v int32) *CreateAppOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppOtaTaskResponse) SetBody(v *CreateAppOtaTaskResponseBody) *CreateAppOtaTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAppOtaTaskResponse) Validate() error {
	return dara.Validate(s)
}
