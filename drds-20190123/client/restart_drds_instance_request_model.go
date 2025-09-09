// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDrdsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *RestartDrdsInstanceRequest
	GetDrdsInstanceId() *string
}

type RestartDrdsInstanceRequest struct {
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RestartDrdsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RestartDrdsInstanceRequest) SetDrdsInstanceId(v string) *RestartDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RestartDrdsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
