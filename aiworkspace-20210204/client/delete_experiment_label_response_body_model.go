// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExperimentLabelResponseBody
	GetRequestId() *string
}

type DeleteExperimentLabelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperimentLabelResponseBody) SetRequestId(v string) *DeleteExperimentLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
