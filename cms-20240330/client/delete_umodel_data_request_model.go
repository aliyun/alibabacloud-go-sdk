// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteUmodelDataRequest
	GetDomain() *string
	SetKind(v string) *DeleteUmodelDataRequest
	GetKind() *string
	SetName(v string) *DeleteUmodelDataRequest
	GetName() *string
}

type DeleteUmodelDataRequest struct {
	// example:
	//
	// apm
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// metric_set
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeleteUmodelDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteUmodelDataRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteUmodelDataRequest) GetKind() *string {
	return s.Kind
}

func (s *DeleteUmodelDataRequest) GetName() *string {
	return s.Name
}

func (s *DeleteUmodelDataRequest) SetDomain(v string) *DeleteUmodelDataRequest {
	s.Domain = &v
	return s
}

func (s *DeleteUmodelDataRequest) SetKind(v string) *DeleteUmodelDataRequest {
	s.Kind = &v
	return s
}

func (s *DeleteUmodelDataRequest) SetName(v string) *DeleteUmodelDataRequest {
	s.Name = &v
	return s
}

func (s *DeleteUmodelDataRequest) Validate() error {
	return dara.Validate(s)
}
