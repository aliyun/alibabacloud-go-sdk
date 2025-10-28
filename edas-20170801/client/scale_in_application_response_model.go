// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleInApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleInApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ScaleInApplicationResponseBody) *ScaleInApplicationResponse
	GetBody() *ScaleInApplicationResponseBody
}

type ScaleInApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleInApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleInApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleInApplicationResponse) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleInApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleInApplicationResponse) GetBody() *ScaleInApplicationResponseBody {
	return s.Body
}

func (s *ScaleInApplicationResponse) SetHeaders(v map[string]*string) *ScaleInApplicationResponse {
	s.Headers = v
	return s
}

func (s *ScaleInApplicationResponse) SetStatusCode(v int32) *ScaleInApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleInApplicationResponse) SetBody(v *ScaleInApplicationResponseBody) *ScaleInApplicationResponse {
	s.Body = v
	return s
}

func (s *ScaleInApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
