// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataMaskingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDataMaskingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDataMaskingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CheckDataMaskingInstanceResponseBody) *CheckDataMaskingInstanceResponse
	GetBody() *CheckDataMaskingInstanceResponseBody
}

type CheckDataMaskingInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataMaskingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataMaskingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDataMaskingInstanceResponse) GoString() string {
	return s.String()
}

func (s *CheckDataMaskingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDataMaskingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDataMaskingInstanceResponse) GetBody() *CheckDataMaskingInstanceResponseBody {
	return s.Body
}

func (s *CheckDataMaskingInstanceResponse) SetHeaders(v map[string]*string) *CheckDataMaskingInstanceResponse {
	s.Headers = v
	return s
}

func (s *CheckDataMaskingInstanceResponse) SetStatusCode(v int32) *CheckDataMaskingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataMaskingInstanceResponse) SetBody(v *CheckDataMaskingInstanceResponseBody) *CheckDataMaskingInstanceResponse {
	s.Body = v
	return s
}

func (s *CheckDataMaskingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
