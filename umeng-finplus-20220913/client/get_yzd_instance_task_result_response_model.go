// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYzdInstanceTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYzdInstanceTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYzdInstanceTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetYzdInstanceTaskResultResponseBody) *GetYzdInstanceTaskResultResponse
	GetBody() *GetYzdInstanceTaskResultResponseBody
}

type GetYzdInstanceTaskResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYzdInstanceTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYzdInstanceTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYzdInstanceTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYzdInstanceTaskResultResponse) GetBody() *GetYzdInstanceTaskResultResponseBody {
	return s.Body
}

func (s *GetYzdInstanceTaskResultResponse) SetHeaders(v map[string]*string) *GetYzdInstanceTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetYzdInstanceTaskResultResponse) SetStatusCode(v int32) *GetYzdInstanceTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponse) SetBody(v *GetYzdInstanceTaskResultResponseBody) *GetYzdInstanceTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetYzdInstanceTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
