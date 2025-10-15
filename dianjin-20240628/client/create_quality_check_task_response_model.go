// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityCheckTaskResponseBody) *CreateQualityCheckTaskResponse
	GetBody() *CreateQualityCheckTaskResponseBody
}

type CreateQualityCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityCheckTaskResponse) GetBody() *CreateQualityCheckTaskResponseBody {
	return s.Body
}

func (s *CreateQualityCheckTaskResponse) SetHeaders(v map[string]*string) *CreateQualityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityCheckTaskResponse) SetStatusCode(v int32) *CreateQualityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityCheckTaskResponse) SetBody(v *CreateQualityCheckTaskResponseBody) *CreateQualityCheckTaskResponse {
	s.Body = v
	return s
}

func (s *CreateQualityCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
