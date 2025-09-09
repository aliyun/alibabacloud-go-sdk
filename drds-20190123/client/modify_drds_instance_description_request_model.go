// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDrdsInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDrdsInstanceDescriptionRequest
	GetDescription() *string
	SetDrdsInstanceId(v string) *ModifyDrdsInstanceDescriptionRequest
	GetDrdsInstanceId() *string
}

type ModifyDrdsInstanceDescriptionRequest struct {
	// The description of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDrdsInstanceDescriptionRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDescription(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDrdsInstanceId(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
