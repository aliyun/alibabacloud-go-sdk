// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewTaskParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewTaskParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewTaskParametersResponse
	GetStatusCode() *int32
	SetBody(v *PreviewTaskParametersResponseBody) *PreviewTaskParametersResponse
	GetBody() *PreviewTaskParametersResponseBody
}

type PreviewTaskParametersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewTaskParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewTaskParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewTaskParametersResponse) GoString() string {
	return s.String()
}

func (s *PreviewTaskParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewTaskParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewTaskParametersResponse) GetBody() *PreviewTaskParametersResponseBody {
	return s.Body
}

func (s *PreviewTaskParametersResponse) SetHeaders(v map[string]*string) *PreviewTaskParametersResponse {
	s.Headers = v
	return s
}

func (s *PreviewTaskParametersResponse) SetStatusCode(v int32) *PreviewTaskParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewTaskParametersResponse) SetBody(v *PreviewTaskParametersResponseBody) *PreviewTaskParametersResponse {
	s.Body = v
	return s
}

func (s *PreviewTaskParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
