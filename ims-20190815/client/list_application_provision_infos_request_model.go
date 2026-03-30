// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationProvisionInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceType(v string) *ListApplicationProvisionInfosRequest
	GetSourceType() *string
}

type ListApplicationProvisionInfosRequest struct {
	// The source of the applications. Valid values:
	//
	// 	- inner: The applications are from the current account.
	//
	// 	- external: The applications are from other accounts.
	//
	// example:
	//
	// external
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListApplicationProvisionInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListApplicationProvisionInfosRequest) SetSourceType(v string) *ListApplicationProvisionInfosRequest {
	s.SourceType = &v
	return s
}

func (s *ListApplicationProvisionInfosRequest) Validate() error {
	return dara.Validate(s)
}
