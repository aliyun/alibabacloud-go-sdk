// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSegmentationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoSegmentationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoSegmentationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoSegmentationTaskResponseBody) *CreateVideoSegmentationTaskResponse
	GetBody() *CreateVideoSegmentationTaskResponseBody
}

type CreateVideoSegmentationTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoSegmentationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoSegmentationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSegmentationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoSegmentationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoSegmentationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoSegmentationTaskResponse) GetBody() *CreateVideoSegmentationTaskResponseBody {
	return s.Body
}

func (s *CreateVideoSegmentationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoSegmentationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoSegmentationTaskResponse) SetStatusCode(v int32) *CreateVideoSegmentationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoSegmentationTaskResponse) SetBody(v *CreateVideoSegmentationTaskResponseBody) *CreateVideoSegmentationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoSegmentationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
