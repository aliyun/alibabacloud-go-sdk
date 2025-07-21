// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppOtaVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppOtaVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppOtaVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppOtaVersionResponseBody) *CreateAppOtaVersionResponse
	GetBody() *CreateAppOtaVersionResponseBody
}

type CreateAppOtaVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppOtaVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppOtaVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppOtaVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppOtaVersionResponse) GetBody() *CreateAppOtaVersionResponseBody {
	return s.Body
}

func (s *CreateAppOtaVersionResponse) SetHeaders(v map[string]*string) *CreateAppOtaVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppOtaVersionResponse) SetStatusCode(v int32) *CreateAppOtaVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppOtaVersionResponse) SetBody(v *CreateAppOtaVersionResponseBody) *CreateAppOtaVersionResponse {
	s.Body = v
	return s
}

func (s *CreateAppOtaVersionResponse) Validate() error {
	return dara.Validate(s)
}
