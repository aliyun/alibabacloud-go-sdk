// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutId(v string) *ModifyAppLayoutResponseBody
	GetLayoutId() *string
	SetRequestId(v string) *ModifyAppLayoutResponseBody
	GetRequestId() *string
}

type ModifyAppLayoutResponseBody struct {
	// example:
	//
	// 167466539798442****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppLayoutResponseBody) GetLayoutId() *string {
	return s.LayoutId
}

func (s *ModifyAppLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppLayoutResponseBody) SetLayoutId(v string) *ModifyAppLayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *ModifyAppLayoutResponseBody) SetRequestId(v string) *ModifyAppLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
