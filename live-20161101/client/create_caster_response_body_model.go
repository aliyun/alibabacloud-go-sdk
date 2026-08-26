// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *CreateCasterResponseBody
	GetCasterId() *string
	SetRequestId(v string) *CreateCasterResponseBody
	GetRequestId() *string
}

type CreateCasterResponseBody struct {
	// The ID of the production studio. You can use this ID as a request parameter to query stream URLs, start the production studio, add video resources, add layouts, query the layout list, add components, and add playlists.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCasterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCasterResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *CreateCasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCasterResponseBody) SetCasterId(v string) *CreateCasterResponseBody {
	s.CasterId = &v
	return s
}

func (s *CreateCasterResponseBody) SetRequestId(v string) *CreateCasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCasterResponseBody) Validate() error {
	return dara.Validate(s)
}
