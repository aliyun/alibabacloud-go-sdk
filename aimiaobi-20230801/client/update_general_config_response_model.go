// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneralConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGeneralConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGeneralConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGeneralConfigResponseBody) *UpdateGeneralConfigResponse
	GetBody() *UpdateGeneralConfigResponseBody
}

type UpdateGeneralConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGeneralConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGeneralConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneralConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateGeneralConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGeneralConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGeneralConfigResponse) GetBody() *UpdateGeneralConfigResponseBody {
	return s.Body
}

func (s *UpdateGeneralConfigResponse) SetHeaders(v map[string]*string) *UpdateGeneralConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateGeneralConfigResponse) SetStatusCode(v int32) *UpdateGeneralConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGeneralConfigResponse) SetBody(v *UpdateGeneralConfigResponseBody) *UpdateGeneralConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateGeneralConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
