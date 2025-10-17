// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJobStepCheckpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyJobStepCheckpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyJobStepCheckpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyJobStepCheckpointResponseBody) *ModifyJobStepCheckpointResponse
	GetBody() *ModifyJobStepCheckpointResponseBody
}

type ModifyJobStepCheckpointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyJobStepCheckpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyJobStepCheckpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobStepCheckpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyJobStepCheckpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyJobStepCheckpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyJobStepCheckpointResponse) GetBody() *ModifyJobStepCheckpointResponseBody {
	return s.Body
}

func (s *ModifyJobStepCheckpointResponse) SetHeaders(v map[string]*string) *ModifyJobStepCheckpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyJobStepCheckpointResponse) SetStatusCode(v int32) *ModifyJobStepCheckpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyJobStepCheckpointResponse) SetBody(v *ModifyJobStepCheckpointResponseBody) *ModifyJobStepCheckpointResponse {
	s.Body = v
	return s
}

func (s *ModifyJobStepCheckpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
