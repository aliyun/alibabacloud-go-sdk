// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateModel(v string) *UpdateDataSourceRequest
	GetUpdateModel() *string
}

type UpdateDataSourceRequest struct {
	// This parameter is required.
	UpdateModel *string `json:"UpdateModel,omitempty" xml:"UpdateModel,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) GetUpdateModel() *string {
	return s.UpdateModel
}

func (s *UpdateDataSourceRequest) SetUpdateModel(v string) *UpdateDataSourceRequest {
	s.UpdateModel = &v
	return s
}

func (s *UpdateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
