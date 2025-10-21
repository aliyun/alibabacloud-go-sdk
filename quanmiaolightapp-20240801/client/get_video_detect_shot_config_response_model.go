// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetectShotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoDetectShotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoDetectShotConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoDetectShotConfigResponseBody) *GetVideoDetectShotConfigResponse
	GetBody() *GetVideoDetectShotConfigResponseBody
}

type GetVideoDetectShotConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoDetectShotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoDetectShotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoDetectShotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoDetectShotConfigResponse) GetBody() *GetVideoDetectShotConfigResponseBody {
	return s.Body
}

func (s *GetVideoDetectShotConfigResponse) SetHeaders(v map[string]*string) *GetVideoDetectShotConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVideoDetectShotConfigResponse) SetStatusCode(v int32) *GetVideoDetectShotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoDetectShotConfigResponse) SetBody(v *GetVideoDetectShotConfigResponseBody) *GetVideoDetectShotConfigResponse {
	s.Body = v
	return s
}

func (s *GetVideoDetectShotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
