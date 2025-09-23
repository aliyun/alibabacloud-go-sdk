// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceShareAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceShareAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceShareAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceShareAssociationsResponseBody) *ListResourceShareAssociationsResponse
	GetBody() *ListResourceShareAssociationsResponseBody
}

type ListResourceShareAssociationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceShareAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceShareAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceShareAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceShareAssociationsResponse) GetBody() *ListResourceShareAssociationsResponseBody {
	return s.Body
}

func (s *ListResourceShareAssociationsResponse) SetHeaders(v map[string]*string) *ListResourceShareAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceShareAssociationsResponse) SetStatusCode(v int32) *ListResourceShareAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceShareAssociationsResponse) SetBody(v *ListResourceShareAssociationsResponseBody) *ListResourceShareAssociationsResponse {
	s.Body = v
	return s
}

func (s *ListResourceShareAssociationsResponse) Validate() error {
	return dara.Validate(s)
}
