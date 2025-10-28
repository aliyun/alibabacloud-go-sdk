// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeployGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListDeployGroupRequest
	GetAppId() *string
}

type ListDeployGroupRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListDeployGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDeployGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListDeployGroupRequest) SetAppId(v string) *ListDeployGroupRequest {
	s.AppId = &v
	return s
}

func (s *ListDeployGroupRequest) Validate() error {
	return dara.Validate(s)
}
