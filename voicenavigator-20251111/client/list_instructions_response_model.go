// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstructionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstructionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstructionsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstructionsResponseBody) *ListInstructionsResponse
	GetBody() *ListInstructionsResponseBody
}

type ListInstructionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstructionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstructionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstructionsResponse) GoString() string {
	return s.String()
}

func (s *ListInstructionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstructionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstructionsResponse) GetBody() *ListInstructionsResponseBody {
	return s.Body
}

func (s *ListInstructionsResponse) SetHeaders(v map[string]*string) *ListInstructionsResponse {
	s.Headers = v
	return s
}

func (s *ListInstructionsResponse) SetStatusCode(v int32) *ListInstructionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstructionsResponse) SetBody(v *ListInstructionsResponseBody) *ListInstructionsResponse {
	s.Body = v
	return s
}

func (s *ListInstructionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
