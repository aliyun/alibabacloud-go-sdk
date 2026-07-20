// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineBuildConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRoutineBuildConfigurationResponseBody
	GetRequestId() *string
}

type UpdateRoutineBuildConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3558df77-8a7a-4060-a900-2d7949403836
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRoutineBuildConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineBuildConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoutineBuildConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRoutineBuildConfigurationResponseBody) SetRequestId(v string) *UpdateRoutineBuildConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
