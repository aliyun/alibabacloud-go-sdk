// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggTaskGroupStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggTaskGroupStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggTaskGroupStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggTaskGroupStatusResponseBody) *UpdateAggTaskGroupStatusResponse
	GetBody() *UpdateAggTaskGroupStatusResponseBody
}

type UpdateAggTaskGroupStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggTaskGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggTaskGroupStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggTaskGroupStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggTaskGroupStatusResponse) GetBody() *UpdateAggTaskGroupStatusResponseBody {
	return s.Body
}

func (s *UpdateAggTaskGroupStatusResponse) SetHeaders(v map[string]*string) *UpdateAggTaskGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggTaskGroupStatusResponse) SetStatusCode(v int32) *UpdateAggTaskGroupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponse) SetBody(v *UpdateAggTaskGroupStatusResponseBody) *UpdateAggTaskGroupStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAggTaskGroupStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
