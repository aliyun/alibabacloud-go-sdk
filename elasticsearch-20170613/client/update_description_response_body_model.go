// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDescriptionResponseBody
	GetRequestId() *string
	SetResult(v *UpdateDescriptionResponseBodyResult) *UpdateDescriptionResponseBody
	GetResult() *UpdateDescriptionResponseBodyResult
}

type UpdateDescriptionResponseBody struct {
	// example:
	//
	// FDF34727-1664-44C1-A8DA-3EB72D60****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateDescriptionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDescriptionResponseBody) GetResult() *UpdateDescriptionResponseBodyResult {
	return s.Result
}

func (s *UpdateDescriptionResponseBody) SetRequestId(v string) *UpdateDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDescriptionResponseBody) SetResult(v *UpdateDescriptionResponseBodyResult) *UpdateDescriptionResponseBody {
	s.Result = v
	return s
}

func (s *UpdateDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDescriptionResponseBodyResult struct {
	// example:
	//
	// aliyunes_test_name
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateDescriptionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateDescriptionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateDescriptionResponseBodyResult) SetDescription(v string) *UpdateDescriptionResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateDescriptionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
