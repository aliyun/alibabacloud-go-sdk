// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStudioLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStudioLayoutResponseBody
	GetRequestId() *string
}

type ModifyStudioLayoutResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStudioLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStudioLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStudioLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStudioLayoutResponseBody) SetRequestId(v string) *ModifyStudioLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStudioLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
