// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddYikeProductionMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddYikeProductionMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddYikeProductionMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddYikeProductionMembersResponseBody) *AddYikeProductionMembersResponse
	GetBody() *AddYikeProductionMembersResponseBody
}

type AddYikeProductionMembersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddYikeProductionMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddYikeProductionMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddYikeProductionMembersResponse) GoString() string {
	return s.String()
}

func (s *AddYikeProductionMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddYikeProductionMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddYikeProductionMembersResponse) GetBody() *AddYikeProductionMembersResponseBody {
	return s.Body
}

func (s *AddYikeProductionMembersResponse) SetHeaders(v map[string]*string) *AddYikeProductionMembersResponse {
	s.Headers = v
	return s
}

func (s *AddYikeProductionMembersResponse) SetStatusCode(v int32) *AddYikeProductionMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddYikeProductionMembersResponse) SetBody(v *AddYikeProductionMembersResponseBody) *AddYikeProductionMembersResponse {
	s.Body = v
	return s
}

func (s *AddYikeProductionMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
