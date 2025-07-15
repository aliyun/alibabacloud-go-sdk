// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRebalanceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRebalanceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRebalanceInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListRebalanceInfoResponseBody) *ListRebalanceInfoResponse
	GetBody() *ListRebalanceInfoResponseBody
}

type ListRebalanceInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRebalanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRebalanceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRebalanceInfoResponse) GoString() string {
	return s.String()
}

func (s *ListRebalanceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRebalanceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRebalanceInfoResponse) GetBody() *ListRebalanceInfoResponseBody {
	return s.Body
}

func (s *ListRebalanceInfoResponse) SetHeaders(v map[string]*string) *ListRebalanceInfoResponse {
	s.Headers = v
	return s
}

func (s *ListRebalanceInfoResponse) SetStatusCode(v int32) *ListRebalanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRebalanceInfoResponse) SetBody(v *ListRebalanceInfoResponseBody) *ListRebalanceInfoResponse {
	s.Body = v
	return s
}

func (s *ListRebalanceInfoResponse) Validate() error {
	return dara.Validate(s)
}
