// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearPairDrillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearPairDrillResponseBody
	GetRequestId() *string
}

type ClearPairDrillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearPairDrillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearPairDrillResponseBody) GoString() string {
	return s.String()
}

func (s *ClearPairDrillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearPairDrillResponseBody) SetRequestId(v string) *ClearPairDrillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearPairDrillResponseBody) Validate() error {
	return dara.Validate(s)
}
