// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDisasterRecoveryItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDisasterRecoveryItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListDisasterRecoveryItemsResponseBody) *ListDisasterRecoveryItemsResponse
	GetBody() *ListDisasterRecoveryItemsResponseBody
}

type ListDisasterRecoveryItemsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisasterRecoveryItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisasterRecoveryItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryItemsResponse) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDisasterRecoveryItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDisasterRecoveryItemsResponse) GetBody() *ListDisasterRecoveryItemsResponseBody {
	return s.Body
}

func (s *ListDisasterRecoveryItemsResponse) SetHeaders(v map[string]*string) *ListDisasterRecoveryItemsResponse {
	s.Headers = v
	return s
}

func (s *ListDisasterRecoveryItemsResponse) SetStatusCode(v int32) *ListDisasterRecoveryItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponse) SetBody(v *ListDisasterRecoveryItemsResponseBody) *ListDisasterRecoveryItemsResponse {
	s.Body = v
	return s
}

func (s *ListDisasterRecoveryItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
