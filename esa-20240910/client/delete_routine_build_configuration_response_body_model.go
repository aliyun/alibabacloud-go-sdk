// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineBuildConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRoutineBuildConfigurationResponseBody
	GetRequestId() *string
}

type DeleteRoutineBuildConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoutineBuildConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineBuildConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineBuildConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineBuildConfigurationResponseBody) SetRequestId(v string) *DeleteRoutineBuildConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineBuildConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
