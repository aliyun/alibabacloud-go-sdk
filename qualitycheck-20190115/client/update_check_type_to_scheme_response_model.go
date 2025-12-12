// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckTypeToSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCheckTypeToSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCheckTypeToSchemeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCheckTypeToSchemeResponseBody) *UpdateCheckTypeToSchemeResponse
	GetBody() *UpdateCheckTypeToSchemeResponseBody
}

type UpdateCheckTypeToSchemeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCheckTypeToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCheckTypeToSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckTypeToSchemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCheckTypeToSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCheckTypeToSchemeResponse) GetBody() *UpdateCheckTypeToSchemeResponseBody {
	return s.Body
}

func (s *UpdateCheckTypeToSchemeResponse) SetHeaders(v map[string]*string) *UpdateCheckTypeToSchemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckTypeToSchemeResponse) SetStatusCode(v int32) *UpdateCheckTypeToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponse) SetBody(v *UpdateCheckTypeToSchemeResponseBody) *UpdateCheckTypeToSchemeResponse {
	s.Body = v
	return s
}

func (s *UpdateCheckTypeToSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
