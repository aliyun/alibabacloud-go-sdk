// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseRuleCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseRuleCacheResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseRuleCacheResponseBody) *ModifyDefenseRuleCacheResponse
	GetBody() *ModifyDefenseRuleCacheResponseBody
}

type ModifyDefenseRuleCacheResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseRuleCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseRuleCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleCacheResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseRuleCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseRuleCacheResponse) GetBody() *ModifyDefenseRuleCacheResponseBody {
	return s.Body
}

func (s *ModifyDefenseRuleCacheResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleCacheResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleCacheResponse) SetStatusCode(v int32) *ModifyDefenseRuleCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleCacheResponse) SetBody(v *ModifyDefenseRuleCacheResponseBody) *ModifyDefenseRuleCacheResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseRuleCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
