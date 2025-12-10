// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExperimentMetaResponseBody
	GetRequestId() *string
}

type UpdateExperimentMetaResponseBody struct {
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentMetaResponseBody) SetRequestId(v string) *UpdateExperimentMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
