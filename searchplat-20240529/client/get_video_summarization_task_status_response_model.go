// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSummarizationTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoSummarizationTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoSummarizationTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoSummarizationTaskStatusResponseBody) *GetVideoSummarizationTaskStatusResponse
	GetBody() *GetVideoSummarizationTaskStatusResponseBody
}

type GetVideoSummarizationTaskStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoSummarizationTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoSummarizationTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoSummarizationTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoSummarizationTaskStatusResponse) GetBody() *GetVideoSummarizationTaskStatusResponseBody {
	return s.Body
}

func (s *GetVideoSummarizationTaskStatusResponse) SetHeaders(v map[string]*string) *GetVideoSummarizationTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponse) SetStatusCode(v int32) *GetVideoSummarizationTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponse) SetBody(v *GetVideoSummarizationTaskStatusResponseBody) *GetVideoSummarizationTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
