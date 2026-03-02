// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTTSVerifyLinkResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendTTSVerifyLinkResult
	GetRequestId() *string
	SetResult(v bool) *SendTTSVerifyLinkResult
	GetResult() *bool
}

type SendTTSVerifyLinkResult struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SendTTSVerifyLinkResult) String() string {
	return dara.Prettify(s)
}

func (s SendTTSVerifyLinkResult) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkResult) GetRequestId() *string {
	return s.RequestId
}

func (s *SendTTSVerifyLinkResult) GetResult() *bool {
	return s.Result
}

func (s *SendTTSVerifyLinkResult) SetRequestId(v string) *SendTTSVerifyLinkResult {
	s.RequestId = &v
	return s
}

func (s *SendTTSVerifyLinkResult) SetResult(v bool) *SendTTSVerifyLinkResult {
	s.Result = &v
	return s
}

func (s *SendTTSVerifyLinkResult) Validate() error {
	return dara.Validate(s)
}
