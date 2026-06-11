// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhitePortalMenuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWhitePortalMenuResponseBody
	GetRequestId() *string
	SetResult(v []*ListWhitePortalMenuResponseBodyResult) *ListWhitePortalMenuResponseBody
	GetResult() []*ListWhitePortalMenuResponseBodyResult
	SetSuccess(v bool) *ListWhitePortalMenuResponseBody
	GetSuccess() *bool
}

type ListWhitePortalMenuResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC4E1**********0DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The whitelist.
	Result []*ListWhitePortalMenuResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWhitePortalMenuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWhitePortalMenuResponseBody) GoString() string {
	return s.String()
}

func (s *ListWhitePortalMenuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWhitePortalMenuResponseBody) GetResult() []*ListWhitePortalMenuResponseBodyResult {
	return s.Result
}

func (s *ListWhitePortalMenuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWhitePortalMenuResponseBody) SetRequestId(v string) *ListWhitePortalMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWhitePortalMenuResponseBody) SetResult(v []*ListWhitePortalMenuResponseBodyResult) *ListWhitePortalMenuResponseBody {
	s.Result = v
	return s
}

func (s *ListWhitePortalMenuResponseBody) SetSuccess(v bool) *ListWhitePortalMenuResponseBody {
	s.Success = &v
	return s
}

func (s *ListWhitePortalMenuResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWhitePortalMenuResponseBodyResult struct {
	// The authorization type for the menu. Valid values:
	//
	// - 1: View
	//
	// - 3: Export and view
	//
	// example:
	//
	// 1
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// The ID of the authorization object. If the authorization is at the workspace or organization level, this parameter returns the workspace ID or organization ID.
	//
	// example:
	//
	// 8a4***********1e769
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The type of the authorization object. Valid values:
	//
	// - 0: User
	//
	// - 1: User group
	//
	// - 3: Workspace or organization level
	//
	// example:
	//
	// 0
	ReceiverType *int32 `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
}

func (s ListWhitePortalMenuResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListWhitePortalMenuResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWhitePortalMenuResponseBodyResult) GetAuthPointsValue() *int32 {
	return s.AuthPointsValue
}

func (s *ListWhitePortalMenuResponseBodyResult) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *ListWhitePortalMenuResponseBodyResult) GetReceiverType() *int32 {
	return s.ReceiverType
}

func (s *ListWhitePortalMenuResponseBodyResult) SetAuthPointsValue(v int32) *ListWhitePortalMenuResponseBodyResult {
	s.AuthPointsValue = &v
	return s
}

func (s *ListWhitePortalMenuResponseBodyResult) SetReceiverId(v string) *ListWhitePortalMenuResponseBodyResult {
	s.ReceiverId = &v
	return s
}

func (s *ListWhitePortalMenuResponseBodyResult) SetReceiverType(v int32) *ListWhitePortalMenuResponseBodyResult {
	s.ReceiverType = &v
	return s
}

func (s *ListWhitePortalMenuResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
