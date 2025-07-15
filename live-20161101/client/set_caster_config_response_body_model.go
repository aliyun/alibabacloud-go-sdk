// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *SetCasterConfigResponseBody
	GetCasterId() *string
	SetRequestId(v string) *SetCasterConfigResponseBody
	GetRequestId() *string
}

type SetCasterConfigResponseBody struct {
	// The ID of the production studio. You can specify the ID in a request to query the streaming URLs of the production studio, start the production studio, add a video resource, a layout, a component, or a playlist to the production studio, or query layouts of the production studio.
	//
	// example:
	//
	// b4810848-bcf9-4aef-bd4a-e6bba2d9****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCasterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCasterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetCasterConfigResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *SetCasterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCasterConfigResponseBody) SetCasterId(v string) *SetCasterConfigResponseBody {
	s.CasterId = &v
	return s
}

func (s *SetCasterConfigResponseBody) SetRequestId(v string) *SetCasterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCasterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
