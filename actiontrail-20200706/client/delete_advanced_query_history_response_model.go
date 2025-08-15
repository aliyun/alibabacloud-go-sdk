// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvancedQueryHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAdvancedQueryHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAdvancedQueryHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAdvancedQueryHistoryResponseBody) *DeleteAdvancedQueryHistoryResponse
	GetBody() *DeleteAdvancedQueryHistoryResponseBody
}

type DeleteAdvancedQueryHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdvancedQueryHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdvancedQueryHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvancedQueryHistoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdvancedQueryHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAdvancedQueryHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAdvancedQueryHistoryResponse) GetBody() *DeleteAdvancedQueryHistoryResponseBody {
	return s.Body
}

func (s *DeleteAdvancedQueryHistoryResponse) SetHeaders(v map[string]*string) *DeleteAdvancedQueryHistoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdvancedQueryHistoryResponse) SetStatusCode(v int32) *DeleteAdvancedQueryHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdvancedQueryHistoryResponse) SetBody(v *DeleteAdvancedQueryHistoryResponseBody) *DeleteAdvancedQueryHistoryResponse {
	s.Body = v
	return s
}

func (s *DeleteAdvancedQueryHistoryResponse) Validate() error {
	return dara.Validate(s)
}
