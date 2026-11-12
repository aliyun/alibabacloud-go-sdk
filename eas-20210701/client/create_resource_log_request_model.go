// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogStore(v string) *CreateResourceLogRequest
	GetLogStore() *string
	SetProjectName(v string) *CreateResourceLogRequest
	GetProjectName() *string
}

type CreateResourceLogRequest struct {
	// The Simple Log Service (SLS) Logstore. For more information about how to query Logstores, see [ListLogStores](https://help.aliyun.com/document_detail/426970.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// access_log
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the Simple Log Service (SLS) project for the resource group. The project must be created in Simple Log Service in advance. This operation does not support automatic creation of the project. You can call the ListProject operation to query existing projects, or call the CreateProject operation to create a project first.
	//
	// This parameter is required.
	//
	// example:
	//
	// eas-r-asdasdasd-sls
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateResourceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceLogRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceLogRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *CreateResourceLogRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateResourceLogRequest) SetLogStore(v string) *CreateResourceLogRequest {
	s.LogStore = &v
	return s
}

func (s *CreateResourceLogRequest) SetProjectName(v string) *CreateResourceLogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateResourceLogRequest) Validate() error {
	return dara.Validate(s)
}
