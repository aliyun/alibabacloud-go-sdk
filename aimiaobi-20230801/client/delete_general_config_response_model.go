// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGeneralConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGeneralConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGeneralConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGeneralConfigResponseBody) *DeleteGeneralConfigResponse
	GetBody() *DeleteGeneralConfigResponseBody
}

type DeleteGeneralConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGeneralConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGeneralConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGeneralConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteGeneralConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGeneralConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGeneralConfigResponse) GetBody() *DeleteGeneralConfigResponseBody {
	return s.Body
}

func (s *DeleteGeneralConfigResponse) SetHeaders(v map[string]*string) *DeleteGeneralConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteGeneralConfigResponse) SetStatusCode(v int32) *DeleteGeneralConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGeneralConfigResponse) SetBody(v *DeleteGeneralConfigResponseBody) *DeleteGeneralConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteGeneralConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
