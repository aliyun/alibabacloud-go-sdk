// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMultiDimTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMultiDimTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMultiDimTableResponse
	GetStatusCode() *int32
	SetBody(v *AddMultiDimTableResponseBody) *AddMultiDimTableResponse
	GetBody() *AddMultiDimTableResponseBody
}

type AddMultiDimTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMultiDimTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMultiDimTableResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableResponse) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMultiDimTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMultiDimTableResponse) GetBody() *AddMultiDimTableResponseBody {
	return s.Body
}

func (s *AddMultiDimTableResponse) SetHeaders(v map[string]*string) *AddMultiDimTableResponse {
	s.Headers = v
	return s
}

func (s *AddMultiDimTableResponse) SetStatusCode(v int32) *AddMultiDimTableResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMultiDimTableResponse) SetBody(v *AddMultiDimTableResponseBody) *AddMultiDimTableResponse {
	s.Body = v
	return s
}

func (s *AddMultiDimTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
