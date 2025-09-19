// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileProtectRuleResponseBody
	GetRequestId() *string
}

type UpdateFileProtectRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 339227F4-C9E1-526F-8347-1099C11F65FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileProtectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileProtectRuleResponseBody) SetRequestId(v string) *UpdateFileProtectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileProtectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
