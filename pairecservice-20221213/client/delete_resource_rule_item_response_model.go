// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRuleItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceRuleItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceRuleItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceRuleItemResponseBody) *DeleteResourceRuleItemResponse
	GetBody() *DeleteResourceRuleItemResponseBody
}

type DeleteResourceRuleItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceRuleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceRuleItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRuleItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceRuleItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceRuleItemResponse) GetBody() *DeleteResourceRuleItemResponseBody {
	return s.Body
}

func (s *DeleteResourceRuleItemResponse) SetHeaders(v map[string]*string) *DeleteResourceRuleItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceRuleItemResponse) SetStatusCode(v int32) *DeleteResourceRuleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceRuleItemResponse) SetBody(v *DeleteResourceRuleItemResponseBody) *DeleteResourceRuleItemResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceRuleItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
