// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecalculateCarbonEmissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecalculateCarbonEmissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecalculateCarbonEmissionResponse
	GetStatusCode() *int32
	SetBody(v *RecalculateCarbonEmissionResponseBody) *RecalculateCarbonEmissionResponse
	GetBody() *RecalculateCarbonEmissionResponseBody
}

type RecalculateCarbonEmissionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecalculateCarbonEmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecalculateCarbonEmissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RecalculateCarbonEmissionResponse) GoString() string {
	return s.String()
}

func (s *RecalculateCarbonEmissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecalculateCarbonEmissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecalculateCarbonEmissionResponse) GetBody() *RecalculateCarbonEmissionResponseBody {
	return s.Body
}

func (s *RecalculateCarbonEmissionResponse) SetHeaders(v map[string]*string) *RecalculateCarbonEmissionResponse {
	s.Headers = v
	return s
}

func (s *RecalculateCarbonEmissionResponse) SetStatusCode(v int32) *RecalculateCarbonEmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecalculateCarbonEmissionResponse) SetBody(v *RecalculateCarbonEmissionResponseBody) *RecalculateCarbonEmissionResponse {
	s.Body = v
	return s
}

func (s *RecalculateCarbonEmissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
