// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDetectionTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageDetectionTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageDetectionTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetImageDetectionTaskResultResponseBody) *GetImageDetectionTaskResultResponse
	GetBody() *GetImageDetectionTaskResultResponseBody
}

type GetImageDetectionTaskResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageDetectionTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageDetectionTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageDetectionTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageDetectionTaskResultResponse) GetBody() *GetImageDetectionTaskResultResponseBody {
	return s.Body
}

func (s *GetImageDetectionTaskResultResponse) SetHeaders(v map[string]*string) *GetImageDetectionTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageDetectionTaskResultResponse) SetStatusCode(v int32) *GetImageDetectionTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageDetectionTaskResultResponse) SetBody(v *GetImageDetectionTaskResultResponseBody) *GetImageDetectionTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetImageDetectionTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
