// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTrailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceTrailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceTrailResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceTrailResponseBody) *CreateServiceTrailResponse
	GetBody() *CreateServiceTrailResponseBody
}

type CreateServiceTrailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceTrailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceTrailResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTrailResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceTrailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceTrailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceTrailResponse) GetBody() *CreateServiceTrailResponseBody {
	return s.Body
}

func (s *CreateServiceTrailResponse) SetHeaders(v map[string]*string) *CreateServiceTrailResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceTrailResponse) SetStatusCode(v int32) *CreateServiceTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceTrailResponse) SetBody(v *CreateServiceTrailResponseBody) *CreateServiceTrailResponse {
	s.Body = v
	return s
}

func (s *CreateServiceTrailResponse) Validate() error {
	return dara.Validate(s)
}
