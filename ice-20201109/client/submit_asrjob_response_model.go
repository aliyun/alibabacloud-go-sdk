// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitASRJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitASRJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitASRJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitASRJobResponseBody) *SubmitASRJobResponse
	GetBody() *SubmitASRJobResponseBody
}

type SubmitASRJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitASRJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitASRJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitASRJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitASRJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitASRJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitASRJobResponse) GetBody() *SubmitASRJobResponseBody {
	return s.Body
}

func (s *SubmitASRJobResponse) SetHeaders(v map[string]*string) *SubmitASRJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitASRJobResponse) SetStatusCode(v int32) *SubmitASRJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitASRJobResponse) SetBody(v *SubmitASRJobResponseBody) *SubmitASRJobResponse {
	s.Body = v
	return s
}

func (s *SubmitASRJobResponse) Validate() error {
	return dara.Validate(s)
}
