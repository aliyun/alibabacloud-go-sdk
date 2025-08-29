// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExperimentGroupResponseBody
	GetRequestId() *string
}

type DeleteExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A009D9BE-C85E-57B2-AE05-BD78BB6EBF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperimentGroupResponseBody) SetRequestId(v string) *DeleteExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
