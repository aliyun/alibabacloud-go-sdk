// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteResourceRequest
	GetId() *string
	SetProjectId(v int64) *DeleteResourceRequest
	GetProjectId() *int64
}

type DeleteResourceRequest struct {
	// The unique identifier of the resource file.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID. You can use this parameter to specify the DataWorks workspace on which you want to perform the API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) GetId() *string {
	return s.Id
}

func (s *DeleteResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteResourceRequest) SetId(v string) *DeleteResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteResourceRequest) SetProjectId(v int64) *DeleteResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteResourceRequest) Validate() error {
	return dara.Validate(s)
}
