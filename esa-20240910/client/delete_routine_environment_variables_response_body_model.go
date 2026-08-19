// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineEnvironmentVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedKeys(v []*string) *DeleteRoutineEnvironmentVariablesResponseBody
	GetDeletedKeys() []*string
	SetFailedKeys(v []*string) *DeleteRoutineEnvironmentVariablesResponseBody
	GetFailedKeys() []*string
	SetRequestId(v string) *DeleteRoutineEnvironmentVariablesResponseBody
	GetRequestId() *string
}

type DeleteRoutineEnvironmentVariablesResponseBody struct {
	// The list of environment variable keys that were deleted successfully.
	DeletedKeys []*string `json:"DeletedKeys,omitempty" xml:"DeletedKeys,omitempty" type:"Repeated"`
	// The list of environment variable keys that failed to be deleted.
	FailedKeys []*string `json:"FailedKeys,omitempty" xml:"FailedKeys,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoutineEnvironmentVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineEnvironmentVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) GetDeletedKeys() []*string {
	return s.DeletedKeys
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) GetFailedKeys() []*string {
	return s.FailedKeys
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) SetDeletedKeys(v []*string) *DeleteRoutineEnvironmentVariablesResponseBody {
	s.DeletedKeys = v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) SetFailedKeys(v []*string) *DeleteRoutineEnvironmentVariablesResponseBody {
	s.FailedKeys = v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) SetRequestId(v string) *DeleteRoutineEnvironmentVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
