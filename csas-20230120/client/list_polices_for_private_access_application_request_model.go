// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForPrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListPolicesForPrivateAccessApplicationRequest
	GetApplicationIds() []*string
}

type ListPolicesForPrivateAccessApplicationRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListPolicesForPrivateAccessApplicationRequest) SetApplicationIds(v []*string) *ListPolicesForPrivateAccessApplicationRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
