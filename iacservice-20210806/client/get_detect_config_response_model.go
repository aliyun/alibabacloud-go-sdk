// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDetectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDetectConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetDetectConfigResponseBody) *GetDetectConfigResponse
	GetBody() *GetDetectConfigResponseBody
}

type GetDetectConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDetectConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDetectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDetectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDetectConfigResponse) GetBody() *GetDetectConfigResponseBody {
	return s.Body
}

func (s *GetDetectConfigResponse) SetHeaders(v map[string]*string) *GetDetectConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDetectConfigResponse) SetStatusCode(v int32) *GetDetectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectConfigResponse) SetBody(v *GetDetectConfigResponseBody) *GetDetectConfigResponse {
	s.Body = v
	return s
}

func (s *GetDetectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
