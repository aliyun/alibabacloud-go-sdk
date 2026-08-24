// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVirusScanAdditionalListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListIds(v []*string) *RemoveVirusScanAdditionalListsRequest
	GetListIds() []*string
}

type RemoveVirusScanAdditionalListsRequest struct {
	// The collection of entry IDs to remove. At least one entry ID must be specified.
	//
	// This parameter is required.
	ListIds []*string `json:"ListIds,omitempty" xml:"ListIds,omitempty" type:"Repeated"`
}

func (s RemoveVirusScanAdditionalListsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveVirusScanAdditionalListsRequest) GoString() string {
	return s.String()
}

func (s *RemoveVirusScanAdditionalListsRequest) GetListIds() []*string {
	return s.ListIds
}

func (s *RemoveVirusScanAdditionalListsRequest) SetListIds(v []*string) *RemoveVirusScanAdditionalListsRequest {
	s.ListIds = v
	return s
}

func (s *RemoveVirusScanAdditionalListsRequest) Validate() error {
	return dara.Validate(s)
}
