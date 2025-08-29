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
	// Id of the request
	//
	// example:
	//
	// 2A734D87-2212-5C84-B63A-1AC87CA843D4
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
