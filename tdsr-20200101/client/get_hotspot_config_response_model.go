// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotConfigResponseBody) *GetHotspotConfigResponse
	GetBody() *GetHotspotConfigResponseBody
}

type GetHotspotConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotConfigResponse) GetBody() *GetHotspotConfigResponseBody {
	return s.Body
}

func (s *GetHotspotConfigResponse) SetHeaders(v map[string]*string) *GetHotspotConfigResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotConfigResponse) SetStatusCode(v int32) *GetHotspotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotConfigResponse) SetBody(v *GetHotspotConfigResponseBody) *GetHotspotConfigResponse {
	s.Body = v
	return s
}

func (s *GetHotspotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
