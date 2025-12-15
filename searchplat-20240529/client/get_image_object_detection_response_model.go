// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageObjectDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageObjectDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageObjectDetectionResponse
	GetStatusCode() *int32
	SetBody(v *GetImageObjectDetectionResponseBody) *GetImageObjectDetectionResponse
	GetBody() *GetImageObjectDetectionResponseBody
}

type GetImageObjectDetectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageObjectDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageObjectDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionResponse) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageObjectDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageObjectDetectionResponse) GetBody() *GetImageObjectDetectionResponseBody {
	return s.Body
}

func (s *GetImageObjectDetectionResponse) SetHeaders(v map[string]*string) *GetImageObjectDetectionResponse {
	s.Headers = v
	return s
}

func (s *GetImageObjectDetectionResponse) SetStatusCode(v int32) *GetImageObjectDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageObjectDetectionResponse) SetBody(v *GetImageObjectDetectionResponseBody) *GetImageObjectDetectionResponse {
	s.Body = v
	return s
}

func (s *GetImageObjectDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
