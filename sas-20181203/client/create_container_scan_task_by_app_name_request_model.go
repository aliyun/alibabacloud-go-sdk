// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContainerScanTaskByAppNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppNames(v string) *CreateContainerScanTaskByAppNameRequest
	GetAppNames() *string
	SetClusterId(v string) *CreateContainerScanTaskByAppNameRequest
	GetClusterId() *string
	SetLang(v string) *CreateContainerScanTaskByAppNameRequest
	GetLang() *string
}

type CreateContainerScanTaskByAppNameRequest struct {
	// The name of the container application.
	//
	// example:
	//
	// app-centos-01
	AppNames *string `json:"AppNames,omitempty" xml:"AppNames,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cfb7a55a81f7246b5ac18845ea79a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The language of the content within the request and response.
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

func (s CreateContainerScanTaskByAppNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskByAppNameRequest) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskByAppNameRequest) GetAppNames() *string {
	return s.AppNames
}

func (s *CreateContainerScanTaskByAppNameRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateContainerScanTaskByAppNameRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateContainerScanTaskByAppNameRequest) SetAppNames(v string) *CreateContainerScanTaskByAppNameRequest {
	s.AppNames = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameRequest) SetClusterId(v string) *CreateContainerScanTaskByAppNameRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameRequest) SetLang(v string) *CreateContainerScanTaskByAppNameRequest {
	s.Lang = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameRequest) Validate() error {
	return dara.Validate(s)
}
