// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileProtectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFileProtectRuleResponseBody
	GetRequestId() *string
}

type DeleteFileProtectRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7FD1C1DC-AA67-510A-A022-5D23310C3658
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileProtectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileProtectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileProtectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileProtectRuleResponseBody) SetRequestId(v string) *DeleteFileProtectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileProtectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
