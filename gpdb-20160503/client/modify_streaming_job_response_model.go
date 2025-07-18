// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStreamingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStreamingJobResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStreamingJobResponseBody) *ModifyStreamingJobResponse
	GetBody() *ModifyStreamingJobResponseBody
}

type ModifyStreamingJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStreamingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStreamingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyStreamingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStreamingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStreamingJobResponse) GetBody() *ModifyStreamingJobResponseBody {
	return s.Body
}

func (s *ModifyStreamingJobResponse) SetHeaders(v map[string]*string) *ModifyStreamingJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyStreamingJobResponse) SetStatusCode(v int32) *ModifyStreamingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStreamingJobResponse) SetBody(v *ModifyStreamingJobResponseBody) *ModifyStreamingJobResponse {
	s.Body = v
	return s
}

func (s *ModifyStreamingJobResponse) Validate() error {
	return dara.Validate(s)
}
