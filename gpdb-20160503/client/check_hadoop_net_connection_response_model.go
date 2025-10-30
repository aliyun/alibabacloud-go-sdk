// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHadoopNetConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckHadoopNetConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckHadoopNetConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CheckHadoopNetConnectionResponseBody) *CheckHadoopNetConnectionResponse
	GetBody() *CheckHadoopNetConnectionResponseBody
}

type CheckHadoopNetConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckHadoopNetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckHadoopNetConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckHadoopNetConnectionResponse) GoString() string {
	return s.String()
}

func (s *CheckHadoopNetConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckHadoopNetConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckHadoopNetConnectionResponse) GetBody() *CheckHadoopNetConnectionResponseBody {
	return s.Body
}

func (s *CheckHadoopNetConnectionResponse) SetHeaders(v map[string]*string) *CheckHadoopNetConnectionResponse {
	s.Headers = v
	return s
}

func (s *CheckHadoopNetConnectionResponse) SetStatusCode(v int32) *CheckHadoopNetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckHadoopNetConnectionResponse) SetBody(v *CheckHadoopNetConnectionResponseBody) *CheckHadoopNetConnectionResponse {
	s.Body = v
	return s
}

func (s *CheckHadoopNetConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
