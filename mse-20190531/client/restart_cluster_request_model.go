// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *RestartClusterRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *RestartClusterRequest
	GetClusterId() *string
	SetInstanceId(v string) *RestartClusterRequest
	GetInstanceId() *string
	SetPodNameList(v string) *RestartClusterRequest
	GetPodNameList() *string
	SetRequestPars(v string) *RestartClusterRequest
	GetRequestPars() *string
}

type RestartClusterRequest struct {
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
	// The ID of the cluster.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-78v1l83****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The names of pods. You can specify the names of multiple pods at a time. Separate multiple pod names with commas (,). Example: mse-a8aba010-1629719288255-reg-center-0-1,mse-a8aba010-1629719288255-reg-center-0-2.
	//
	// The specified pods must belong to the current cluster and be associated with the specified instance. Otherwise, a restart exception occurs.
	//
	// example:
	//
	// mse-a8aba010-1629719288255-reg-center-0-1
	PodNameList *string `json:"PodNameList,omitempty" xml:"PodNameList,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s RestartClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartClusterRequest) GoString() string {
	return s.String()
}

func (s *RestartClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *RestartClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RestartClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartClusterRequest) GetPodNameList() *string {
	return s.PodNameList
}

func (s *RestartClusterRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *RestartClusterRequest) SetAcceptLanguage(v string) *RestartClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *RestartClusterRequest) SetClusterId(v string) *RestartClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RestartClusterRequest) SetInstanceId(v string) *RestartClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartClusterRequest) SetPodNameList(v string) *RestartClusterRequest {
	s.PodNameList = &v
	return s
}

func (s *RestartClusterRequest) SetRequestPars(v string) *RestartClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *RestartClusterRequest) Validate() error {
	return dara.Validate(s)
}
