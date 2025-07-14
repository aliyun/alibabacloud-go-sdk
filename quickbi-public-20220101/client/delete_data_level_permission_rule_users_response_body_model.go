// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLevelPermissionRuleUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataLevelPermissionRuleUsersResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteDataLevelPermissionRuleUsersResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteDataLevelPermissionRuleUsersResponseBody
	GetSuccess() *bool
}

type DeleteDataLevelPermissionRuleUsersResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLevelPermissionRuleUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLevelPermissionRuleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) SetRequestId(v string) *DeleteDataLevelPermissionRuleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) SetResult(v bool) *DeleteDataLevelPermissionRuleUsersResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) SetSuccess(v bool) *DeleteDataLevelPermissionRuleUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
