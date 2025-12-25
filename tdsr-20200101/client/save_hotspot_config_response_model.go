// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveHotspotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveHotspotConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveHotspotConfigResponseBody) *SaveHotspotConfigResponse
	GetBody() *SaveHotspotConfigResponseBody
}

type SaveHotspotConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveHotspotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveHotspotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveHotspotConfigResponse) GetBody() *SaveHotspotConfigResponseBody {
	return s.Body
}

func (s *SaveHotspotConfigResponse) SetHeaders(v map[string]*string) *SaveHotspotConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveHotspotConfigResponse) SetStatusCode(v int32) *SaveHotspotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveHotspotConfigResponse) SetBody(v *SaveHotspotConfigResponseBody) *SaveHotspotConfigResponse {
	s.Body = v
	return s
}

func (s *SaveHotspotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
