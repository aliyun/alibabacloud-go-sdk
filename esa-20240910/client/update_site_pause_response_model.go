// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSitePauseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSitePauseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSitePauseResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSitePauseResponseBody) *UpdateSitePauseResponse
	GetBody() *UpdateSitePauseResponseBody
}

type UpdateSitePauseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSitePauseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSitePauseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSitePauseResponse) GoString() string {
	return s.String()
}

func (s *UpdateSitePauseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSitePauseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSitePauseResponse) GetBody() *UpdateSitePauseResponseBody {
	return s.Body
}

func (s *UpdateSitePauseResponse) SetHeaders(v map[string]*string) *UpdateSitePauseResponse {
	s.Headers = v
	return s
}

func (s *UpdateSitePauseResponse) SetStatusCode(v int32) *UpdateSitePauseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSitePauseResponse) SetBody(v *UpdateSitePauseResponseBody) *UpdateSitePauseResponse {
	s.Body = v
	return s
}

func (s *UpdateSitePauseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
