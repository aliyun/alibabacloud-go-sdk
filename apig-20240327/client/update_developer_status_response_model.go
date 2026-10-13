// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeveloperStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeveloperStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeveloperStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeveloperStatusResponseBody) *UpdateDeveloperStatusResponse
	GetBody() *UpdateDeveloperStatusResponseBody
}

type UpdateDeveloperStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeveloperStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeveloperStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeveloperStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeveloperStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeveloperStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeveloperStatusResponse) GetBody() *UpdateDeveloperStatusResponseBody {
	return s.Body
}

func (s *UpdateDeveloperStatusResponse) SetHeaders(v map[string]*string) *UpdateDeveloperStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeveloperStatusResponse) SetStatusCode(v int32) *UpdateDeveloperStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeveloperStatusResponse) SetBody(v *UpdateDeveloperStatusResponseBody) *UpdateDeveloperStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateDeveloperStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
