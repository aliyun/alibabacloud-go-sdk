// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupBasicConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApgroupBasicConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApgroupBasicConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveApgroupBasicConfigResponseBody) *SaveApgroupBasicConfigResponse
	GetBody() *SaveApgroupBasicConfigResponseBody
}

type SaveApgroupBasicConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApgroupBasicConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApgroupBasicConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupBasicConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveApgroupBasicConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApgroupBasicConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApgroupBasicConfigResponse) GetBody() *SaveApgroupBasicConfigResponseBody {
	return s.Body
}

func (s *SaveApgroupBasicConfigResponse) SetHeaders(v map[string]*string) *SaveApgroupBasicConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveApgroupBasicConfigResponse) SetStatusCode(v int32) *SaveApgroupBasicConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApgroupBasicConfigResponse) SetBody(v *SaveApgroupBasicConfigResponseBody) *SaveApgroupBasicConfigResponse {
	s.Body = v
	return s
}

func (s *SaveApgroupBasicConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
