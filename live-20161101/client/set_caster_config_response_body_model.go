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
	// The production studio ID. This ID can be used as a request parameter for querying the production studio stream address, starting the production studio, adding video resources, adding layouts, querying the layout list, adding components, and adding a program list.
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
