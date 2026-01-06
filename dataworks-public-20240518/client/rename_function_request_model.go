// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RenameFunctionRequest
	GetId() *string
	SetName(v string) *RenameFunctionRequest
	GetName() *string
	SetProjectId(v int64) *RenameFunctionRequest
	GetProjectId() *int64
}

type RenameFunctionRequest struct {
	// The unique identifier of the Data Studio UDF.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the Data Studio UDF.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameFunctionRequest) GoString() string {
	return s.String()
}

func (s *RenameFunctionRequest) GetId() *string {
	return s.Id
}

func (s *RenameFunctionRequest) GetName() *string {
	return s.Name
}

func (s *RenameFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameFunctionRequest) SetId(v string) *RenameFunctionRequest {
	s.Id = &v
	return s
}

func (s *RenameFunctionRequest) SetName(v string) *RenameFunctionRequest {
	s.Name = &v
	return s
}

func (s *RenameFunctionRequest) SetProjectId(v int64) *RenameFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *RenameFunctionRequest) Validate() error {
	return dara.Validate(s)
}
