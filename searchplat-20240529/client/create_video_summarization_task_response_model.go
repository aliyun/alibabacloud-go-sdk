// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSummarizationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoSummarizationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoSummarizationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoSummarizationTaskResponseBody) *CreateVideoSummarizationTaskResponse
	GetBody() *CreateVideoSummarizationTaskResponseBody
}

type CreateVideoSummarizationTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoSummarizationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoSummarizationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoSummarizationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoSummarizationTaskResponse) GetBody() *CreateVideoSummarizationTaskResponseBody {
	return s.Body
}

func (s *CreateVideoSummarizationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoSummarizationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoSummarizationTaskResponse) SetStatusCode(v int32) *CreateVideoSummarizationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoSummarizationTaskResponse) SetBody(v *CreateVideoSummarizationTaskResponseBody) *CreateVideoSummarizationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoSummarizationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
