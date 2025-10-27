// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSuspEventQuaraFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackSuspEventQuaraFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackSuspEventQuaraFileResponse
	GetStatusCode() *int32
	SetBody(v *RollbackSuspEventQuaraFileResponseBody) *RollbackSuspEventQuaraFileResponse
	GetBody() *RollbackSuspEventQuaraFileResponseBody
}

type RollbackSuspEventQuaraFileResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackSuspEventQuaraFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackSuspEventQuaraFileResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponse) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackSuspEventQuaraFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackSuspEventQuaraFileResponse) GetBody() *RollbackSuspEventQuaraFileResponseBody {
	return s.Body
}

func (s *RollbackSuspEventQuaraFileResponse) SetHeaders(v map[string]*string) *RollbackSuspEventQuaraFileResponse {
	s.Headers = v
	return s
}

func (s *RollbackSuspEventQuaraFileResponse) SetStatusCode(v int32) *RollbackSuspEventQuaraFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackSuspEventQuaraFileResponse) SetBody(v *RollbackSuspEventQuaraFileResponseBody) *RollbackSuspEventQuaraFileResponse {
	s.Body = v
	return s
}

func (s *RollbackSuspEventQuaraFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
