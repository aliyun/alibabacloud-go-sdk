// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchClusterMasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchClusterMasterResponseBody
	GetRequestId() *string
}

type SwitchClusterMasterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchClusterMasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchClusterMasterResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchClusterMasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchClusterMasterResponseBody) SetRequestId(v string) *SwitchClusterMasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchClusterMasterResponseBody) Validate() error {
	return dara.Validate(s)
}
