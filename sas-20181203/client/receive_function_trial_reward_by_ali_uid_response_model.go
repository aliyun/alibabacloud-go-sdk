// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveFunctionTrialRewardByAliUidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReceiveFunctionTrialRewardByAliUidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReceiveFunctionTrialRewardByAliUidResponse
	GetStatusCode() *int32
	SetBody(v *ReceiveFunctionTrialRewardByAliUidResponseBody) *ReceiveFunctionTrialRewardByAliUidResponse
	GetBody() *ReceiveFunctionTrialRewardByAliUidResponseBody
}

type ReceiveFunctionTrialRewardByAliUidResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReceiveFunctionTrialRewardByAliUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReceiveFunctionTrialRewardByAliUidResponse) String() string {
	return dara.Prettify(s)
}

func (s ReceiveFunctionTrialRewardByAliUidResponse) GoString() string {
	return s.String()
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) GetBody() *ReceiveFunctionTrialRewardByAliUidResponseBody {
	return s.Body
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) SetHeaders(v map[string]*string) *ReceiveFunctionTrialRewardByAliUidResponse {
	s.Headers = v
	return s
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) SetStatusCode(v int32) *ReceiveFunctionTrialRewardByAliUidResponse {
	s.StatusCode = &v
	return s
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) SetBody(v *ReceiveFunctionTrialRewardByAliUidResponseBody) *ReceiveFunctionTrialRewardByAliUidResponse {
	s.Body = v
	return s
}

func (s *ReceiveFunctionTrialRewardByAliUidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
