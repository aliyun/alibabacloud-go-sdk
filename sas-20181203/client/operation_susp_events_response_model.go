// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationSuspEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperationSuspEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperationSuspEventsResponse
	GetStatusCode() *int32
	SetBody(v *OperationSuspEventsResponseBody) *OperationSuspEventsResponse
	GetBody() *OperationSuspEventsResponseBody
}

type OperationSuspEventsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperationSuspEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperationSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *OperationSuspEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperationSuspEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperationSuspEventsResponse) GetBody() *OperationSuspEventsResponseBody {
	return s.Body
}

func (s *OperationSuspEventsResponse) SetHeaders(v map[string]*string) *OperationSuspEventsResponse {
	s.Headers = v
	return s
}

func (s *OperationSuspEventsResponse) SetStatusCode(v int32) *OperationSuspEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperationSuspEventsResponse) SetBody(v *OperationSuspEventsResponseBody) *OperationSuspEventsResponse {
	s.Body = v
	return s
}

func (s *OperationSuspEventsResponse) Validate() error {
	return dara.Validate(s)
}
