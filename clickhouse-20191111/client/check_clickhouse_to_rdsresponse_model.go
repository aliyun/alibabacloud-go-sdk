// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckClickhouseToRDSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckClickhouseToRDSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckClickhouseToRDSResponse
	GetStatusCode() *int32
	SetBody(v *CheckClickhouseToRDSResponseBody) *CheckClickhouseToRDSResponse
	GetBody() *CheckClickhouseToRDSResponseBody
}

type CheckClickhouseToRDSResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckClickhouseToRDSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckClickhouseToRDSResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckClickhouseToRDSResponse) GoString() string {
	return s.String()
}

func (s *CheckClickhouseToRDSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckClickhouseToRDSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckClickhouseToRDSResponse) GetBody() *CheckClickhouseToRDSResponseBody {
	return s.Body
}

func (s *CheckClickhouseToRDSResponse) SetHeaders(v map[string]*string) *CheckClickhouseToRDSResponse {
	s.Headers = v
	return s
}

func (s *CheckClickhouseToRDSResponse) SetStatusCode(v int32) *CheckClickhouseToRDSResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckClickhouseToRDSResponse) SetBody(v *CheckClickhouseToRDSResponseBody) *CheckClickhouseToRDSResponse {
	s.Body = v
	return s
}

func (s *CheckClickhouseToRDSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
