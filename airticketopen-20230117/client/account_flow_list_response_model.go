// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountFlowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountFlowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountFlowListResponse
	GetStatusCode() *int32
	SetBody(v *AccountFlowListResponseBody) *AccountFlowListResponse
	GetBody() *AccountFlowListResponseBody
}

type AccountFlowListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountFlowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountFlowListResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountFlowListResponse) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountFlowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountFlowListResponse) GetBody() *AccountFlowListResponseBody {
	return s.Body
}

func (s *AccountFlowListResponse) SetHeaders(v map[string]*string) *AccountFlowListResponse {
	s.Headers = v
	return s
}

func (s *AccountFlowListResponse) SetStatusCode(v int32) *AccountFlowListResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountFlowListResponse) SetBody(v *AccountFlowListResponseBody) *AccountFlowListResponse {
	s.Body = v
	return s
}

func (s *AccountFlowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
