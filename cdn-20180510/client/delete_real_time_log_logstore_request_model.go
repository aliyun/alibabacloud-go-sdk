// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRealTimeLogLogstoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DeleteRealTimeLogLogstoreRequest
	GetLogstore() *string
	SetProject(v string) *DeleteRealTimeLogLogstoreRequest
	GetProject() *string
	SetRegion(v string) *DeleteRealTimeLogLogstoreRequest
	GetRegion() *string
}

type DeleteRealTimeLogLogstoreRequest struct {
	// The name of the Logstore to which log entries are delivered.
	//
	// This parameter is required.
	//
	// example:
	//
	// LogstoreName
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProjectName
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed. For more information, see [Regions that support real-time log delivery](https://help.aliyun.com/document_detail/144883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteRealTimeLogLogstoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRealTimeLogLogstoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteRealTimeLogLogstoreRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DeleteRealTimeLogLogstoreRequest) GetProject() *string {
	return s.Project
}

func (s *DeleteRealTimeLogLogstoreRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteRealTimeLogLogstoreRequest) SetLogstore(v string) *DeleteRealTimeLogLogstoreRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteRealTimeLogLogstoreRequest) SetProject(v string) *DeleteRealTimeLogLogstoreRequest {
	s.Project = &v
	return s
}

func (s *DeleteRealTimeLogLogstoreRequest) SetRegion(v string) *DeleteRealTimeLogLogstoreRequest {
	s.Region = &v
	return s
}

func (s *DeleteRealTimeLogLogstoreRequest) Validate() error {
	return dara.Validate(s)
}
