// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToPrePaidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformToPrePaidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformToPrePaidResponse
	GetStatusCode() *int32
	SetBody(v *TransformToPrePaidResponseBody) *TransformToPrePaidResponse
	GetBody() *TransformToPrePaidResponseBody
}

type TransformToPrePaidResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformToPrePaidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformToPrePaidResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformToPrePaidResponse) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformToPrePaidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformToPrePaidResponse) GetBody() *TransformToPrePaidResponseBody {
	return s.Body
}

func (s *TransformToPrePaidResponse) SetHeaders(v map[string]*string) *TransformToPrePaidResponse {
	s.Headers = v
	return s
}

func (s *TransformToPrePaidResponse) SetStatusCode(v int32) *TransformToPrePaidResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformToPrePaidResponse) SetBody(v *TransformToPrePaidResponseBody) *TransformToPrePaidResponse {
	s.Body = v
	return s
}

func (s *TransformToPrePaidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
