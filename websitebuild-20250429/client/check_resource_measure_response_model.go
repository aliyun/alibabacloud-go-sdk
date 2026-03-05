// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceMeasureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckResourceMeasureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckResourceMeasureResponse
	GetStatusCode() *int32
	SetBody(v *CheckResourceMeasureResponseBody) *CheckResourceMeasureResponse
	GetBody() *CheckResourceMeasureResponseBody
}

type CheckResourceMeasureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResourceMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResourceMeasureResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceMeasureResponse) GoString() string {
	return s.String()
}

func (s *CheckResourceMeasureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckResourceMeasureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckResourceMeasureResponse) GetBody() *CheckResourceMeasureResponseBody {
	return s.Body
}

func (s *CheckResourceMeasureResponse) SetHeaders(v map[string]*string) *CheckResourceMeasureResponse {
	s.Headers = v
	return s
}

func (s *CheckResourceMeasureResponse) SetStatusCode(v int32) *CheckResourceMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResourceMeasureResponse) SetBody(v *CheckResourceMeasureResponseBody) *CheckResourceMeasureResponse {
	s.Body = v
	return s
}

func (s *CheckResourceMeasureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
