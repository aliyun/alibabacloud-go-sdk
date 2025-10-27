// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetImageSensitiveFileStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetImageSensitiveFileStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetImageSensitiveFileStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetImageSensitiveFileStatusResponseBody) *SetImageSensitiveFileStatusResponse
	GetBody() *SetImageSensitiveFileStatusResponseBody
}

type SetImageSensitiveFileStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetImageSensitiveFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetImageSensitiveFileStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetImageSensitiveFileStatusResponse) GoString() string {
	return s.String()
}

func (s *SetImageSensitiveFileStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetImageSensitiveFileStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetImageSensitiveFileStatusResponse) GetBody() *SetImageSensitiveFileStatusResponseBody {
	return s.Body
}

func (s *SetImageSensitiveFileStatusResponse) SetHeaders(v map[string]*string) *SetImageSensitiveFileStatusResponse {
	s.Headers = v
	return s
}

func (s *SetImageSensitiveFileStatusResponse) SetStatusCode(v int32) *SetImageSensitiveFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetImageSensitiveFileStatusResponse) SetBody(v *SetImageSensitiveFileStatusResponseBody) *SetImageSensitiveFileStatusResponse {
	s.Body = v
	return s
}

func (s *SetImageSensitiveFileStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
