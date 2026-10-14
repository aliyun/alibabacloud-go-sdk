// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeC2paBucketSignTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeC2paBucketSignTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeC2paBucketSignTaskResponse
	GetStatusCode() *int32
	SetBody(v *ResumeC2paBucketSignTaskResponseBody) *ResumeC2paBucketSignTaskResponse
	GetBody() *ResumeC2paBucketSignTaskResponseBody
}

type ResumeC2paBucketSignTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeC2paBucketSignTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeC2paBucketSignTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeC2paBucketSignTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeC2paBucketSignTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeC2paBucketSignTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeC2paBucketSignTaskResponse) GetBody() *ResumeC2paBucketSignTaskResponseBody {
	return s.Body
}

func (s *ResumeC2paBucketSignTaskResponse) SetHeaders(v map[string]*string) *ResumeC2paBucketSignTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeC2paBucketSignTaskResponse) SetStatusCode(v int32) *ResumeC2paBucketSignTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponse) SetBody(v *ResumeC2paBucketSignTaskResponseBody) *ResumeC2paBucketSignTaskResponse {
	s.Body = v
	return s
}

func (s *ResumeC2paBucketSignTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
