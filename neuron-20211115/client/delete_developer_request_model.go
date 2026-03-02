// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *DeleteDeveloperRequest
	GetEnterpriseId() *int64
}

type DeleteDeveloperRequest struct {
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
}

func (s DeleteDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeveloperRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeveloperRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *DeleteDeveloperRequest) SetEnterpriseId(v int64) *DeleteDeveloperRequest {
	s.EnterpriseId = &v
	return s
}

func (s *DeleteDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
