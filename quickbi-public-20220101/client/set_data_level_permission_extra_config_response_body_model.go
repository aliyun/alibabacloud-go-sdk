// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionExtraConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDataLevelPermissionExtraConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetDataLevelPermissionExtraConfigResponseBody
	GetResult() *bool
	SetSuccess(v bool) *SetDataLevelPermissionExtraConfigResponseBody
	GetSuccess() *bool
}

type SetDataLevelPermissionExtraConfigResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B70E1FBD-E533-52F2-A7A1-E02B92F78DDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result of the API execution. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataLevelPermissionExtraConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionExtraConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) SetRequestId(v string) *SetDataLevelPermissionExtraConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) SetResult(v bool) *SetDataLevelPermissionExtraConfigResponseBody {
	s.Result = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) SetSuccess(v bool) *SetDataLevelPermissionExtraConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
