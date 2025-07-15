// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStudioLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutId(v string) *AddStudioLayoutResponseBody
	GetLayoutId() *string
	SetRequestId(v string) *AddStudioLayoutResponseBody
	GetRequestId() *string
}

type AddStudioLayoutResponseBody struct {
	// The ID of the layout. You can use the ID as a request parameter in the following operations: DeleteStudioLayout, ModifyStudioLayout, and DescribeStudioLayouts.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddStudioLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddStudioLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *AddStudioLayoutResponseBody) GetLayoutId() *string {
	return s.LayoutId
}

func (s *AddStudioLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddStudioLayoutResponseBody) SetLayoutId(v string) *AddStudioLayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *AddStudioLayoutResponseBody) SetRequestId(v string) *AddStudioLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddStudioLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
