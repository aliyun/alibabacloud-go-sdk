// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetExperimentLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetExperimentLabelsResponseBody
	GetRequestId() *string
}

type SetExperimentLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetExperimentLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetExperimentLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *SetExperimentLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetExperimentLabelsResponseBody) SetRequestId(v string) *SetExperimentLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetExperimentLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
