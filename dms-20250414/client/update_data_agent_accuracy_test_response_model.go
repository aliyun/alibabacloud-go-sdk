// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentAccuracyTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataAgentAccuracyTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataAgentAccuracyTestResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataAgentAccuracyTestResponseBody) *UpdateDataAgentAccuracyTestResponse
	GetBody() *UpdateDataAgentAccuracyTestResponseBody
}

type UpdateDataAgentAccuracyTestResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataAgentAccuracyTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataAgentAccuracyTestResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentAccuracyTestResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentAccuracyTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataAgentAccuracyTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataAgentAccuracyTestResponse) GetBody() *UpdateDataAgentAccuracyTestResponseBody {
	return s.Body
}

func (s *UpdateDataAgentAccuracyTestResponse) SetHeaders(v map[string]*string) *UpdateDataAgentAccuracyTestResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponse) SetStatusCode(v int32) *UpdateDataAgentAccuracyTestResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponse) SetBody(v *UpdateDataAgentAccuracyTestResponseBody) *UpdateDataAgentAccuracyTestResponse {
	s.Body = v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
