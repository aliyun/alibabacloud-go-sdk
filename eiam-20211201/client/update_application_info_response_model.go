// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationInfoResponseBody) *UpdateApplicationInfoResponse
	GetBody() *UpdateApplicationInfoResponseBody
}

type UpdateApplicationInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationInfoResponse) GetBody() *UpdateApplicationInfoResponseBody {
	return s.Body
}

func (s *UpdateApplicationInfoResponse) SetHeaders(v map[string]*string) *UpdateApplicationInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationInfoResponse) SetStatusCode(v int32) *UpdateApplicationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationInfoResponse) SetBody(v *UpdateApplicationInfoResponseBody) *UpdateApplicationInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
