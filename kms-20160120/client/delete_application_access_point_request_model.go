// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteApplicationAccessPointRequest
	GetName() *string
}

type DeleteApplicationAccessPointRequest struct {
	// The name of the AAP that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteApplicationAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationAccessPointRequest) GetName() *string {
	return s.Name
}

func (s *DeleteApplicationAccessPointRequest) SetName(v string) *DeleteApplicationAccessPointRequest {
	s.Name = &v
	return s
}

func (s *DeleteApplicationAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
