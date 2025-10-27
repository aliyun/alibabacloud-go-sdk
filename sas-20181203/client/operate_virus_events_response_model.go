// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateVirusEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateVirusEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateVirusEventsResponse
	GetStatusCode() *int32
	SetBody(v *OperateVirusEventsResponseBody) *OperateVirusEventsResponse
	GetBody() *OperateVirusEventsResponseBody
}

type OperateVirusEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateVirusEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateVirusEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateVirusEventsResponse) GoString() string {
	return s.String()
}

func (s *OperateVirusEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateVirusEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateVirusEventsResponse) GetBody() *OperateVirusEventsResponseBody {
	return s.Body
}

func (s *OperateVirusEventsResponse) SetHeaders(v map[string]*string) *OperateVirusEventsResponse {
	s.Headers = v
	return s
}

func (s *OperateVirusEventsResponse) SetStatusCode(v int32) *OperateVirusEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateVirusEventsResponse) SetBody(v *OperateVirusEventsResponseBody) *OperateVirusEventsResponse {
	s.Body = v
	return s
}

func (s *OperateVirusEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
