// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmbeddedStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThirdPartAuthFlag(v bool) *UpdateEmbeddedStatusRequest
	GetThirdPartAuthFlag() *bool
	SetWorksId(v string) *UpdateEmbeddedStatusRequest
	GetWorksId() *string
}

type UpdateEmbeddedStatusRequest struct {
	// Whether to enable the embedding feature for the work. Valid values:
	//
	// 	- true: enables embedding.
	//
	// 	- false: disables embedding.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ThirdPartAuthFlag *bool `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The ID of the work.
	//
	// 	- Batch modification is supported. Separate multiple values with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 897ce25e-f993-4abd-af84-d13c5610****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s UpdateEmbeddedStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmbeddedStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmbeddedStatusRequest) GetThirdPartAuthFlag() *bool {
	return s.ThirdPartAuthFlag
}

func (s *UpdateEmbeddedStatusRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *UpdateEmbeddedStatusRequest) SetThirdPartAuthFlag(v bool) *UpdateEmbeddedStatusRequest {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *UpdateEmbeddedStatusRequest) SetWorksId(v string) *UpdateEmbeddedStatusRequest {
	s.WorksId = &v
	return s
}

func (s *UpdateEmbeddedStatusRequest) Validate() error {
	return dara.Validate(s)
}
