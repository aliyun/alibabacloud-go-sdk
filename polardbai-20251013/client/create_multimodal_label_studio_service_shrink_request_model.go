// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalLabelStudioServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalLabelStudioServiceShrinkRequest
	GetDBClusterId() *string
	SetDatasetIdsShrink(v string) *CreateMultimodalLabelStudioServiceShrinkRequest
	GetDatasetIdsShrink() *string
	SetPassword(v string) *CreateMultimodalLabelStudioServiceShrinkRequest
	GetPassword() *string
	SetUsername(v string) *CreateMultimodalLabelStudioServiceShrinkRequest
	GetUsername() *string
}

type CreateMultimodalLabelStudioServiceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	DatasetIdsShrink *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-admin@db4ai.com
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateMultimodalLabelStudioServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalLabelStudioServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) GetDatasetIdsShrink() *string {
	return s.DatasetIdsShrink
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) SetDBClusterId(v string) *CreateMultimodalLabelStudioServiceShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) SetDatasetIdsShrink(v string) *CreateMultimodalLabelStudioServiceShrinkRequest {
	s.DatasetIdsShrink = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) SetPassword(v string) *CreateMultimodalLabelStudioServiceShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) SetUsername(v string) *CreateMultimodalLabelStudioServiceShrinkRequest {
	s.Username = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
