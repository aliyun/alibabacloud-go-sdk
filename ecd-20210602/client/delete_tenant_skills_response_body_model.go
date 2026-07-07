// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTenantSkillsResponseBody
	GetRequestId() *string
}

type DeleteTenantSkillsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTenantSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTenantSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTenantSkillsResponseBody) SetRequestId(v string) *DeleteTenantSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTenantSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
