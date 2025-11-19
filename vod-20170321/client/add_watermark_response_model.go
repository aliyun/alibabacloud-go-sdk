// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *AddWatermarkResponseBody) *AddWatermarkResponse
	GetBody() *AddWatermarkResponseBody
}

type AddWatermarkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkResponse) GoString() string {
	return s.String()
}

func (s *AddWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWatermarkResponse) GetBody() *AddWatermarkResponseBody {
	return s.Body
}

func (s *AddWatermarkResponse) SetHeaders(v map[string]*string) *AddWatermarkResponse {
	s.Headers = v
	return s
}

func (s *AddWatermarkResponse) SetStatusCode(v int32) *AddWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWatermarkResponse) SetBody(v *AddWatermarkResponseBody) *AddWatermarkResponse {
	s.Body = v
	return s
}

func (s *AddWatermarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
