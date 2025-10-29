// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAckNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAckNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAckNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *ListAckNamespacesResponseBody) *ListAckNamespacesResponse
	GetBody() *ListAckNamespacesResponseBody
}

type ListAckNamespacesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAckNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAckNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAckNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAckNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAckNamespacesResponse) GetBody() *ListAckNamespacesResponseBody {
	return s.Body
}

func (s *ListAckNamespacesResponse) SetHeaders(v map[string]*string) *ListAckNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListAckNamespacesResponse) SetStatusCode(v int32) *ListAckNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAckNamespacesResponse) SetBody(v *ListAckNamespacesResponseBody) *ListAckNamespacesResponse {
	s.Body = v
	return s
}

func (s *ListAckNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
