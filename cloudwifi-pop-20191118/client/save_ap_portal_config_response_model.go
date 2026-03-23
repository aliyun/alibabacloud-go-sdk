// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApPortalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApPortalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApPortalConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveApPortalConfigResponseBody) *SaveApPortalConfigResponse
	GetBody() *SaveApPortalConfigResponseBody
}

type SaveApPortalConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApPortalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApPortalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApPortalConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveApPortalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApPortalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApPortalConfigResponse) GetBody() *SaveApPortalConfigResponseBody {
	return s.Body
}

func (s *SaveApPortalConfigResponse) SetHeaders(v map[string]*string) *SaveApPortalConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveApPortalConfigResponse) SetStatusCode(v int32) *SaveApPortalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApPortalConfigResponse) SetBody(v *SaveApPortalConfigResponseBody) *SaveApPortalConfigResponse {
	s.Body = v
	return s
}

func (s *SaveApPortalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
