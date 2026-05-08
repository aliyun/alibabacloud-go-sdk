// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageToVideoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitImageToVideoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitImageToVideoTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitImageToVideoTaskResponseBody) *SubmitImageToVideoTaskResponse
	GetBody() *SubmitImageToVideoTaskResponseBody
}

type SubmitImageToVideoTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImageToVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImageToVideoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageToVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitImageToVideoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitImageToVideoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitImageToVideoTaskResponse) GetBody() *SubmitImageToVideoTaskResponseBody {
	return s.Body
}

func (s *SubmitImageToVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitImageToVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitImageToVideoTaskResponse) SetStatusCode(v int32) *SubmitImageToVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageToVideoTaskResponse) SetBody(v *SubmitImageToVideoTaskResponseBody) *SubmitImageToVideoTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitImageToVideoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
