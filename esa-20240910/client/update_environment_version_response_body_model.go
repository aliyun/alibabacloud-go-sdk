// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEnvironmentVersionResponseBody
	GetRequestId() *string
}

type UpdateEnvironmentVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvironmentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvironmentVersionResponseBody) SetRequestId(v string) *UpdateEnvironmentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvironmentVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
