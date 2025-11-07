// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartNisInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartNisInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartNisInspectionTaskResponseBody) *StartNisInspectionTaskResponse
	GetBody() *StartNisInspectionTaskResponseBody
}

type StartNisInspectionTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartNisInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartNisInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartNisInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *StartNisInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartNisInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartNisInspectionTaskResponse) GetBody() *StartNisInspectionTaskResponseBody {
	return s.Body
}

func (s *StartNisInspectionTaskResponse) SetHeaders(v map[string]*string) *StartNisInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *StartNisInspectionTaskResponse) SetStatusCode(v int32) *StartNisInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartNisInspectionTaskResponse) SetBody(v *StartNisInspectionTaskResponseBody) *StartNisInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *StartNisInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
