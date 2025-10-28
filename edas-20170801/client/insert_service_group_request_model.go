// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *InsertServiceGroupRequest
	GetGroupName() *string
}

type InsertServiceGroupRequest struct {
	// The name of the service group that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// edas-test-group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s InsertServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *InsertServiceGroupRequest) SetGroupName(v string) *InsertServiceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *InsertServiceGroupRequest) Validate() error {
	return dara.Validate(s)
}
