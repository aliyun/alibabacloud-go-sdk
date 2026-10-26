// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySQLCollectorPolicyResponseBody
	GetRequestId() *string
}

type ModifySQLCollectorPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySQLCollectorPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySQLCollectorPolicyResponseBody) SetRequestId(v string) *ModifySQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySQLCollectorPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
