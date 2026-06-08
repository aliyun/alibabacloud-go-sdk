// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPendingApprovalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPendingApprovalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPendingApprovalsResponse
	GetStatusCode() *int32
	SetBody(v *ListPendingApprovalsResponseBody) *ListPendingApprovalsResponse
	GetBody() *ListPendingApprovalsResponseBody
}

type ListPendingApprovalsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPendingApprovalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPendingApprovalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponse) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPendingApprovalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPendingApprovalsResponse) GetBody() *ListPendingApprovalsResponseBody {
	return s.Body
}

func (s *ListPendingApprovalsResponse) SetHeaders(v map[string]*string) *ListPendingApprovalsResponse {
	s.Headers = v
	return s
}

func (s *ListPendingApprovalsResponse) SetStatusCode(v int32) *ListPendingApprovalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPendingApprovalsResponse) SetBody(v *ListPendingApprovalsResponseBody) *ListPendingApprovalsResponse {
	s.Body = v
	return s
}

func (s *ListPendingApprovalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
