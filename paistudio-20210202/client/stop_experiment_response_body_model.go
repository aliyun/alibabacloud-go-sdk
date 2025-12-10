// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopExperimentResponseBody
	GetRequestId() *string
}

type StopExperimentResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *StopExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopExperimentResponseBody) SetRequestId(v string) *StopExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
