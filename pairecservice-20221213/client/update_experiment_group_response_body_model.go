// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExperimentGroupResponseBody
	GetRequestId() *string
}

type UpdateExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentGroupResponseBody) SetRequestId(v string) *UpdateExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
