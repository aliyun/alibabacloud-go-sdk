// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHiveAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyHiveAttributeRequest
	GetDescription() *string
	SetHiveId(v string) *ModifyHiveAttributeRequest
	GetHiveId() *string
	SetName(v string) *ModifyHiveAttributeRequest
	GetName() *string
}

type ModifyHiveAttributeRequest struct {
	// example:
	//
	// gb-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID
	//
	// This parameter is required.
	//
	// example:
	//
	// hive-3b506f0868a7451ba15e0e890706033a
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// example:
	//
	// yy-test2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyHiveAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHiveAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHiveAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyHiveAttributeRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *ModifyHiveAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyHiveAttributeRequest) SetDescription(v string) *ModifyHiveAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyHiveAttributeRequest) SetHiveId(v string) *ModifyHiveAttributeRequest {
	s.HiveId = &v
	return s
}

func (s *ModifyHiveAttributeRequest) SetName(v string) *ModifyHiveAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyHiveAttributeRequest) Validate() error {
	return dara.Validate(s)
}
