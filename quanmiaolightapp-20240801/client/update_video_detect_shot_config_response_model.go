// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoDetectShotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoDetectShotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoDetectShotConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoDetectShotConfigResponseBody) *UpdateVideoDetectShotConfigResponse
	GetBody() *UpdateVideoDetectShotConfigResponseBody
}

type UpdateVideoDetectShotConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoDetectShotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoDetectShotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoDetectShotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoDetectShotConfigResponse) GetBody() *UpdateVideoDetectShotConfigResponseBody {
	return s.Body
}

func (s *UpdateVideoDetectShotConfigResponse) SetHeaders(v map[string]*string) *UpdateVideoDetectShotConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoDetectShotConfigResponse) SetStatusCode(v int32) *UpdateVideoDetectShotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoDetectShotConfigResponse) SetBody(v *UpdateVideoDetectShotConfigResponseBody) *UpdateVideoDetectShotConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoDetectShotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
