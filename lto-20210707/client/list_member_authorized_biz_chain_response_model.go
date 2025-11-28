// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberAuthorizedBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemberAuthorizedBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemberAuthorizedBizChainResponse
	GetStatusCode() *int32
	SetBody(v *ListMemberAuthorizedBizChainResponseBody) *ListMemberAuthorizedBizChainResponse
	GetBody() *ListMemberAuthorizedBizChainResponseBody
}

type ListMemberAuthorizedBizChainResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemberAuthorizedBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemberAuthorizedBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAuthorizedBizChainResponse) GoString() string {
	return s.String()
}

func (s *ListMemberAuthorizedBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemberAuthorizedBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemberAuthorizedBizChainResponse) GetBody() *ListMemberAuthorizedBizChainResponseBody {
	return s.Body
}

func (s *ListMemberAuthorizedBizChainResponse) SetHeaders(v map[string]*string) *ListMemberAuthorizedBizChainResponse {
	s.Headers = v
	return s
}

func (s *ListMemberAuthorizedBizChainResponse) SetStatusCode(v int32) *ListMemberAuthorizedBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponse) SetBody(v *ListMemberAuthorizedBizChainResponseBody) *ListMemberAuthorizedBizChainResponse {
	s.Body = v
	return s
}

func (s *ListMemberAuthorizedBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
