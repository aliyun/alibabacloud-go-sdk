// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoRenderJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoRenderJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoRenderJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoRenderJobResponseBody) *SubmitVideoRenderJobResponse
	GetBody() *SubmitVideoRenderJobResponseBody
}

type SubmitVideoRenderJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoRenderJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoRenderJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoRenderJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoRenderJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoRenderJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoRenderJobResponse) GetBody() *SubmitVideoRenderJobResponseBody {
	return s.Body
}

func (s *SubmitVideoRenderJobResponse) SetHeaders(v map[string]*string) *SubmitVideoRenderJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoRenderJobResponse) SetStatusCode(v int32) *SubmitVideoRenderJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoRenderJobResponse) SetBody(v *SubmitVideoRenderJobResponseBody) *SubmitVideoRenderJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoRenderJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
