// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVirusScanAdditionalListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListIds(v []*string) *AddVirusScanAdditionalListsResponseBody
	GetListIds() []*string
	SetRequestId(v string) *AddVirusScanAdditionalListsResponseBody
	GetRequestId() *string
}

type AddVirusScanAdditionalListsResponseBody struct {
	// The list of IDs for the newly added entries. The order is consistent with AdditionalLists in the request.
	ListIds []*string `json:"ListIds,omitempty" xml:"ListIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVirusScanAdditionalListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVirusScanAdditionalListsResponseBody) GoString() string {
	return s.String()
}

func (s *AddVirusScanAdditionalListsResponseBody) GetListIds() []*string {
	return s.ListIds
}

func (s *AddVirusScanAdditionalListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVirusScanAdditionalListsResponseBody) SetListIds(v []*string) *AddVirusScanAdditionalListsResponseBody {
	s.ListIds = v
	return s
}

func (s *AddVirusScanAdditionalListsResponseBody) SetRequestId(v string) *AddVirusScanAdditionalListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVirusScanAdditionalListsResponseBody) Validate() error {
	return dara.Validate(s)
}
