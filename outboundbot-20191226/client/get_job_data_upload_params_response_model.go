// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDataUploadParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobDataUploadParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobDataUploadParamsResponse
	GetStatusCode() *int32
	SetBody(v *GetJobDataUploadParamsResponseBody) *GetJobDataUploadParamsResponse
	GetBody() *GetJobDataUploadParamsResponseBody
}

type GetJobDataUploadParamsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobDataUploadParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobDataUploadParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobDataUploadParamsResponse) GoString() string {
	return s.String()
}

func (s *GetJobDataUploadParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobDataUploadParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobDataUploadParamsResponse) GetBody() *GetJobDataUploadParamsResponseBody {
	return s.Body
}

func (s *GetJobDataUploadParamsResponse) SetHeaders(v map[string]*string) *GetJobDataUploadParamsResponse {
	s.Headers = v
	return s
}

func (s *GetJobDataUploadParamsResponse) SetStatusCode(v int32) *GetJobDataUploadParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobDataUploadParamsResponse) SetBody(v *GetJobDataUploadParamsResponseBody) *GetJobDataUploadParamsResponse {
	s.Body = v
	return s
}

func (s *GetJobDataUploadParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
