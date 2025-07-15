// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterProgramResponseBody
	GetCasterId() *string
	SetRequestId(v string) *DeleteCasterProgramResponseBody
	GetRequestId() *string
}

type DeleteCasterProgramResponseBody struct {
	// The ID of the production studio. You can specify the ID as a request parameter in the API operation that is used to add, delete, or modify episodes in the production studio.
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

func (s DeleteCasterProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterProgramResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterProgramResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterProgramResponseBody) SetCasterId(v string) *DeleteCasterProgramResponseBody {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterProgramResponseBody) SetRequestId(v string) *DeleteCasterProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterProgramResponseBody) Validate() error {
	return dara.Validate(s)
}
