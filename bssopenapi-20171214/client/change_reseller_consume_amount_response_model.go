// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResellerConsumeAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeResellerConsumeAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeResellerConsumeAmountResponse
	GetStatusCode() *int32
	SetBody(v *ChangeResellerConsumeAmountResponseBody) *ChangeResellerConsumeAmountResponse
	GetBody() *ChangeResellerConsumeAmountResponseBody
}

type ChangeResellerConsumeAmountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResellerConsumeAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResellerConsumeAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeResellerConsumeAmountResponse) GoString() string {
	return s.String()
}

func (s *ChangeResellerConsumeAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeResellerConsumeAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeResellerConsumeAmountResponse) GetBody() *ChangeResellerConsumeAmountResponseBody {
	return s.Body
}

func (s *ChangeResellerConsumeAmountResponse) SetHeaders(v map[string]*string) *ChangeResellerConsumeAmountResponse {
	s.Headers = v
	return s
}

func (s *ChangeResellerConsumeAmountResponse) SetStatusCode(v int32) *ChangeResellerConsumeAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponse) SetBody(v *ChangeResellerConsumeAmountResponseBody) *ChangeResellerConsumeAmountResponse {
	s.Body = v
	return s
}

func (s *ChangeResellerConsumeAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
