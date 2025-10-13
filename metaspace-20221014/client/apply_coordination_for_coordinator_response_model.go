// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationForCoordinatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyCoordinationForCoordinatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyCoordinationForCoordinatorResponse
	GetStatusCode() *int32
	SetBody(v *ApplyCoordinationForCoordinatorResponseBody) *ApplyCoordinationForCoordinatorResponse
	GetBody() *ApplyCoordinationForCoordinatorResponseBody
}

type ApplyCoordinationForCoordinatorResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCoordinationForCoordinatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCoordinationForCoordinatorResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForCoordinatorResponse) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForCoordinatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyCoordinationForCoordinatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyCoordinationForCoordinatorResponse) GetBody() *ApplyCoordinationForCoordinatorResponseBody {
	return s.Body
}

func (s *ApplyCoordinationForCoordinatorResponse) SetHeaders(v map[string]*string) *ApplyCoordinationForCoordinatorResponse {
	s.Headers = v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponse) SetStatusCode(v int32) *ApplyCoordinationForCoordinatorResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponse) SetBody(v *ApplyCoordinationForCoordinatorResponseBody) *ApplyCoordinationForCoordinatorResponse {
	s.Body = v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponse) Validate() error {
	return dara.Validate(s)
}
