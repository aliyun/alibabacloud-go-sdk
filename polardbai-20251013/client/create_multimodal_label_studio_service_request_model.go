// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalLabelStudioServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalLabelStudioServiceRequest
	GetDBClusterId() *string
	SetDatasetIds(v []*string) *CreateMultimodalLabelStudioServiceRequest
	GetDatasetIds() []*string
	SetPassword(v string) *CreateMultimodalLabelStudioServiceRequest
	GetPassword() *string
	SetUsername(v string) *CreateMultimodalLabelStudioServiceRequest
	GetUsername() *string
}

type CreateMultimodalLabelStudioServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	DatasetIds []*string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty" type:"Repeated"`
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

func (s CreateMultimodalLabelStudioServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalLabelStudioServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalLabelStudioServiceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalLabelStudioServiceRequest) GetDatasetIds() []*string {
	return s.DatasetIds
}

func (s *CreateMultimodalLabelStudioServiceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateMultimodalLabelStudioServiceRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateMultimodalLabelStudioServiceRequest) SetDBClusterId(v string) *CreateMultimodalLabelStudioServiceRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceRequest) SetDatasetIds(v []*string) *CreateMultimodalLabelStudioServiceRequest {
	s.DatasetIds = v
	return s
}

func (s *CreateMultimodalLabelStudioServiceRequest) SetPassword(v string) *CreateMultimodalLabelStudioServiceRequest {
	s.Password = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceRequest) SetUsername(v string) *CreateMultimodalLabelStudioServiceRequest {
	s.Username = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceRequest) Validate() error {
	return dara.Validate(s)
}
