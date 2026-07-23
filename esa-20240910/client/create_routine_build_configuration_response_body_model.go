// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRoutineBuildConfigurationResponseBody
	GetRequestId() *string
	SetRoutineBuildConfigurationId(v int64) *CreateRoutineBuildConfigurationResponseBody
	GetRoutineBuildConfigurationId() *int64
}

type CreateRoutineBuildConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ER build configuration ID.
	//
	// example:
	//
	// 3472165674357056
	RoutineBuildConfigurationId *int64 `json:"RoutineBuildConfigurationId,omitempty" xml:"RoutineBuildConfigurationId,omitempty"`
}

func (s CreateRoutineBuildConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineBuildConfigurationResponseBody) GetRoutineBuildConfigurationId() *int64 {
	return s.RoutineBuildConfigurationId
}

func (s *CreateRoutineBuildConfigurationResponseBody) SetRequestId(v string) *CreateRoutineBuildConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineBuildConfigurationResponseBody) SetRoutineBuildConfigurationId(v int64) *CreateRoutineBuildConfigurationResponseBody {
	s.RoutineBuildConfigurationId = &v
	return s
}

func (s *CreateRoutineBuildConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
