// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRCSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncRCSecurityGroupResponseBody
	GetRequestId() *string
}

type SyncRCSecurityGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 300333A0-68E5-59CE-94AD-75153D17639E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncRCSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncRCSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SyncRCSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncRCSecurityGroupResponseBody) SetRequestId(v string) *SyncRCSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncRCSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
