// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTraceAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTraceAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTraceAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveTraceAppConfigResponseBody) *SaveTraceAppConfigResponse
	GetBody() *SaveTraceAppConfigResponseBody
}

type SaveTraceAppConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTraceAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTraceAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTraceAppConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTraceAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTraceAppConfigResponse) GetBody() *SaveTraceAppConfigResponseBody {
	return s.Body
}

func (s *SaveTraceAppConfigResponse) SetHeaders(v map[string]*string) *SaveTraceAppConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveTraceAppConfigResponse) SetStatusCode(v int32) *SaveTraceAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTraceAppConfigResponse) SetBody(v *SaveTraceAppConfigResponseBody) *SaveTraceAppConfigResponse {
	s.Body = v
	return s
}

func (s *SaveTraceAppConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
