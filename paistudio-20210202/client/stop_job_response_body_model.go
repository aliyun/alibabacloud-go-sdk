// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopJobResponseBody
	GetRequestId() *string
}

type StopJobResponseBody struct {
	// example:
	//
	// 431C53C4-BDD0-588F-8081-4437B00852B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopJobResponseBody) Validate() error {
	return dara.Validate(s)
}
