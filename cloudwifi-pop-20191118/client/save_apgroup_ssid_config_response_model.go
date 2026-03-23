// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupSsidConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApgroupSsidConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApgroupSsidConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveApgroupSsidConfigResponseBody) *SaveApgroupSsidConfigResponse
	GetBody() *SaveApgroupSsidConfigResponseBody
}

type SaveApgroupSsidConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApgroupSsidConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApgroupSsidConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupSsidConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveApgroupSsidConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApgroupSsidConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApgroupSsidConfigResponse) GetBody() *SaveApgroupSsidConfigResponseBody {
	return s.Body
}

func (s *SaveApgroupSsidConfigResponse) SetHeaders(v map[string]*string) *SaveApgroupSsidConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveApgroupSsidConfigResponse) SetStatusCode(v int32) *SaveApgroupSsidConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApgroupSsidConfigResponse) SetBody(v *SaveApgroupSsidConfigResponseBody) *SaveApgroupSsidConfigResponse {
	s.Body = v
	return s
}

func (s *SaveApgroupSsidConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
