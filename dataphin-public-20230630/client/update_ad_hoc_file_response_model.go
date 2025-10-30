// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdHocFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdHocFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdHocFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdHocFileResponseBody) *UpdateAdHocFileResponse
	GetBody() *UpdateAdHocFileResponseBody
}

type UpdateAdHocFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdHocFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdHocFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdHocFileResponse) GetBody() *UpdateAdHocFileResponseBody {
	return s.Body
}

func (s *UpdateAdHocFileResponse) SetHeaders(v map[string]*string) *UpdateAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdHocFileResponse) SetStatusCode(v int32) *UpdateAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdHocFileResponse) SetBody(v *UpdateAdHocFileResponseBody) *UpdateAdHocFileResponse {
	s.Body = v
	return s
}

func (s *UpdateAdHocFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
