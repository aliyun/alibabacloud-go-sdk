// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkNodeWorkforceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkNodeWorkforceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkNodeWorkforceResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkNodeWorkforceResponseBody) *AddWorkNodeWorkforceResponse
	GetBody() *AddWorkNodeWorkforceResponseBody
}

type AddWorkNodeWorkforceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkNodeWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkNodeWorkforceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkNodeWorkforceResponse) GoString() string {
	return s.String()
}

func (s *AddWorkNodeWorkforceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkNodeWorkforceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkNodeWorkforceResponse) GetBody() *AddWorkNodeWorkforceResponseBody {
	return s.Body
}

func (s *AddWorkNodeWorkforceResponse) SetHeaders(v map[string]*string) *AddWorkNodeWorkforceResponse {
	s.Headers = v
	return s
}

func (s *AddWorkNodeWorkforceResponse) SetStatusCode(v int32) *AddWorkNodeWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkNodeWorkforceResponse) SetBody(v *AddWorkNodeWorkforceResponseBody) *AddWorkNodeWorkforceResponse {
	s.Body = v
	return s
}

func (s *AddWorkNodeWorkforceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
