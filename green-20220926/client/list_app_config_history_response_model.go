// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConfigHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppConfigHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppConfigHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListAppConfigHistoryResponseBody) *ListAppConfigHistoryResponse
	GetBody() *ListAppConfigHistoryResponseBody
}

type ListAppConfigHistoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppConfigHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListAppConfigHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppConfigHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppConfigHistoryResponse) GetBody() *ListAppConfigHistoryResponseBody {
	return s.Body
}

func (s *ListAppConfigHistoryResponse) SetHeaders(v map[string]*string) *ListAppConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListAppConfigHistoryResponse) SetStatusCode(v int32) *ListAppConfigHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppConfigHistoryResponse) SetBody(v *ListAppConfigHistoryResponseBody) *ListAppConfigHistoryResponse {
	s.Body = v
	return s
}

func (s *ListAppConfigHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
