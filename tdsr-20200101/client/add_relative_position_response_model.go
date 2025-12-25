// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRelativePositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRelativePositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRelativePositionResponse
	GetStatusCode() *int32
	SetBody(v *AddRelativePositionResponseBody) *AddRelativePositionResponse
	GetBody() *AddRelativePositionResponseBody
}

type AddRelativePositionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRelativePositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRelativePositionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRelativePositionResponse) GoString() string {
	return s.String()
}

func (s *AddRelativePositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRelativePositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRelativePositionResponse) GetBody() *AddRelativePositionResponseBody {
	return s.Body
}

func (s *AddRelativePositionResponse) SetHeaders(v map[string]*string) *AddRelativePositionResponse {
	s.Headers = v
	return s
}

func (s *AddRelativePositionResponse) SetStatusCode(v int32) *AddRelativePositionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRelativePositionResponse) SetBody(v *AddRelativePositionResponseBody) *AddRelativePositionResponse {
	s.Body = v
	return s
}

func (s *AddRelativePositionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
