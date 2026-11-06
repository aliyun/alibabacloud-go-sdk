// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakebaseS3AccountPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLakebaseS3AccountPolicyResponseBody
	GetRequestId() *string
}

type ModifyLakebaseS3AccountPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLakebaseS3AccountPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakebaseS3AccountPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLakebaseS3AccountPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLakebaseS3AccountPolicyResponseBody) SetRequestId(v string) *ModifyLakebaseS3AccountPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
