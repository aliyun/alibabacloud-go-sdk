// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyRelatedApprovalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyRelatedApprovalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyRelatedApprovalsResponse
	GetStatusCode() *int32
	SetBody(v *ListMyRelatedApprovalsResponseBody) *ListMyRelatedApprovalsResponse
	GetBody() *ListMyRelatedApprovalsResponseBody
}

type ListMyRelatedApprovalsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyRelatedApprovalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyRelatedApprovalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponse) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyRelatedApprovalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyRelatedApprovalsResponse) GetBody() *ListMyRelatedApprovalsResponseBody {
	return s.Body
}

func (s *ListMyRelatedApprovalsResponse) SetHeaders(v map[string]*string) *ListMyRelatedApprovalsResponse {
	s.Headers = v
	return s
}

func (s *ListMyRelatedApprovalsResponse) SetStatusCode(v int32) *ListMyRelatedApprovalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyRelatedApprovalsResponse) SetBody(v *ListMyRelatedApprovalsResponseBody) *ListMyRelatedApprovalsResponse {
	s.Body = v
	return s
}

func (s *ListMyRelatedApprovalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
