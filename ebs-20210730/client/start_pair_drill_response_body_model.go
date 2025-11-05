// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPairDrillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDrillId(v string) *StartPairDrillResponseBody
	GetDrillId() *string
	SetRequestId(v string) *StartPairDrillResponseBody
	GetRequestId() *string
}

type StartPairDrillResponseBody struct {
	// The drill ID.
	//
	// example:
	//
	// drill-xxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPairDrillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPairDrillResponseBody) GoString() string {
	return s.String()
}

func (s *StartPairDrillResponseBody) GetDrillId() *string {
	return s.DrillId
}

func (s *StartPairDrillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPairDrillResponseBody) SetDrillId(v string) *StartPairDrillResponseBody {
	s.DrillId = &v
	return s
}

func (s *StartPairDrillResponseBody) SetRequestId(v string) *StartPairDrillResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPairDrillResponseBody) Validate() error {
	return dara.Validate(s)
}
