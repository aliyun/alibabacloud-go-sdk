// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyNetworkAclEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyNetworkAclEntriesResponseBody
	GetRequestId() *string
}

type CopyNetworkAclEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6608E72F-F276-440F-ABEF-419971CEC4D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyNetworkAclEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyNetworkAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *CopyNetworkAclEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyNetworkAclEntriesResponseBody) SetRequestId(v string) *CopyNetworkAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyNetworkAclEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
