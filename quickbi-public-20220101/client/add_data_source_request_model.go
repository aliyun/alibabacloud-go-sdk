// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddModel(v string) *AddDataSourceRequest
	GetAddModel() *string
}

type AddDataSourceRequest struct {
	// This parameter is required.
	AddModel *string `json:"AddModel,omitempty" xml:"AddModel,omitempty"`
}

func (s AddDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceRequest) GetAddModel() *string {
	return s.AddModel
}

func (s *AddDataSourceRequest) SetAddModel(v string) *AddDataSourceRequest {
	s.AddModel = &v
	return s
}

func (s *AddDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
