// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelationType(v string) *GetWorkitemRelationsRequest
	GetRelationType() *string
}

type GetWorkitemRelationsRequest struct {
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s GetWorkitemRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemRelationsRequest) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *GetWorkitemRelationsRequest) SetRelationType(v string) *GetWorkitemRelationsRequest {
	s.RelationType = &v
	return s
}

func (s *GetWorkitemRelationsRequest) Validate() error {
	return dara.Validate(s)
}
