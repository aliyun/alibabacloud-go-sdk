// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDefectDetectionTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageDefectDetectionTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageDefectDetectionTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetImageDefectDetectionTaskResultResponseBody) *GetImageDefectDetectionTaskResultResponse
	GetBody() *GetImageDefectDetectionTaskResultResponseBody
}

type GetImageDefectDetectionTaskResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageDefectDetectionTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageDefectDetectionTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageDefectDetectionTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageDefectDetectionTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageDefectDetectionTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageDefectDetectionTaskResultResponse) GetBody() *GetImageDefectDetectionTaskResultResponseBody {
	return s.Body
}

func (s *GetImageDefectDetectionTaskResultResponse) SetHeaders(v map[string]*string) *GetImageDefectDetectionTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponse) SetStatusCode(v int32) *GetImageDefectDetectionTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponse) SetBody(v *GetImageDefectDetectionTaskResultResponseBody) *GetImageDefectDetectionTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
