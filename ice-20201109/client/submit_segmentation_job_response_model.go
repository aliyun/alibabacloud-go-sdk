// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSegmentationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSegmentationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSegmentationJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSegmentationJobResponseBody) *SubmitSegmentationJobResponse
	GetBody() *SubmitSegmentationJobResponseBody
}

type SubmitSegmentationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSegmentationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSegmentationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSegmentationJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSegmentationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSegmentationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSegmentationJobResponse) GetBody() *SubmitSegmentationJobResponseBody {
	return s.Body
}

func (s *SubmitSegmentationJobResponse) SetHeaders(v map[string]*string) *SubmitSegmentationJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSegmentationJobResponse) SetStatusCode(v int32) *SubmitSegmentationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSegmentationJobResponse) SetBody(v *SubmitSegmentationJobResponseBody) *SubmitSegmentationJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSegmentationJobResponse) Validate() error {
	return dara.Validate(s)
}
