// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterLayoutResponseBody
	GetCasterId() *string
	SetLayoutId(v string) *DeleteCasterLayoutResponseBody
	GetLayoutId() *string
	SetRequestId(v string) *DeleteCasterLayoutResponseBody
	GetRequestId() *string
}

type DeleteCasterLayoutResponseBody struct {
	// The ID of the production studio. You can use the ID as a request parameter in the API operation that is used to modify a layout in the production studio, query layouts in the production studio, add a component in the production studio, or query components in the production studio.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the layout. You can use the ID as a request parameter in the API operation that is used to query layouts in the production studio.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterLayoutResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterLayoutResponseBody) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DeleteCasterLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterLayoutResponseBody) SetCasterId(v string) *DeleteCasterLayoutResponseBody {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterLayoutResponseBody) SetLayoutId(v string) *DeleteCasterLayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *DeleteCasterLayoutResponseBody) SetRequestId(v string) *DeleteCasterLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
