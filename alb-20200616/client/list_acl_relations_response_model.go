// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAclRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAclRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAclRelationsResponseBody) *ListAclRelationsResponse
	GetBody() *ListAclRelationsResponseBody
}

type ListAclRelationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAclRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAclRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAclRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAclRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAclRelationsResponse) GetBody() *ListAclRelationsResponseBody {
	return s.Body
}

func (s *ListAclRelationsResponse) SetHeaders(v map[string]*string) *ListAclRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListAclRelationsResponse) SetStatusCode(v int32) *ListAclRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclRelationsResponse) SetBody(v *ListAclRelationsResponseBody) *ListAclRelationsResponse {
	s.Body = v
	return s
}

func (s *ListAclRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
