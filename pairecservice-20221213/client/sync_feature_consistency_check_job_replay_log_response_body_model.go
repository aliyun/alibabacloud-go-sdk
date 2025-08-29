// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncFeatureConsistencyCheckJobReplayLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncFeatureConsistencyCheckJobReplayLogResponseBody
	GetRequestId() *string
}

type SyncFeatureConsistencyCheckJobReplayLogResponseBody struct {
	// example:
	//
	// C7D0B48F-0105-52B9-B60A-FA7606E2234D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponseBody) GoString() string {
	return s.String()
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponseBody) SetRequestId(v string) *SyncFeatureConsistencyCheckJobReplayLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponseBody) Validate() error {
	return dara.Validate(s)
}
