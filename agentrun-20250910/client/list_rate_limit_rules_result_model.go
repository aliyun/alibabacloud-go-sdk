// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRateLimitRulesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRateLimitRulesResult
	GetCode() *string
	SetData(v *ListRateLimitRulesOutput) *ListRateLimitRulesResult
	GetData() *ListRateLimitRulesOutput
	SetRequestId(v string) *ListRateLimitRulesResult
	GetRequestId() *string
}

type ListRateLimitRulesResult struct {
	// A value of `SUCCESS` indicates a successful request. If the request fails, an error code is returned.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Detailed information about the rate limit rules.
	//
	// example:
	//
	// {}
	Data *ListRateLimitRulesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request ID, used for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRateLimitRulesResult) String() string {
	return dara.Prettify(s)
}

func (s ListRateLimitRulesResult) GoString() string {
	return s.String()
}

func (s *ListRateLimitRulesResult) GetCode() *string {
	return s.Code
}

func (s *ListRateLimitRulesResult) GetData() *ListRateLimitRulesOutput {
	return s.Data
}

func (s *ListRateLimitRulesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRateLimitRulesResult) SetCode(v string) *ListRateLimitRulesResult {
	s.Code = &v
	return s
}

func (s *ListRateLimitRulesResult) SetData(v *ListRateLimitRulesOutput) *ListRateLimitRulesResult {
	s.Data = v
	return s
}

func (s *ListRateLimitRulesResult) SetRequestId(v string) *ListRateLimitRulesResult {
	s.RequestId = &v
	return s
}

func (s *ListRateLimitRulesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
