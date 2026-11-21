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
	SetDryRun(v bool) *DeleteUmodelDataRequest
	GetDryRun() *bool
	SetKind(v string) *DeleteUmodelDataRequest
	GetKind() *string
	SetName(v string) *DeleteUmodelDataRequest
	GetName() *string
}

type DeleteUmodelDataRequest struct {
	// The name of the specific Umodel data. If you do not specify this parameter, all Umodel data is included.
	//
	// example:
	//
	// apm
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// Specifies whether to perform a dry run of the request. Valid values:
	//
	// - true: sends a check request and does not perform actual changes.
	//
	// - false: sends a normal request.
	//
	// Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The kind of the specific Umodel data. If you do not specify this parameter, all kinds are included.
	//
	// example:
	//
	// metric_set
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the specific Umodel data. If you do not specify this parameter, all Umodel data is included.
	//
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

func (s *DeleteUmodelDataRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *DeleteUmodelDataRequest) SetDryRun(v bool) *DeleteUmodelDataRequest {
	s.DryRun = &v
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
