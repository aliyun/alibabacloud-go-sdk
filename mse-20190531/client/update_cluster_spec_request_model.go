// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateClusterSpecRequest
	GetAcceptLanguage() *string
	SetAutoPay(v bool) *UpdateClusterSpecRequest
	GetAutoPay() *bool
	SetClusterId(v string) *UpdateClusterSpecRequest
	GetClusterId() *string
	SetClusterSpecification(v string) *UpdateClusterSpecRequest
	GetClusterSpecification() *string
	SetInstanceCount(v int32) *UpdateClusterSpecRequest
	GetInstanceCount() *int32
	SetInstanceId(v string) *UpdateClusterSpecRequest
	GetInstanceId() *string
	SetMseVersion(v string) *UpdateClusterSpecRequest
	GetMseVersion() *string
	SetPubNetworkFlow(v int32) *UpdateClusterSpecRequest
	GetPubNetworkFlow() *int32
}

type UpdateClusterSpecRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AutoPay        *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The destination engine specifications.
	//
	// example:
	//
	// MSE_SC_2_4_200_c
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	// The number of destination nodes.
	//
	// example:
	//
	// 3
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The MSE version.
	//
	// example:
	//
	// mse_pro
	MseVersion     *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	PubNetworkFlow *int32  `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
}

func (s UpdateClusterSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterSpecRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateClusterSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateClusterSpecRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterSpecRequest) GetClusterSpecification() *string {
	return s.ClusterSpecification
}

func (s *UpdateClusterSpecRequest) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *UpdateClusterSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateClusterSpecRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *UpdateClusterSpecRequest) GetPubNetworkFlow() *int32 {
	return s.PubNetworkFlow
}

func (s *UpdateClusterSpecRequest) SetAcceptLanguage(v string) *UpdateClusterSpecRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetAutoPay(v bool) *UpdateClusterSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetClusterId(v string) *UpdateClusterSpecRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetClusterSpecification(v string) *UpdateClusterSpecRequest {
	s.ClusterSpecification = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetInstanceCount(v int32) *UpdateClusterSpecRequest {
	s.InstanceCount = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetInstanceId(v string) *UpdateClusterSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetMseVersion(v string) *UpdateClusterSpecRequest {
	s.MseVersion = &v
	return s
}

func (s *UpdateClusterSpecRequest) SetPubNetworkFlow(v int32) *UpdateClusterSpecRequest {
	s.PubNetworkFlow = &v
	return s
}

func (s *UpdateClusterSpecRequest) Validate() error {
	return dara.Validate(s)
}
