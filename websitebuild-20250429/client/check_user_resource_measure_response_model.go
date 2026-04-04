// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserResourceMeasureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUserResourceMeasureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUserResourceMeasureResponse
	GetStatusCode() *int32
	SetBody(v *CheckUserResourceMeasureResponseBody) *CheckUserResourceMeasureResponse
	GetBody() *CheckUserResourceMeasureResponseBody
}

type CheckUserResourceMeasureResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserResourceMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserResourceMeasureResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserResourceMeasureResponse) GoString() string {
	return s.String()
}

func (s *CheckUserResourceMeasureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUserResourceMeasureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUserResourceMeasureResponse) GetBody() *CheckUserResourceMeasureResponseBody {
	return s.Body
}

func (s *CheckUserResourceMeasureResponse) SetHeaders(v map[string]*string) *CheckUserResourceMeasureResponse {
	s.Headers = v
	return s
}

func (s *CheckUserResourceMeasureResponse) SetStatusCode(v int32) *CheckUserResourceMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponse) SetBody(v *CheckUserResourceMeasureResponseBody) *CheckUserResourceMeasureResponse {
	s.Body = v
	return s
}

func (s *CheckUserResourceMeasureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
