// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDingTalkOnlineTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DingTalkOnlineTestResponseBody
	GetRequestId() *string
}

type DingTalkOnlineTestResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2E96F605-1BE3-5954-83A5-AE96C617****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DingTalkOnlineTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DingTalkOnlineTestResponseBody) GoString() string {
	return s.String()
}

func (s *DingTalkOnlineTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DingTalkOnlineTestResponseBody) SetRequestId(v string) *DingTalkOnlineTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DingTalkOnlineTestResponseBody) Validate() error {
	return dara.Validate(s)
}
