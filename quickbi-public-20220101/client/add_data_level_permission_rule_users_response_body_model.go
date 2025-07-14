// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataLevelPermissionRuleUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDataLevelPermissionRuleUsersResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddDataLevelPermissionRuleUsersResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AddDataLevelPermissionRuleUsersResponseBody
	GetSuccess() *bool
}

type AddDataLevelPermissionRuleUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:\\n\\n	- true: The request was successful.\\n	- false: The request failed.\\n
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:\\n\\n	- true: The request was successful.\\n	- false: The request failed.\\n
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDataLevelPermissionRuleUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataLevelPermissionRuleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) SetRequestId(v string) *AddDataLevelPermissionRuleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) SetResult(v bool) *AddDataLevelPermissionRuleUsersResponseBody {
	s.Result = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) SetSuccess(v bool) *AddDataLevelPermissionRuleUsersResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
