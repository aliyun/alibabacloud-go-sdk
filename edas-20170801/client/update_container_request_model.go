// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateContainerRequest
	GetAppId() *string
	SetBuildPackId(v int32) *UpdateContainerRequest
	GetBuildPackId() *int32
}

type UpdateContainerRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// e83acea6-****-47e1-96ae-c0e953772cdc
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The build package number of EDAS Container. You can obtain the build package number in the Build package number column in the EDAS Container release notes table. For more information, see [Release notes for EDAS Container](https://help.aliyun.com/document_detail/92614.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 59
	BuildPackId *int32 `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
}

func (s UpdateContainerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateContainerRequest) GetBuildPackId() *int32 {
	return s.BuildPackId
}

func (s *UpdateContainerRequest) SetAppId(v string) *UpdateContainerRequest {
	s.AppId = &v
	return s
}

func (s *UpdateContainerRequest) SetBuildPackId(v int32) *UpdateContainerRequest {
	s.BuildPackId = &v
	return s
}

func (s *UpdateContainerRequest) Validate() error {
	return dara.Validate(s)
}
