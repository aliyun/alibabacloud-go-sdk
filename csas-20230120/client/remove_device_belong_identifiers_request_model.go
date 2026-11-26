// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDeviceBelongIdentifiersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifierIds(v []*string) *RemoveDeviceBelongIdentifiersRequest
	GetIdentifierIds() []*string
}

type RemoveDeviceBelongIdentifiersRequest struct {
	// The list of identifier record IDs to remove. You can specify 1 to 1,000 IDs per batch. All IDs must exist and belong to the current tenant. Otherwise, an InvalidParameter error is returned and no records are deleted.
	//
	// This parameter is required.
	IdentifierIds []*string `json:"IdentifierIds,omitempty" xml:"IdentifierIds,omitempty" type:"Repeated"`
}

func (s RemoveDeviceBelongIdentifiersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDeviceBelongIdentifiersRequest) GoString() string {
	return s.String()
}

func (s *RemoveDeviceBelongIdentifiersRequest) GetIdentifierIds() []*string {
	return s.IdentifierIds
}

func (s *RemoveDeviceBelongIdentifiersRequest) SetIdentifierIds(v []*string) *RemoveDeviceBelongIdentifiersRequest {
	s.IdentifierIds = v
	return s
}

func (s *RemoveDeviceBelongIdentifiersRequest) Validate() error {
	return dara.Validate(s)
}
