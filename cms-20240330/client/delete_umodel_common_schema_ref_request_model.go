// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelCommonSchemaRefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v string) *DeleteUmodelCommonSchemaRefRequest
	GetGroup() *string
}

type DeleteUmodelCommonSchemaRefRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// apm-common
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
}

func (s DeleteUmodelCommonSchemaRefRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelCommonSchemaRefRequest) GoString() string {
	return s.String()
}

func (s *DeleteUmodelCommonSchemaRefRequest) GetGroup() *string {
	return s.Group
}

func (s *DeleteUmodelCommonSchemaRefRequest) SetGroup(v string) *DeleteUmodelCommonSchemaRefRequest {
	s.Group = &v
	return s
}

func (s *DeleteUmodelCommonSchemaRefRequest) Validate() error {
	return dara.Validate(s)
}
