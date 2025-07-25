// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrainingJobResponseBody
	GetRequestId() *string
}

type DeleteTrainingJobResponseBody struct {
	// example:
	//
	// 4cc83062-9bcb-4ab3-979e-2e571a35834f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrainingJobResponseBody) SetRequestId(v string) *DeleteTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrainingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
