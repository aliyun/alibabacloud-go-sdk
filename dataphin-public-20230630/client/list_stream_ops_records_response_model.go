// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamOpsRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamOpsRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamOpsRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamOpsRecordsResponseBody) *ListStreamOpsRecordsResponse
	GetBody() *ListStreamOpsRecordsResponseBody
}

type ListStreamOpsRecordsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamOpsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamOpsRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamOpsRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListStreamOpsRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamOpsRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamOpsRecordsResponse) GetBody() *ListStreamOpsRecordsResponseBody {
	return s.Body
}

func (s *ListStreamOpsRecordsResponse) SetHeaders(v map[string]*string) *ListStreamOpsRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListStreamOpsRecordsResponse) SetStatusCode(v int32) *ListStreamOpsRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamOpsRecordsResponse) SetBody(v *ListStreamOpsRecordsResponseBody) *ListStreamOpsRecordsResponse {
	s.Body = v
	return s
}

func (s *ListStreamOpsRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
