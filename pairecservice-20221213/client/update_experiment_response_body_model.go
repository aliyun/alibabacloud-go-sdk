// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExperimentResponseBody
	GetRequestId() *string
}

type UpdateExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A760D972-1475-58C0-BBB3-92B5FB08904F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentResponseBody) SetRequestId(v string) *UpdateExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
