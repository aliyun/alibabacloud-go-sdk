// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpResponseHeaderModificationRuleResponseBody) *DeleteHttpResponseHeaderModificationRuleResponse
	GetBody() *DeleteHttpResponseHeaderModificationRuleResponseBody
}

type DeleteHttpResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) GetBody() *DeleteHttpResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *DeleteHttpResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *DeleteHttpResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) SetBody(v *DeleteHttpResponseHeaderModificationRuleResponseBody) *DeleteHttpResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
