// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateQwenEduApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateQwenEduApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateQwenEduApiResponse
	GetStatusCode() *int32
	SetBody(v *ActivateQwenEduApiResponseBody) *ActivateQwenEduApiResponse
	GetBody() *ActivateQwenEduApiResponseBody
}

type ActivateQwenEduApiResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateQwenEduApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateQwenEduApiResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateQwenEduApiResponse) GoString() string {
	return s.String()
}

func (s *ActivateQwenEduApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateQwenEduApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateQwenEduApiResponse) GetBody() *ActivateQwenEduApiResponseBody {
	return s.Body
}

func (s *ActivateQwenEduApiResponse) SetHeaders(v map[string]*string) *ActivateQwenEduApiResponse {
	s.Headers = v
	return s
}

func (s *ActivateQwenEduApiResponse) SetStatusCode(v int32) *ActivateQwenEduApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateQwenEduApiResponse) SetBody(v *ActivateQwenEduApiResponseBody) *ActivateQwenEduApiResponse {
	s.Body = v
	return s
}

func (s *ActivateQwenEduApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
