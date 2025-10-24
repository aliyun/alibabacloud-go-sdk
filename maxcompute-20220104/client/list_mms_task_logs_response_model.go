// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTaskLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsTaskLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsTaskLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsTaskLogsResponseBody) *ListMmsTaskLogsResponse
	GetBody() *ListMmsTaskLogsResponseBody
}

type ListMmsTaskLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTaskLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTaskLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTaskLogsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTaskLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsTaskLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsTaskLogsResponse) GetBody() *ListMmsTaskLogsResponseBody {
	return s.Body
}

func (s *ListMmsTaskLogsResponse) SetHeaders(v map[string]*string) *ListMmsTaskLogsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTaskLogsResponse) SetStatusCode(v int32) *ListMmsTaskLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTaskLogsResponse) SetBody(v *ListMmsTaskLogsResponseBody) *ListMmsTaskLogsResponse {
	s.Body = v
	return s
}

func (s *ListMmsTaskLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
