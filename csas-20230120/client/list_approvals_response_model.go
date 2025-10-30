// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApprovalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApprovalsResponse
	GetStatusCode() *int32
	SetBody(v *ListApprovalsResponseBody) *ListApprovalsResponse
	GetBody() *ListApprovalsResponseBody
}

type ListApprovalsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApprovalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApprovalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalsResponse) GoString() string {
	return s.String()
}

func (s *ListApprovalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApprovalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApprovalsResponse) GetBody() *ListApprovalsResponseBody {
	return s.Body
}

func (s *ListApprovalsResponse) SetHeaders(v map[string]*string) *ListApprovalsResponse {
	s.Headers = v
	return s
}

func (s *ListApprovalsResponse) SetStatusCode(v int32) *ListApprovalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApprovalsResponse) SetBody(v *ListApprovalsResponseBody) *ListApprovalsResponse {
	s.Body = v
	return s
}

func (s *ListApprovalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
