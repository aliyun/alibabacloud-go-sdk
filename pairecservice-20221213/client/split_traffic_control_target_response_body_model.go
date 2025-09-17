// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SplitTrafficControlTargetResponseBody
	GetRequestId() *string
}

type SplitTrafficControlTargetResponseBody struct {
	// Id of the request。
	//
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SplitTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SplitTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *SplitTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SplitTrafficControlTargetResponseBody) SetRequestId(v string) *SplitTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SplitTrafficControlTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
