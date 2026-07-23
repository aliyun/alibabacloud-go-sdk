// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRoutineBuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRoutineBuildResponseBody
	GetRequestId() *string
}

type StopRoutineBuildResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRoutineBuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRoutineBuildResponseBody) GoString() string {
	return s.String()
}

func (s *StopRoutineBuildResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRoutineBuildResponseBody) SetRequestId(v string) *StopRoutineBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRoutineBuildResponseBody) Validate() error {
	return dara.Validate(s)
}
