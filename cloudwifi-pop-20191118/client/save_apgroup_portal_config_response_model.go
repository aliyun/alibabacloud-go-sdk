// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupPortalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApgroupPortalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApgroupPortalConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveApgroupPortalConfigResponseBody) *SaveApgroupPortalConfigResponse
	GetBody() *SaveApgroupPortalConfigResponseBody
}

type SaveApgroupPortalConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApgroupPortalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApgroupPortalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupPortalConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveApgroupPortalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApgroupPortalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApgroupPortalConfigResponse) GetBody() *SaveApgroupPortalConfigResponseBody {
	return s.Body
}

func (s *SaveApgroupPortalConfigResponse) SetHeaders(v map[string]*string) *SaveApgroupPortalConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveApgroupPortalConfigResponse) SetStatusCode(v int32) *SaveApgroupPortalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApgroupPortalConfigResponse) SetBody(v *SaveApgroupPortalConfigResponseBody) *SaveApgroupPortalConfigResponse {
	s.Body = v
	return s
}

func (s *SaveApgroupPortalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
