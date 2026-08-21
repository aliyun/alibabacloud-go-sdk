// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationL7SwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListPrivateAccessApplicationL7SwitchesRequest
	GetApplicationIds() []*string
}

type ListPrivateAccessApplicationL7SwitchesRequest struct {
	// The IDs of internal-facing applications. You can specify up to 100 internal-facing application IDs.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessApplicationL7SwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationL7SwitchesRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationL7SwitchesRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListPrivateAccessApplicationL7SwitchesRequest) SetApplicationIds(v []*string) *ListPrivateAccessApplicationL7SwitchesRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesRequest) Validate() error {
	return dara.Validate(s)
}
