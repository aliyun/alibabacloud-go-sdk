// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksBloodRelationshipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorksId(v string) *QueryWorksBloodRelationshipRequest
	GetWorksId() *string
}

type QueryWorksBloodRelationshipRequest struct {
	// The ID of the data work.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcd****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryWorksBloodRelationshipRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksBloodRelationshipRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryWorksBloodRelationshipRequest) SetWorksId(v string) *QueryWorksBloodRelationshipRequest {
	s.WorksId = &v
	return s
}

func (s *QueryWorksBloodRelationshipRequest) Validate() error {
	return dara.Validate(s)
}
