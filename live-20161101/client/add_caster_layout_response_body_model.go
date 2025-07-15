// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutId(v string) *AddCasterLayoutResponseBody
	GetLayoutId() *string
	SetRequestId(v string) *AddCasterLayoutResponseBody
	GetRequestId() *string
}

type AddCasterLayoutResponseBody struct {
	// The ID of the layout.
	//
	// Record the ID as it can be used to manage the layout being created.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCasterLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutResponseBody) GetLayoutId() *string {
	return s.LayoutId
}

func (s *AddCasterLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterLayoutResponseBody) SetLayoutId(v string) *AddCasterLayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *AddCasterLayoutResponseBody) SetRequestId(v string) *AddCasterLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
