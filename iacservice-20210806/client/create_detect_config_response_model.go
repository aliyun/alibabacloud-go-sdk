// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDetectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDetectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDetectConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateDetectConfigResponseBody) *CreateDetectConfigResponse
	GetBody() *CreateDetectConfigResponseBody
}

type CreateDetectConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDetectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDetectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateDetectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDetectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDetectConfigResponse) GetBody() *CreateDetectConfigResponseBody {
	return s.Body
}

func (s *CreateDetectConfigResponse) SetHeaders(v map[string]*string) *CreateDetectConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateDetectConfigResponse) SetStatusCode(v int32) *CreateDetectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDetectConfigResponse) SetBody(v *CreateDetectConfigResponseBody) *CreateDetectConfigResponse {
	s.Body = v
	return s
}

func (s *CreateDetectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
