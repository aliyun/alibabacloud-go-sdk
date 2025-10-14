// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppLogRiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEdgeContainerAppLogRiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEdgeContainerAppLogRiverResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEdgeContainerAppLogRiverResponseBody) *UpdateEdgeContainerAppLogRiverResponse
	GetBody() *UpdateEdgeContainerAppLogRiverResponseBody
}

type UpdateEdgeContainerAppLogRiverResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEdgeContainerAppLogRiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEdgeContainerAppLogRiverResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppLogRiverResponse) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppLogRiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEdgeContainerAppLogRiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEdgeContainerAppLogRiverResponse) GetBody() *UpdateEdgeContainerAppLogRiverResponseBody {
	return s.Body
}

func (s *UpdateEdgeContainerAppLogRiverResponse) SetHeaders(v map[string]*string) *UpdateEdgeContainerAppLogRiverResponse {
	s.Headers = v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverResponse) SetStatusCode(v int32) *UpdateEdgeContainerAppLogRiverResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverResponse) SetBody(v *UpdateEdgeContainerAppLogRiverResponseBody) *UpdateEdgeContainerAppLogRiverResponse {
	s.Body = v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
