// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterComponentResponseBody
	GetCasterId() *string
	SetComponentId(v string) *DeleteCasterComponentResponseBody
	GetComponentId() *string
	SetRequestId(v string) *DeleteCasterComponentResponseBody
	GetRequestId() *string
}

type DeleteCasterComponentResponseBody struct {
	// The ID of the production studio. You can use the ID as a request parameter in the API operation that is called to query the components in the production studio, add an episode list to the production studio, or modify a component in the production studio.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The component ID. You can use the ID as a request parameter in the API operation that is called to query the component in the production studio or modify the component in the production studio.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterComponentResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterComponentResponseBody) GetComponentId() *string {
	return s.ComponentId
}

func (s *DeleteCasterComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterComponentResponseBody) SetCasterId(v string) *DeleteCasterComponentResponseBody {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterComponentResponseBody) SetComponentId(v string) *DeleteCasterComponentResponseBody {
	s.ComponentId = &v
	return s
}

func (s *DeleteCasterComponentResponseBody) SetRequestId(v string) *DeleteCasterComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
