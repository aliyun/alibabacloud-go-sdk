// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhiteRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWhiteRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWhiteRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWhiteRuleListResponseBody) *DeleteWhiteRuleListResponse
	GetBody() *DeleteWhiteRuleListResponseBody
}

type DeleteWhiteRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWhiteRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWhiteRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWhiteRuleListResponse) GetBody() *DeleteWhiteRuleListResponseBody {
	return s.Body
}

func (s *DeleteWhiteRuleListResponse) SetHeaders(v map[string]*string) *DeleteWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhiteRuleListResponse) SetStatusCode(v int32) *DeleteWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWhiteRuleListResponse) SetBody(v *DeleteWhiteRuleListResponseBody) *DeleteWhiteRuleListResponse {
	s.Body = v
	return s
}

func (s *DeleteWhiteRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
