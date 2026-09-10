// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *AddDataCheckTaskResponseBody) *AddDataCheckTaskResponse
	GetBody() *AddDataCheckTaskResponseBody
}

type AddDataCheckTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *AddDataCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataCheckTaskResponse) GetBody() *AddDataCheckTaskResponseBody {
	return s.Body
}

func (s *AddDataCheckTaskResponse) SetHeaders(v map[string]*string) *AddDataCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *AddDataCheckTaskResponse) SetStatusCode(v int32) *AddDataCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataCheckTaskResponse) SetBody(v *AddDataCheckTaskResponseBody) *AddDataCheckTaskResponse {
	s.Body = v
	return s
}

func (s *AddDataCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
