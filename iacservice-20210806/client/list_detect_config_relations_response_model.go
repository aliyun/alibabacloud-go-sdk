// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectConfigRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDetectConfigRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDetectConfigRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListDetectConfigRelationsResponseBody) *ListDetectConfigRelationsResponse
	GetBody() *ListDetectConfigRelationsResponseBody
}

type ListDetectConfigRelationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDetectConfigRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDetectConfigRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListDetectConfigRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDetectConfigRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDetectConfigRelationsResponse) GetBody() *ListDetectConfigRelationsResponseBody {
	return s.Body
}

func (s *ListDetectConfigRelationsResponse) SetHeaders(v map[string]*string) *ListDetectConfigRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListDetectConfigRelationsResponse) SetStatusCode(v int32) *ListDetectConfigRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDetectConfigRelationsResponse) SetBody(v *ListDetectConfigRelationsResponseBody) *ListDetectConfigRelationsResponse {
	s.Body = v
	return s
}

func (s *ListDetectConfigRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
