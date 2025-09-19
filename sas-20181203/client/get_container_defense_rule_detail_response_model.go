// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContainerDefenseRuleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContainerDefenseRuleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContainerDefenseRuleDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetContainerDefenseRuleDetailResponseBody) *GetContainerDefenseRuleDetailResponse
	GetBody() *GetContainerDefenseRuleDetailResponseBody
}

type GetContainerDefenseRuleDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContainerDefenseRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContainerDefenseRuleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContainerDefenseRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetContainerDefenseRuleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContainerDefenseRuleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContainerDefenseRuleDetailResponse) GetBody() *GetContainerDefenseRuleDetailResponseBody {
	return s.Body
}

func (s *GetContainerDefenseRuleDetailResponse) SetHeaders(v map[string]*string) *GetContainerDefenseRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponse) SetStatusCode(v int32) *GetContainerDefenseRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponse) SetBody(v *GetContainerDefenseRuleDetailResponseBody) *GetContainerDefenseRuleDetailResponse {
	s.Body = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponse) Validate() error {
	return dara.Validate(s)
}
