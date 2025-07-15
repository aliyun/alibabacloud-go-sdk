// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutId(v string) *ModifyCasterLayoutResponseBody
	GetLayoutId() *string
	SetRequestId(v string) *ModifyCasterLayoutResponseBody
	GetRequestId() *string
}

type ModifyCasterLayoutResponseBody struct {
	// The ID of the layout. You can use this ID to query information about the layout.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCasterLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutResponseBody) GetLayoutId() *string {
	return s.LayoutId
}

func (s *ModifyCasterLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCasterLayoutResponseBody) SetLayoutId(v string) *ModifyCasterLayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *ModifyCasterLayoutResponseBody) SetRequestId(v string) *ModifyCasterLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
