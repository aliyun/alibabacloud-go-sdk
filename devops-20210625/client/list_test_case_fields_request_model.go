// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTestCaseFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceIdentifier(v string) *ListTestCaseFieldsRequest
	GetSpaceIdentifier() *string
}

type ListTestCaseFieldsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
}

func (s ListTestCaseFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTestCaseFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListTestCaseFieldsRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListTestCaseFieldsRequest) SetSpaceIdentifier(v string) *ListTestCaseFieldsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListTestCaseFieldsRequest) Validate() error {
	return dara.Validate(s)
}
