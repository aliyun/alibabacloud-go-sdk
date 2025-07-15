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
	// The ID of the production studio. You can specify the ID in a request to query the streaming URLs of the production studio, start the production studio, add a video resource, a layout, a component, or a playlist to the production studio, or query layouts of the production studio.
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
