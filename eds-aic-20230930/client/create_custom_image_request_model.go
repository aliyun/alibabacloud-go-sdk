// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCustomImageRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCustomImageRequest
	GetDescription() *string
	SetImageName(v string) *CreateCustomImageRequest
	GetImageName() *string
	SetInstanceId(v string) *CreateCustomImageRequest
	GetInstanceId() *string
}

type CreateCustomImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. By default, this parameter is left empty. The token cannot exceed 64 characters in length.
	//
	// example:
	//
	// 20393E53-8FF1-524C-B494-B478A5369733
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the custom image.
	//
	// example:
	//
	// create for cc5g group auth rules test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom image name
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the cloud phone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCustomImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomImageRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateCustomImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomImageRequest) SetClientToken(v string) *CreateCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomImageRequest) SetDescription(v string) *CreateCustomImageRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomImageRequest) SetImageName(v string) *CreateCustomImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateCustomImageRequest) SetInstanceId(v string) *CreateCustomImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomImageRequest) Validate() error {
	return dara.Validate(s)
}
