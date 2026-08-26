// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterResponseBody
	GetCasterId() *string
	SetRequestId(v string) *DeleteCasterResponseBody
	GetRequestId() *string
}

type DeleteCasterResponseBody struct {
	// The ID of the production studio. This ID is used as a request parameter to add video resources, layouts, components, and program lists, or to query the layout list.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterResponseBody) SetCasterId(v string) *DeleteCasterResponseBody {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterResponseBody) SetRequestId(v string) *DeleteCasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterResponseBody) Validate() error {
	return dara.Validate(s)
}
