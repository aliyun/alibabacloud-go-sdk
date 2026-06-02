// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawFlowResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawFlowResponseBody) *WithdrawFlowResponse
	GetBody() *WithdrawFlowResponseBody
}

type WithdrawFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawFlowResponse) GoString() string {
	return s.String()
}

func (s *WithdrawFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawFlowResponse) GetBody() *WithdrawFlowResponseBody {
	return s.Body
}

func (s *WithdrawFlowResponse) SetHeaders(v map[string]*string) *WithdrawFlowResponse {
	s.Headers = v
	return s
}

func (s *WithdrawFlowResponse) SetStatusCode(v int32) *WithdrawFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawFlowResponse) SetBody(v *WithdrawFlowResponseBody) *WithdrawFlowResponse {
	s.Body = v
	return s
}

func (s *WithdrawFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
