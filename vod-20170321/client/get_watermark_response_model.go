// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *GetWatermarkResponseBody) *GetWatermarkResponse
	GetBody() *GetWatermarkResponseBody
}

type GetWatermarkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWatermarkResponse) GetBody() *GetWatermarkResponseBody {
	return s.Body
}

func (s *GetWatermarkResponse) SetHeaders(v map[string]*string) *GetWatermarkResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarkResponse) SetStatusCode(v int32) *GetWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWatermarkResponse) SetBody(v *GetWatermarkResponseBody) *GetWatermarkResponse {
	s.Body = v
	return s
}

func (s *GetWatermarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
