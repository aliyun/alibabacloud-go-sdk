// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiDatasourceParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiDatasourceParametersResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyApiDatasourceParametersResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ModifyApiDatasourceParametersResponseBody
	GetSuccess() *bool
}

type ModifyApiDatasourceParametersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApiDatasourceParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiDatasourceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiDatasourceParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiDatasourceParametersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyApiDatasourceParametersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyApiDatasourceParametersResponseBody) SetRequestId(v string) *ModifyApiDatasourceParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponseBody) SetResult(v bool) *ModifyApiDatasourceParametersResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponseBody) SetSuccess(v bool) *ModifyApiDatasourceParametersResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
