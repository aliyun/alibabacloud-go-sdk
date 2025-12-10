// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishExperimentResponseBody
	GetRequestId() *string
}

type PublishExperimentResponseBody struct {
	// example:
	//
	// 5A26A7FA-EEF0-5A6B-BA76-06067547C11F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *PublishExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishExperimentResponseBody) SetRequestId(v string) *PublishExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
