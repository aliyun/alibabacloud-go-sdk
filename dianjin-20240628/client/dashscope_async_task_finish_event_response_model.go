// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashscopeAsyncTaskFinishEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DashscopeAsyncTaskFinishEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DashscopeAsyncTaskFinishEventResponse
	GetStatusCode() *int32
	SetBody(v *DashscopeAsyncTaskFinishEventResponseBody) *DashscopeAsyncTaskFinishEventResponse
	GetBody() *DashscopeAsyncTaskFinishEventResponseBody
}

type DashscopeAsyncTaskFinishEventResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DashscopeAsyncTaskFinishEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DashscopeAsyncTaskFinishEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DashscopeAsyncTaskFinishEventResponse) GoString() string {
	return s.String()
}

func (s *DashscopeAsyncTaskFinishEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DashscopeAsyncTaskFinishEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DashscopeAsyncTaskFinishEventResponse) GetBody() *DashscopeAsyncTaskFinishEventResponseBody {
	return s.Body
}

func (s *DashscopeAsyncTaskFinishEventResponse) SetHeaders(v map[string]*string) *DashscopeAsyncTaskFinishEventResponse {
	s.Headers = v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponse) SetStatusCode(v int32) *DashscopeAsyncTaskFinishEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponse) SetBody(v *DashscopeAsyncTaskFinishEventResponseBody) *DashscopeAsyncTaskFinishEventResponse {
	s.Body = v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
