// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePublicIpsFromEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemovePublicIpsFromEpnInstanceResponseBody
	GetRequestId() *string
}

type RemovePublicIpsFromEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E	 Request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePublicIpsFromEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePublicIpsFromEpnInstanceResponseBody) SetRequestId(v string) *RemovePublicIpsFromEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
