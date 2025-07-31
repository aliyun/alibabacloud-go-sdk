// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdHocFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAdHocFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAdHocFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateAdHocFileResponseBody) *CreateAdHocFileResponse
	GetBody() *CreateAdHocFileResponseBody
}

type CreateAdHocFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdHocFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAdHocFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAdHocFileResponse) GetBody() *CreateAdHocFileResponseBody {
	return s.Body
}

func (s *CreateAdHocFileResponse) SetHeaders(v map[string]*string) *CreateAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *CreateAdHocFileResponse) SetStatusCode(v int32) *CreateAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdHocFileResponse) SetBody(v *CreateAdHocFileResponseBody) *CreateAdHocFileResponse {
	s.Body = v
	return s
}

func (s *CreateAdHocFileResponse) Validate() error {
	return dara.Validate(s)
}
