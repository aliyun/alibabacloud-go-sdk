// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBruteForceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableBruteForceRecordResponseBody
	GetRequestId() *string
}

type DisableBruteForceRecordResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5EFF53F7-9B2A-58B5-AD06-6B07ACE17133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableBruteForceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableBruteForceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DisableBruteForceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableBruteForceRecordResponseBody) SetRequestId(v string) *DisableBruteForceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableBruteForceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
