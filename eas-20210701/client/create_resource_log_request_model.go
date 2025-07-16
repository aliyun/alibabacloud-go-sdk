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
	// The Logstore of Log Service. For more information about how to query a Logstore, see [ListLogStores](https://help.aliyun.com/document_detail/426970.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// access_log
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project that is associated with the resource group. For more information about how to query the project, see [ListProject](https://help.aliyun.com/document_detail/74955.html).
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
