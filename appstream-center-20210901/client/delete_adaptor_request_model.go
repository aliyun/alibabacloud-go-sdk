// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdaptorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptorId(v string) *DeleteAdaptorRequest
	GetAdaptorId() *string
	SetAppId(v string) *DeleteAdaptorRequest
	GetAppId() *string
	SetAppVersionId(v string) *DeleteAdaptorRequest
	GetAppVersionId() *string
	SetImageId(v string) *DeleteAdaptorRequest
	GetImageId() *string
	SetVersionId(v string) *DeleteAdaptorRequest
	GetVersionId() *string
}

type DeleteAdaptorRequest struct {
	// The ID of the application adaptation to be processed. **Required**. It must point to the same application adaptation as AppId and AppVersionId.
	//
	// example:
	//
	// adt-bp1example****
	AdaptorId *string `json:"AdaptorId,omitempty" xml:"AdaptorId,omitempty"`
	// The ID of the application. **Required**. It must match the application corresponding to AdaptorId.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application version. **Required**. It must match the application version corresponding to AdaptorId. This parameter is not the image version ID.
	//
	// example:
	//
	// apv-bp1example****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The ID of the image to be dissociated. **Must be provided together with VersionId**. If both parameters are not specified at the same time, the adaptation deletion process starts.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the image version to be dissociated. It must be provided together with ImageId and belong to the specified image.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteAdaptorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdaptorRequest) GoString() string {
	return s.String()
}

func (s *DeleteAdaptorRequest) GetAdaptorId() *string {
	return s.AdaptorId
}

func (s *DeleteAdaptorRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAdaptorRequest) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *DeleteAdaptorRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteAdaptorRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeleteAdaptorRequest) SetAdaptorId(v string) *DeleteAdaptorRequest {
	s.AdaptorId = &v
	return s
}

func (s *DeleteAdaptorRequest) SetAppId(v string) *DeleteAdaptorRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAdaptorRequest) SetAppVersionId(v string) *DeleteAdaptorRequest {
	s.AppVersionId = &v
	return s
}

func (s *DeleteAdaptorRequest) SetImageId(v string) *DeleteAdaptorRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteAdaptorRequest) SetVersionId(v string) *DeleteAdaptorRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteAdaptorRequest) Validate() error {
	return dara.Validate(s)
}
