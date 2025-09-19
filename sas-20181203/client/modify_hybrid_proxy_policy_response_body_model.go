// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridProxyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessgae(v string) *ModifyHybridProxyPolicyResponseBody
	GetMessgae() *string
	SetRequestId(v string) *ModifyHybridProxyPolicyResponseBody
	GetRequestId() *string
}

type ModifyHybridProxyPolicyResponseBody struct {
	// The message of the request.
	//
	// example:
	//
	// clusterName data not exist
	Messgae *string `json:"Messgae,omitempty" xml:"Messgae,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridProxyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridProxyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridProxyPolicyResponseBody) GetMessgae() *string {
	return s.Messgae
}

func (s *ModifyHybridProxyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridProxyPolicyResponseBody) SetMessgae(v string) *ModifyHybridProxyPolicyResponseBody {
	s.Messgae = &v
	return s
}

func (s *ModifyHybridProxyPolicyResponseBody) SetRequestId(v string) *ModifyHybridProxyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridProxyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
