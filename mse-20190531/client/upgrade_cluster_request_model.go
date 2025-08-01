// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpgradeClusterRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *UpgradeClusterRequest
	GetInstanceId() *string
	SetRequestPars(v string) *UpgradeClusterRequest
	GetRequestPars() *string
	SetUpgradeVersion(v string) *UpgradeClusterRequest
	GetUpgradeVersion() *string
}

type UpgradeClusterRequest struct {
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
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-nif1w51wi0c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The destination version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.2.1
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s UpgradeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpgradeClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeClusterRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *UpgradeClusterRequest) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *UpgradeClusterRequest) SetAcceptLanguage(v string) *UpgradeClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpgradeClusterRequest) SetInstanceId(v string) *UpgradeClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeClusterRequest) SetRequestPars(v string) *UpgradeClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *UpgradeClusterRequest) SetUpgradeVersion(v string) *UpgradeClusterRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *UpgradeClusterRequest) Validate() error {
	return dara.Validate(s)
}
