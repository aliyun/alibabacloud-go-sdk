// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineEnvironmentVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetRoutineEnvironmentVariablesResponseBody
	GetRequestId() *string
	SetSetKeys(v []*string) *SetRoutineEnvironmentVariablesResponseBody
	GetSetKeys() []*string
}

type SetRoutineEnvironmentVariablesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of environment variable keys that were set successfully.
	//
	// example:
	//
	// ["key1","key2"]
	SetKeys []*string `json:"SetKeys,omitempty" xml:"SetKeys,omitempty" type:"Repeated"`
}

func (s SetRoutineEnvironmentVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineEnvironmentVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *SetRoutineEnvironmentVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRoutineEnvironmentVariablesResponseBody) GetSetKeys() []*string {
	return s.SetKeys
}

func (s *SetRoutineEnvironmentVariablesResponseBody) SetRequestId(v string) *SetRoutineEnvironmentVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesResponseBody) SetSetKeys(v []*string) *SetRoutineEnvironmentVariablesResponseBody {
	s.SetKeys = v
	return s
}

func (s *SetRoutineEnvironmentVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
