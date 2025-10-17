// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSaasInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSaasInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListSaasInfoResponseBody) *ListSaasInfoResponse
	GetBody() *ListSaasInfoResponseBody
}

type ListSaasInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSaasInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSaasInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSaasInfoResponse) GoString() string {
	return s.String()
}

func (s *ListSaasInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSaasInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSaasInfoResponse) GetBody() *ListSaasInfoResponseBody {
	return s.Body
}

func (s *ListSaasInfoResponse) SetHeaders(v map[string]*string) *ListSaasInfoResponse {
	s.Headers = v
	return s
}

func (s *ListSaasInfoResponse) SetStatusCode(v int32) *ListSaasInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSaasInfoResponse) SetBody(v *ListSaasInfoResponseBody) *ListSaasInfoResponse {
	s.Body = v
	return s
}

func (s *ListSaasInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
