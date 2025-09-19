// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppNames(v string) *ModifyContainerScanConfigRequest
	GetAppNames() *string
	SetClusterId(v string) *ModifyContainerScanConfigRequest
	GetClusterId() *string
	SetLang(v string) *ModifyContainerScanConfigRequest
	GetLang() *string
}

type ModifyContainerScanConfigRequest struct {
	// The name of the container application.
	//
	// example:
	//
	// [\\"logtail-ds\\",\\"alicloud-monitor-controller\\",\\"storage-snapshot-manager\\"]
	AppNames *string `json:"AppNames,omitempty" xml:"AppNames,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// cfb7a55a81f7246b5ac18845ea79a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ModifyContainerScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerScanConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerScanConfigRequest) GetAppNames() *string {
	return s.AppNames
}

func (s *ModifyContainerScanConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyContainerScanConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyContainerScanConfigRequest) SetAppNames(v string) *ModifyContainerScanConfigRequest {
	s.AppNames = &v
	return s
}

func (s *ModifyContainerScanConfigRequest) SetClusterId(v string) *ModifyContainerScanConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyContainerScanConfigRequest) SetLang(v string) *ModifyContainerScanConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyContainerScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
