// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *RemoveDrdsInstanceRequest
	GetDrdsInstanceId() *string
}

type RemoveDrdsInstanceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RemoveDrdsInstanceRequest) SetDrdsInstanceId(v string) *RemoveDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveDrdsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
