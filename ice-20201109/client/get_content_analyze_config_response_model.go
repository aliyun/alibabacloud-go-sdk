// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentAnalyzeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContentAnalyzeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContentAnalyzeConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetContentAnalyzeConfigResponseBody) *GetContentAnalyzeConfigResponse
	GetBody() *GetContentAnalyzeConfigResponseBody
}

type GetContentAnalyzeConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContentAnalyzeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContentAnalyzeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContentAnalyzeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetContentAnalyzeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContentAnalyzeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContentAnalyzeConfigResponse) GetBody() *GetContentAnalyzeConfigResponseBody {
	return s.Body
}

func (s *GetContentAnalyzeConfigResponse) SetHeaders(v map[string]*string) *GetContentAnalyzeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetContentAnalyzeConfigResponse) SetStatusCode(v int32) *GetContentAnalyzeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContentAnalyzeConfigResponse) SetBody(v *GetContentAnalyzeConfigResponseBody) *GetContentAnalyzeConfigResponse {
	s.Body = v
	return s
}

func (s *GetContentAnalyzeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
