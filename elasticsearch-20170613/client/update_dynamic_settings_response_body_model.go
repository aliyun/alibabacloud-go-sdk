// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDynamicSettingsResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDynamicSettingsResponseBody
	GetResult() *bool
}

type UpdateDynamicSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of updating the dynamic configurations of the cluster. A value of true indicates that the update was successful, and a value of false indicates that the update failed.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateDynamicSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDynamicSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDynamicSettingsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDynamicSettingsResponseBody) SetRequestId(v string) *UpdateDynamicSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDynamicSettingsResponseBody) SetResult(v bool) *UpdateDynamicSettingsResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDynamicSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
