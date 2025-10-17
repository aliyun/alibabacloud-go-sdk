// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEssayCorrectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitEssayCorrectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitEssayCorrectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitEssayCorrectionTaskResponseBody) *SubmitEssayCorrectionTaskResponse
	GetBody() *SubmitEssayCorrectionTaskResponseBody
}

type SubmitEssayCorrectionTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitEssayCorrectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitEssayCorrectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitEssayCorrectionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitEssayCorrectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitEssayCorrectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitEssayCorrectionTaskResponse) GetBody() *SubmitEssayCorrectionTaskResponseBody {
	return s.Body
}

func (s *SubmitEssayCorrectionTaskResponse) SetHeaders(v map[string]*string) *SubmitEssayCorrectionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitEssayCorrectionTaskResponse) SetStatusCode(v int32) *SubmitEssayCorrectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponse) SetBody(v *SubmitEssayCorrectionTaskResponseBody) *SubmitEssayCorrectionTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitEssayCorrectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
