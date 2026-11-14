// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdList(v []*string) *UnbindAppRequest
	GetAppIdList() []*string
	SetImageId(v string) *UnbindAppRequest
	GetImageId() *string
	SetVersionId(v string) *UnbindAppRequest
	GetVersionId() *string
}

type UnbindAppRequest struct {
	// The list of application IDs to disassociate from the specified image version. Verify that each application is associated with the selected image version.
	//
	// This parameter is required.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The ID of the image from which the applications are to be disassociated. The image must be a custom image or a community image manageable by the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the image version from which the applications are to be disassociated. If this parameter is omitted, the current active version of the image is used. If specified, the version must belong to the image specified by ImageId.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UnbindAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppRequest) GoString() string {
	return s.String()
}

func (s *UnbindAppRequest) GetAppIdList() []*string {
	return s.AppIdList
}

func (s *UnbindAppRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UnbindAppRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UnbindAppRequest) SetAppIdList(v []*string) *UnbindAppRequest {
	s.AppIdList = v
	return s
}

func (s *UnbindAppRequest) SetImageId(v string) *UnbindAppRequest {
	s.ImageId = &v
	return s
}

func (s *UnbindAppRequest) SetVersionId(v string) *UnbindAppRequest {
	s.VersionId = &v
	return s
}

func (s *UnbindAppRequest) Validate() error {
	return dara.Validate(s)
}
