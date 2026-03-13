// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSegmentationTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoSegmentationTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoSegmentationTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoSegmentationTaskStatusResponseBody) *GetVideoSegmentationTaskStatusResponse
	GetBody() *GetVideoSegmentationTaskStatusResponseBody
}

type GetVideoSegmentationTaskStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoSegmentationTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoSegmentationTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoSegmentationTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoSegmentationTaskStatusResponse) GetBody() *GetVideoSegmentationTaskStatusResponseBody {
	return s.Body
}

func (s *GetVideoSegmentationTaskStatusResponse) SetHeaders(v map[string]*string) *GetVideoSegmentationTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponse) SetStatusCode(v int32) *GetVideoSegmentationTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponse) SetBody(v *GetVideoSegmentationTaskStatusResponseBody) *GetVideoSegmentationTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
