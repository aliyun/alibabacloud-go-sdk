// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomizationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomizationConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomizationConfigResponseBody) *DeleteCustomizationConfigResponse
	GetBody() *DeleteCustomizationConfigResponseBody
}

type DeleteCustomizationConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizationConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomizationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomizationConfigResponse) GetBody() *DeleteCustomizationConfigResponseBody {
	return s.Body
}

func (s *DeleteCustomizationConfigResponse) SetHeaders(v map[string]*string) *DeleteCustomizationConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizationConfigResponse) SetStatusCode(v int32) *DeleteCustomizationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizationConfigResponse) SetBody(v *DeleteCustomizationConfigResponseBody) *DeleteCustomizationConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomizationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
