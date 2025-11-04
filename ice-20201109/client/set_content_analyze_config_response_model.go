// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetContentAnalyzeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetContentAnalyzeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetContentAnalyzeConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetContentAnalyzeConfigResponseBody) *SetContentAnalyzeConfigResponse
	GetBody() *SetContentAnalyzeConfigResponseBody
}

type SetContentAnalyzeConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetContentAnalyzeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetContentAnalyzeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetContentAnalyzeConfigResponse) GoString() string {
	return s.String()
}

func (s *SetContentAnalyzeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetContentAnalyzeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetContentAnalyzeConfigResponse) GetBody() *SetContentAnalyzeConfigResponseBody {
	return s.Body
}

func (s *SetContentAnalyzeConfigResponse) SetHeaders(v map[string]*string) *SetContentAnalyzeConfigResponse {
	s.Headers = v
	return s
}

func (s *SetContentAnalyzeConfigResponse) SetStatusCode(v int32) *SetContentAnalyzeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetContentAnalyzeConfigResponse) SetBody(v *SetContentAnalyzeConfigResponseBody) *SetContentAnalyzeConfigResponse {
	s.Body = v
	return s
}

func (s *SetContentAnalyzeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
