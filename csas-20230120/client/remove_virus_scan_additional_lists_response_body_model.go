// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVirusScanAdditionalListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveVirusScanAdditionalListsResponseBody
	GetRequestId() *string
}

type RemoveVirusScanAdditionalListsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVirusScanAdditionalListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveVirusScanAdditionalListsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVirusScanAdditionalListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveVirusScanAdditionalListsResponseBody) SetRequestId(v string) *RemoveVirusScanAdditionalListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVirusScanAdditionalListsResponseBody) Validate() error {
	return dara.Validate(s)
}
