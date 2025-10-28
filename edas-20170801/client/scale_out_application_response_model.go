// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleOutApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleOutApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ScaleOutApplicationResponseBody) *ScaleOutApplicationResponse
	GetBody() *ScaleOutApplicationResponseBody
}

type ScaleOutApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleOutApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleOutApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutApplicationResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleOutApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleOutApplicationResponse) GetBody() *ScaleOutApplicationResponseBody {
	return s.Body
}

func (s *ScaleOutApplicationResponse) SetHeaders(v map[string]*string) *ScaleOutApplicationResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutApplicationResponse) SetStatusCode(v int32) *ScaleOutApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleOutApplicationResponse) SetBody(v *ScaleOutApplicationResponseBody) *ScaleOutApplicationResponse {
	s.Body = v
	return s
}

func (s *ScaleOutApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
