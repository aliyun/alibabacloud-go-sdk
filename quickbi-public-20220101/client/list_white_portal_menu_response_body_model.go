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
	// example:
	//
	// DC4E1**********0DF67E2C3
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListWhitePortalMenuResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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
	// example:
	//
	// 1
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// example:
	//
	// 8a4***********1e769
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
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
