// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaProducingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaProducingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaProducingJobResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaProducingJobResponseBody) *GetMediaProducingJobResponse
	GetBody() *GetMediaProducingJobResponseBody
}

type GetMediaProducingJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaProducingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaProducingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaProducingJobResponse) GetBody() *GetMediaProducingJobResponseBody {
	return s.Body
}

func (s *GetMediaProducingJobResponse) SetHeaders(v map[string]*string) *GetMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *GetMediaProducingJobResponse) SetStatusCode(v int32) *GetMediaProducingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaProducingJobResponse) SetBody(v *GetMediaProducingJobResponseBody) *GetMediaProducingJobResponse {
	s.Body = v
	return s
}

func (s *GetMediaProducingJobResponse) Validate() error {
	return dara.Validate(s)
}
