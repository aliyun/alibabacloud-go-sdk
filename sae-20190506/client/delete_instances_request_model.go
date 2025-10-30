// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteInstancesRequest
	GetAppId() *string
	SetInstanceIds(v string) *DeleteInstancesRequest
	GetInstanceIds() *string
}

type DeleteInstancesRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-aaed579c-1f8a-431e-8705-97d18e91c7b4******
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DeleteInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DeleteInstancesRequest) SetAppId(v string) *DeleteInstancesRequest {
	s.AppId = &v
	return s
}

func (s *DeleteInstancesRequest) SetInstanceIds(v string) *DeleteInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DeleteInstancesRequest) Validate() error {
	return dara.Validate(s)
}
