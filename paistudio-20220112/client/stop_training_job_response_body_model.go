// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTrainingJobResponseBody
	GetRequestId() *string
}

type StopTrainingJobResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTrainingJobResponseBody) SetRequestId(v string) *StopTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTrainingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
