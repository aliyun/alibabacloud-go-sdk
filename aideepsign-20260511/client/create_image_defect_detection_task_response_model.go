// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDefectDetectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageDefectDetectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageDefectDetectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageDefectDetectionTaskResponseBody) *CreateImageDefectDetectionTaskResponse
	GetBody() *CreateImageDefectDetectionTaskResponseBody
}

type CreateImageDefectDetectionTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageDefectDetectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageDefectDetectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDefectDetectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageDefectDetectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageDefectDetectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageDefectDetectionTaskResponse) GetBody() *CreateImageDefectDetectionTaskResponseBody {
	return s.Body
}

func (s *CreateImageDefectDetectionTaskResponse) SetHeaders(v map[string]*string) *CreateImageDefectDetectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageDefectDetectionTaskResponse) SetStatusCode(v int32) *CreateImageDefectDetectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponse) SetBody(v *CreateImageDefectDetectionTaskResponseBody) *CreateImageDefectDetectionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageDefectDetectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
