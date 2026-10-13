// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPenetrationTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPenetrationTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPenetrationTestResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPenetrationTestResponseBody) *SubmitPenetrationTestResponse
	GetBody() *SubmitPenetrationTestResponseBody
}

type SubmitPenetrationTestResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPenetrationTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPenetrationTestResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPenetrationTestResponse) GoString() string {
	return s.String()
}

func (s *SubmitPenetrationTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPenetrationTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPenetrationTestResponse) GetBody() *SubmitPenetrationTestResponseBody {
	return s.Body
}

func (s *SubmitPenetrationTestResponse) SetHeaders(v map[string]*string) *SubmitPenetrationTestResponse {
	s.Headers = v
	return s
}

func (s *SubmitPenetrationTestResponse) SetStatusCode(v int32) *SubmitPenetrationTestResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPenetrationTestResponse) SetBody(v *SubmitPenetrationTestResponseBody) *SubmitPenetrationTestResponse {
	s.Body = v
	return s
}

func (s *SubmitPenetrationTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
