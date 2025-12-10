// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExperimentResponseBody
	GetRequestId() *string
}

type DeleteExperimentResponseBody struct {
	// example:
	//
	// 2ABF5D32-C9EE-55AE-92EE-DB08E8988AD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperimentResponseBody) SetRequestId(v string) *DeleteExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
