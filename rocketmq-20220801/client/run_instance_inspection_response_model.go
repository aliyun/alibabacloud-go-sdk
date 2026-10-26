// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstanceInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunInstanceInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunInstanceInspectionResponse
	GetStatusCode() *int32
	SetBody(v *RunInstanceInspectionResponseBody) *RunInstanceInspectionResponse
	GetBody() *RunInstanceInspectionResponseBody
}

type RunInstanceInspectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunInstanceInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunInstanceInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunInstanceInspectionResponse) GoString() string {
	return s.String()
}

func (s *RunInstanceInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunInstanceInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunInstanceInspectionResponse) GetBody() *RunInstanceInspectionResponseBody {
	return s.Body
}

func (s *RunInstanceInspectionResponse) SetHeaders(v map[string]*string) *RunInstanceInspectionResponse {
	s.Headers = v
	return s
}

func (s *RunInstanceInspectionResponse) SetStatusCode(v int32) *RunInstanceInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunInstanceInspectionResponse) SetBody(v *RunInstanceInspectionResponseBody) *RunInstanceInspectionResponse {
	s.Body = v
	return s
}

func (s *RunInstanceInspectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
