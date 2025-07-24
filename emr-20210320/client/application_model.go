// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplication interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *Application
	GetApplicationName() *string
}

type Application struct {
	// 应用名称。从EMR控制台集群创建页面可查看到指定发行版的应用名称列表。
	//
	// This parameter is required.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
}

func (s Application) String() string {
	return dara.Prettify(s)
}

func (s Application) GoString() string {
	return s.String()
}

func (s *Application) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *Application) SetApplicationName(v string) *Application {
	s.ApplicationName = &v
	return s
}

func (s *Application) Validate() error {
	return dara.Validate(s)
}
