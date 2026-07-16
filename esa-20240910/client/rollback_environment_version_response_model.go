// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackEnvironmentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackEnvironmentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackEnvironmentVersionResponse
	GetStatusCode() *int32
	SetBody(v *RollbackEnvironmentVersionResponseBody) *RollbackEnvironmentVersionResponse
	GetBody() *RollbackEnvironmentVersionResponseBody
}

type RollbackEnvironmentVersionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackEnvironmentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackEnvironmentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackEnvironmentVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackEnvironmentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackEnvironmentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackEnvironmentVersionResponse) GetBody() *RollbackEnvironmentVersionResponseBody {
	return s.Body
}

func (s *RollbackEnvironmentVersionResponse) SetHeaders(v map[string]*string) *RollbackEnvironmentVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackEnvironmentVersionResponse) SetStatusCode(v int32) *RollbackEnvironmentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackEnvironmentVersionResponse) SetBody(v *RollbackEnvironmentVersionResponseBody) *RollbackEnvironmentVersionResponse {
	s.Body = v
	return s
}

func (s *RollbackEnvironmentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
