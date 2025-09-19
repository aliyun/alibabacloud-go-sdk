// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridProxyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *ModifyHybridProxyPolicyRequest
	GetClusterName() *string
	SetPolicyInfo(v string) *ModifyHybridProxyPolicyRequest
	GetPolicyInfo() *string
}

type ModifyHybridProxyPolicyRequest struct {
	// The name of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas-proxy
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The policy of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "policyType": "limitFrequency",
	//
	//         "info":
	//
	//         [
	//
	//             {
	//
	//                 "type": "file",
	//
	//                 "config": "10"
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     {
	//
	//         "policyType": "limitBandWidth",
	//
	//         "info":
	//
	//         [
	//
	//             {
	//
	//                 "type": "net"
	//
	//             },
	//
	//             {
	//
	//                 "type": "process"
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// ]
	PolicyInfo *string `json:"PolicyInfo,omitempty" xml:"PolicyInfo,omitempty"`
}

func (s ModifyHybridProxyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridProxyPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridProxyPolicyRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyHybridProxyPolicyRequest) GetPolicyInfo() *string {
	return s.PolicyInfo
}

func (s *ModifyHybridProxyPolicyRequest) SetClusterName(v string) *ModifyHybridProxyPolicyRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyHybridProxyPolicyRequest) SetPolicyInfo(v string) *ModifyHybridProxyPolicyRequest {
	s.PolicyInfo = &v
	return s
}

func (s *ModifyHybridProxyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
