// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAliasResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyAliasResponseBody
	GetResult() map[string]interface{}
}

type ModifyAliasResponseBody struct {
	// id of request
	//
	// example:
	//
	// F6E3D968-529C-5C40-AFDD-133A8B8FD930
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAliasResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAliasResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyAliasResponseBody) SetRequestId(v string) *ModifyAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAliasResponseBody) SetResult(v map[string]interface{}) *ModifyAliasResponseBody {
	s.Result = v
	return s
}

func (s *ModifyAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
