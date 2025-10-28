// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteLogPathRequest
	GetAppId() *string
	SetPath(v string) *DeleteLogPathRequest
	GetPath() *string
}

type DeleteLogPathRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413**********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The absolute path of the log directory that you want to remove. The value must start and end with a forward slash (`/`) and must contain `/log` or `/logs`. The following directories are the default log directories in Enterprise Distributed Application Service (EDAS):
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
	// example:
	//
	// /temp/log/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteLogPathRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogPathRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogPathRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteLogPathRequest) GetPath() *string {
	return s.Path
}

func (s *DeleteLogPathRequest) SetAppId(v string) *DeleteLogPathRequest {
	s.AppId = &v
	return s
}

func (s *DeleteLogPathRequest) SetPath(v string) *DeleteLogPathRequest {
	s.Path = &v
	return s
}

func (s *DeleteLogPathRequest) Validate() error {
	return dara.Validate(s)
}
