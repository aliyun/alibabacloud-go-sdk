// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLogPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddLogPathRequest
	GetAppId() *string
	SetPath(v string) *AddLogPathRequest
	GetPath() *string
}

type AddLogPathRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413**********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The absolute path of the log directory that you want to add. The value must start and end with a forward slash (`/`) and must contain `/log` or `/logs`. The following directories are the default log directories in Enterprise Distributed Application Service (EDAS):
	//
	// 	- /home/admin/edas-container/logs/
	//
	// 	- /home/admin/taobao-tomcat-7.0.59/logs/
	//
	// 	- /home/admin/taobao-tomcat-production-7.0.59.3/logs/
	//
	// 	- /home/admin/taobao-tomcat-production-7.0.70/logs/
	//
	// 	- /home/admin/edas-agent/logs/
	//
	// This parameter is required.
	//
	// example:
	//
	// /temp/log/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AddLogPathRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLogPathRequest) GoString() string {
	return s.String()
}

func (s *AddLogPathRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddLogPathRequest) GetPath() *string {
	return s.Path
}

func (s *AddLogPathRequest) SetAppId(v string) *AddLogPathRequest {
	s.AppId = &v
	return s
}

func (s *AddLogPathRequest) SetPath(v string) *AddLogPathRequest {
	s.Path = &v
	return s
}

func (s *AddLogPathRequest) Validate() error {
	return dara.Validate(s)
}
