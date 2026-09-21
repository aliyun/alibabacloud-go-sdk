// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudPhoneNodeSystemImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpdateCloudPhoneNodeSystemImageRequest
	GetImageId() *string
	SetNodeIds(v []*string) *UpdateCloudPhoneNodeSystemImageRequest
	GetNodeIds() []*string
}

type UpdateCloudPhoneNodeSystemImageRequest struct {
	// The image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The list of cloud phone normal matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s UpdateCloudPhoneNodeSystemImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPhoneNodeSystemImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudPhoneNodeSystemImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateCloudPhoneNodeSystemImageRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *UpdateCloudPhoneNodeSystemImageRequest) SetImageId(v string) *UpdateCloudPhoneNodeSystemImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageRequest) SetNodeIds(v []*string) *UpdateCloudPhoneNodeSystemImageRequest {
	s.NodeIds = v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageRequest) Validate() error {
	return dara.Validate(s)
}
