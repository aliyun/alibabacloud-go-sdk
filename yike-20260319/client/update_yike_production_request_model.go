// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateYikeProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *UpdateYikeProductionRequest
	GetProductionId() *string
	SetTitle(v string) *UpdateYikeProductionRequest
	GetTitle() *string
}

type UpdateYikeProductionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd_12334**
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateYikeProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateYikeProductionRequest) GoString() string {
	return s.String()
}

func (s *UpdateYikeProductionRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *UpdateYikeProductionRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateYikeProductionRequest) SetProductionId(v string) *UpdateYikeProductionRequest {
	s.ProductionId = &v
	return s
}

func (s *UpdateYikeProductionRequest) SetTitle(v string) *UpdateYikeProductionRequest {
	s.Title = &v
	return s
}

func (s *UpdateYikeProductionRequest) Validate() error {
	return dara.Validate(s)
}
