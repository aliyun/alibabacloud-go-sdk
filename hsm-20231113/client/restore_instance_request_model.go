// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *RestoreInstanceRequest
	GetImageId() *string
	SetInstanceId(v string) *RestoreInstanceRequest
	GetInstanceId() *string
}

type RestoreInstanceRequest struct {
	// The ID of the image that you want to use to restore the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// image-eaOGHkRDQgh4****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestoreInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RestoreInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestoreInstanceRequest) SetImageId(v string) *RestoreInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *RestoreInstanceRequest) SetInstanceId(v string) *RestoreInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestoreInstanceRequest) Validate() error {
	return dara.Validate(s)
}
