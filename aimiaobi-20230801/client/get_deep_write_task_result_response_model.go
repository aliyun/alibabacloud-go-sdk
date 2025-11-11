// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeepWriteTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeepWriteTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeepWriteTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDeepWriteTaskResultResponseBody) *GetDeepWriteTaskResultResponse
	GetBody() *GetDeepWriteTaskResultResponseBody
}

type GetDeepWriteTaskResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeepWriteTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeepWriteTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeepWriteTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeepWriteTaskResultResponse) GetBody() *GetDeepWriteTaskResultResponseBody {
	return s.Body
}

func (s *GetDeepWriteTaskResultResponse) SetHeaders(v map[string]*string) *GetDeepWriteTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetDeepWriteTaskResultResponse) SetStatusCode(v int32) *GetDeepWriteTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeepWriteTaskResultResponse) SetBody(v *GetDeepWriteTaskResultResponseBody) *GetDeepWriteTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetDeepWriteTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
