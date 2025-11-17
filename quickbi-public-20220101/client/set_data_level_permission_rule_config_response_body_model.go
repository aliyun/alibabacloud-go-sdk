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
	// Request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Successfully saved row and column permission information.
	//
	// example:
	//
	// {
	//
	//     "cubeId": "aa574353-41cf-****-8d0d-455176c101fa",
	//
	//     "hitTakeEffect": 1,
	//
	//     "isOpen": 1,
	//
	//     "permissionMode": "COMPLEX",
	//
	//     "ruleContentModel": {
	//
	//         "ruleContent": {
	//
	//             "pathId": [
	//
	//                 "264b7a970b"
	//
	//             ]
	//
	//         },
	//
	//         "ruleContentJson": "{\\"pathId\\":[\\"264b7a970b\\"]}",
	//
	//         "ruleContentType": "COLUMN_FORBID"
	//
	//     },
	//
	//     "ruleId": "3971fa8e-f7e0-****-b6e3-5b3167dd7247",
	//
	//     "ruleLevelType": "COLUMN_LEVEL",
	//
	//     "ruleName": "test",
	//
	//     "ruleTargetScope": "ALL"
	//
	// }
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. The value range is as follows:
	//
	// - true: The request succeeded
	//
	// - false: The request failed
	//
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
