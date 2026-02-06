// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobDataParsingTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobDataParsingTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobDataParsingTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobDataParsingTaskResponseBody) *CreateJobDataParsingTaskResponse
	GetBody() *CreateJobDataParsingTaskResponseBody
}

type CreateJobDataParsingTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobDataParsingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobDataParsingTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobDataParsingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateJobDataParsingTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobDataParsingTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobDataParsingTaskResponse) GetBody() *CreateJobDataParsingTaskResponseBody {
	return s.Body
}

func (s *CreateJobDataParsingTaskResponse) SetHeaders(v map[string]*string) *CreateJobDataParsingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateJobDataParsingTaskResponse) SetStatusCode(v int32) *CreateJobDataParsingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobDataParsingTaskResponse) SetBody(v *CreateJobDataParsingTaskResponseBody) *CreateJobDataParsingTaskResponse {
	s.Body = v
	return s
}

func (s *CreateJobDataParsingTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
