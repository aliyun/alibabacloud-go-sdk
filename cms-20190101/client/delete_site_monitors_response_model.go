// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSiteMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSiteMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSiteMonitorsResponseBody) *DeleteSiteMonitorsResponse
	GetBody() *DeleteSiteMonitorsResponseBody
}

type DeleteSiteMonitorsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSiteMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSiteMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteMonitorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSiteMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSiteMonitorsResponse) GetBody() *DeleteSiteMonitorsResponseBody {
	return s.Body
}

func (s *DeleteSiteMonitorsResponse) SetHeaders(v map[string]*string) *DeleteSiteMonitorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteMonitorsResponse) SetStatusCode(v int32) *DeleteSiteMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSiteMonitorsResponse) SetBody(v *DeleteSiteMonitorsResponseBody) *DeleteSiteMonitorsResponse {
	s.Body = v
	return s
}

func (s *DeleteSiteMonitorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
