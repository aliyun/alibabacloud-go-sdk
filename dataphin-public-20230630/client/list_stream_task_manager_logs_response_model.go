// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamTaskManagerLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamTaskManagerLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamTaskManagerLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamTaskManagerLogsResponseBody) *ListStreamTaskManagerLogsResponse
	GetBody() *ListStreamTaskManagerLogsResponseBody
}

type ListStreamTaskManagerLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamTaskManagerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamTaskManagerLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagerLogsResponse) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagerLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamTaskManagerLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamTaskManagerLogsResponse) GetBody() *ListStreamTaskManagerLogsResponseBody {
	return s.Body
}

func (s *ListStreamTaskManagerLogsResponse) SetHeaders(v map[string]*string) *ListStreamTaskManagerLogsResponse {
	s.Headers = v
	return s
}

func (s *ListStreamTaskManagerLogsResponse) SetStatusCode(v int32) *ListStreamTaskManagerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponse) SetBody(v *ListStreamTaskManagerLogsResponseBody) *ListStreamTaskManagerLogsResponse {
	s.Body = v
	return s
}

func (s *ListStreamTaskManagerLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
