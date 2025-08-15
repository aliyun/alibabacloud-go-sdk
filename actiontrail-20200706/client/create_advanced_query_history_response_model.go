// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedQueryHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAdvancedQueryHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAdvancedQueryHistoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateAdvancedQueryHistoryResponseBody) *CreateAdvancedQueryHistoryResponse
	GetBody() *CreateAdvancedQueryHistoryResponseBody
}

type CreateAdvancedQueryHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdvancedQueryHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdvancedQueryHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedQueryHistoryResponse) GoString() string {
	return s.String()
}

func (s *CreateAdvancedQueryHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAdvancedQueryHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAdvancedQueryHistoryResponse) GetBody() *CreateAdvancedQueryHistoryResponseBody {
	return s.Body
}

func (s *CreateAdvancedQueryHistoryResponse) SetHeaders(v map[string]*string) *CreateAdvancedQueryHistoryResponse {
	s.Headers = v
	return s
}

func (s *CreateAdvancedQueryHistoryResponse) SetStatusCode(v int32) *CreateAdvancedQueryHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdvancedQueryHistoryResponse) SetBody(v *CreateAdvancedQueryHistoryResponseBody) *CreateAdvancedQueryHistoryResponse {
	s.Body = v
	return s
}

func (s *CreateAdvancedQueryHistoryResponse) Validate() error {
	return dara.Validate(s)
}
