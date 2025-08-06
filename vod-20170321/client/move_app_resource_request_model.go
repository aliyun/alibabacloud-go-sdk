// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveAppResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v string) *MoveAppResourceRequest
	GetResourceIds() *string
	SetResourceType(v string) *MoveAppResourceRequest
	GetResourceType() *string
	SetTargetAppId(v string) *MoveAppResourceRequest
	GetTargetAppId() *string
}

type MoveAppResourceRequest struct {
	// The resource ID. You can specify a maximum of 20 IDs at a time. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 9afb4****06de180880e,f7bba****caa546cfe2ba
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **video**: video files.
	//
	// 	- **image**: image files.
	//
	// 	- **attached**: auxiliary media assets.
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the application to which resources are migrated. Default value: **app-1000000**. For more information, see [Use the multi-application service](https://help.aliyun.com/document_detail/113600.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	TargetAppId *string `json:"TargetAppId,omitempty" xml:"TargetAppId,omitempty"`
}

func (s MoveAppResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveAppResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveAppResourceRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *MoveAppResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveAppResourceRequest) GetTargetAppId() *string {
	return s.TargetAppId
}

func (s *MoveAppResourceRequest) SetResourceIds(v string) *MoveAppResourceRequest {
	s.ResourceIds = &v
	return s
}

func (s *MoveAppResourceRequest) SetResourceType(v string) *MoveAppResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveAppResourceRequest) SetTargetAppId(v string) *MoveAppResourceRequest {
	s.TargetAppId = &v
	return s
}

func (s *MoveAppResourceRequest) Validate() error {
	return dara.Validate(s)
}
