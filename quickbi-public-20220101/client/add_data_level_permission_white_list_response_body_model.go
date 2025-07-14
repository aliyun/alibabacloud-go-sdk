// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataLevelPermissionWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDataLevelPermissionWhiteListResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddDataLevelPermissionWhiteListResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AddDataLevelPermissionWhiteListResponseBody
	GetSuccess() *bool
}

type AddDataLevelPermissionWhiteListResponseBody struct {
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

func (s AddDataLevelPermissionWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataLevelPermissionWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataLevelPermissionWhiteListResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddDataLevelPermissionWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataLevelPermissionWhiteListResponseBody) SetRequestId(v string) *AddDataLevelPermissionWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponseBody) SetResult(v bool) *AddDataLevelPermissionWhiteListResponseBody {
	s.Result = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponseBody) SetSuccess(v bool) *AddDataLevelPermissionWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
