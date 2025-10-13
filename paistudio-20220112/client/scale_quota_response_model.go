// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleQuotaResponse
	GetStatusCode() *int32
	SetBody(v *ScaleQuotaResponseBody) *ScaleQuotaResponse
	GetBody() *ScaleQuotaResponseBody
}

type ScaleQuotaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleQuotaResponse) GoString() string {
	return s.String()
}

func (s *ScaleQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleQuotaResponse) GetBody() *ScaleQuotaResponseBody {
	return s.Body
}

func (s *ScaleQuotaResponse) SetHeaders(v map[string]*string) *ScaleQuotaResponse {
	s.Headers = v
	return s
}

func (s *ScaleQuotaResponse) SetStatusCode(v int32) *ScaleQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleQuotaResponse) SetBody(v *ScaleQuotaResponseBody) *ScaleQuotaResponse {
	s.Body = v
	return s
}

func (s *ScaleQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
