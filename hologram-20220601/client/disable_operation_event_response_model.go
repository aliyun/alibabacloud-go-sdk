// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOperationEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableOperationEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableOperationEventResponse
	GetStatusCode() *int32
	SetBody(v *DisableOperationEventResponseBody) *DisableOperationEventResponse
	GetBody() *DisableOperationEventResponseBody
}

type DisableOperationEventResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableOperationEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableOperationEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableOperationEventResponse) GoString() string {
	return s.String()
}

func (s *DisableOperationEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableOperationEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableOperationEventResponse) GetBody() *DisableOperationEventResponseBody {
	return s.Body
}

func (s *DisableOperationEventResponse) SetHeaders(v map[string]*string) *DisableOperationEventResponse {
	s.Headers = v
	return s
}

func (s *DisableOperationEventResponse) SetStatusCode(v int32) *DisableOperationEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableOperationEventResponse) SetBody(v *DisableOperationEventResponseBody) *DisableOperationEventResponse {
	s.Body = v
	return s
}

func (s *DisableOperationEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
