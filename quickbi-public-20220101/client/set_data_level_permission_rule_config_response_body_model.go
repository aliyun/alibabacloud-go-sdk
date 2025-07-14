// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionRuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDataLevelPermissionRuleConfigResponseBody
	GetRequestId() *string
	SetResult(v string) *SetDataLevelPermissionRuleConfigResponseBody
	GetResult() *string
	SetSuccess(v bool) *SetDataLevelPermissionRuleConfigResponseBody
	GetSuccess() *bool
}

type SetDataLevelPermissionRuleConfigResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataLevelPermissionRuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) GetResult() *string {
	return s.Result
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) SetRequestId(v string) *SetDataLevelPermissionRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) SetResult(v string) *SetDataLevelPermissionRuleConfigResponseBody {
	s.Result = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) SetSuccess(v bool) *SetDataLevelPermissionRuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
