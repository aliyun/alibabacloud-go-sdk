// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *ModifyCasterProgramResponseBody
	GetCasterId() *string
	SetRequestId(v string) *ModifyCasterProgramResponseBody
	GetRequestId() *string
}

type ModifyCasterProgramResponseBody struct {
	// The ID of the production studio.
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

func (s ModifyCasterProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterProgramResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCasterProgramResponseBody) SetCasterId(v string) *ModifyCasterProgramResponseBody {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterProgramResponseBody) SetRequestId(v string) *ModifyCasterProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterProgramResponseBody) Validate() error {
	return dara.Validate(s)
}
