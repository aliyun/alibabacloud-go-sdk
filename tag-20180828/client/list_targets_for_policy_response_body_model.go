// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetsForPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsRd(v bool) *ListTargetsForPolicyResponseBody
	GetIsRd() *bool
	SetNextToken(v string) *ListTargetsForPolicyResponseBody
	GetNextToken() *string
	SetRdId(v string) *ListTargetsForPolicyResponseBody
	GetRdId() *string
	SetRequestId(v string) *ListTargetsForPolicyResponseBody
	GetRequestId() *string
	SetTargets(v []*ListTargetsForPolicyResponseBodyTargets) *ListTargetsForPolicyResponseBody
	GetTargets() []*ListTargetsForPolicyResponseBodyTargets
}

type ListTargetsForPolicyResponseBody struct {
	// Indicates whether the object belongs to the resource directory. Valid values:
	//
	// 	- true: The object belongs to the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- false: The object does not belong to the resource directory. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// example:
	//
	// true
	IsRd *bool `json:"IsRd,omitempty" xml:"IsRd,omitempty"`
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource directory.
	//
	// >  This parameter is returned only if you use the Tag Policy feature in multi-account mode.
	//
	// example:
	//
	// rd-3G****
	RdId *string `json:"RdId,omitempty" xml:"RdId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2EE71C8D-6DB8-56AC-8B05-3D4C0116E6A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The objects to which the tag policy is attached.
	Targets []*ListTargetsForPolicyResponseBodyTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s ListTargetsForPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsForPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyResponseBody) GetIsRd() *bool {
	return s.IsRd
}

func (s *ListTargetsForPolicyResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTargetsForPolicyResponseBody) GetRdId() *string {
	return s.RdId
}

func (s *ListTargetsForPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTargetsForPolicyResponseBody) GetTargets() []*ListTargetsForPolicyResponseBodyTargets {
	return s.Targets
}

func (s *ListTargetsForPolicyResponseBody) SetIsRd(v bool) *ListTargetsForPolicyResponseBody {
	s.IsRd = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetNextToken(v string) *ListTargetsForPolicyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetRdId(v string) *ListTargetsForPolicyResponseBody {
	s.RdId = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetRequestId(v string) *ListTargetsForPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetTargets(v []*ListTargetsForPolicyResponseBodyTargets) *ListTargetsForPolicyResponseBody {
	s.Targets = v
	return s
}

func (s *ListTargetsForPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTargetsForPolicyResponseBodyTargets struct {
	// The ID of the object.
	//
	// example:
	//
	// 195320939469****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// 	- ROOT: the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- FOLDER: a folder other than the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- ACCOUNT: a member in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// example:
	//
	// ACCOUNT
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListTargetsForPolicyResponseBodyTargets) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsForPolicyResponseBodyTargets) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyResponseBodyTargets) GetTargetId() *string {
	return s.TargetId
}

func (s *ListTargetsForPolicyResponseBodyTargets) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListTargetsForPolicyResponseBodyTargets) SetTargetId(v string) *ListTargetsForPolicyResponseBodyTargets {
	s.TargetId = &v
	return s
}

func (s *ListTargetsForPolicyResponseBodyTargets) SetTargetType(v int32) *ListTargetsForPolicyResponseBodyTargets {
	s.TargetType = &v
	return s
}

func (s *ListTargetsForPolicyResponseBodyTargets) Validate() error {
	return dara.Validate(s)
}
