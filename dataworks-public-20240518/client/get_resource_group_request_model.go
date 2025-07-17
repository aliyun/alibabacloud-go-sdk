// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetResourceGroupRequest
	GetId() *string
}

type GetResourceGroupRequest struct {
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) GetId() *string {
	return s.Id
}

func (s *GetResourceGroupRequest) SetId(v string) *GetResourceGroupRequest {
	s.Id = &v
	return s
}

func (s *GetResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
