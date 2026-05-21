// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOperationEventScheduleTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOperationEventScheduleTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOperationEventScheduleTimeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOperationEventScheduleTimeResponseBody) *UpdateOperationEventScheduleTimeResponse
	GetBody() *UpdateOperationEventScheduleTimeResponseBody
}

type UpdateOperationEventScheduleTimeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOperationEventScheduleTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOperationEventScheduleTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOperationEventScheduleTimeResponse) GoString() string {
	return s.String()
}

func (s *UpdateOperationEventScheduleTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOperationEventScheduleTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOperationEventScheduleTimeResponse) GetBody() *UpdateOperationEventScheduleTimeResponseBody {
	return s.Body
}

func (s *UpdateOperationEventScheduleTimeResponse) SetHeaders(v map[string]*string) *UpdateOperationEventScheduleTimeResponse {
	s.Headers = v
	return s
}

func (s *UpdateOperationEventScheduleTimeResponse) SetStatusCode(v int32) *UpdateOperationEventScheduleTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOperationEventScheduleTimeResponse) SetBody(v *UpdateOperationEventScheduleTimeResponseBody) *UpdateOperationEventScheduleTimeResponse {
	s.Body = v
	return s
}

func (s *UpdateOperationEventScheduleTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
