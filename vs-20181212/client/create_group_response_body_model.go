// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGbId(v string) *CreateGroupResponseBody
	GetGbId() *string
	SetGbIp(v string) *CreateGroupResponseBody
	GetGbIp() *string
	SetGbPort(v int64) *CreateGroupResponseBody
	GetGbPort() *int64
	SetId(v string) *CreateGroupResponseBody
	GetId() *string
	SetRequestId(v string) *CreateGroupResponseBody
	GetRequestId() *string
}

type CreateGroupResponseBody struct {
	// example:
	//
	// 31000000****00000001
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// example:
	//
	// 10.10.10.10
	GbIp *string `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	// example:
	//
	// 5060
	GbPort *int64 `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	// example:
	//
	// 33763****77224964-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) GetGbId() *string {
	return s.GbId
}

func (s *CreateGroupResponseBody) GetGbIp() *string {
	return s.GbIp
}

func (s *CreateGroupResponseBody) GetGbPort() *int64 {
	return s.GbPort
}

func (s *CreateGroupResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupResponseBody) SetGbId(v string) *CreateGroupResponseBody {
	s.GbId = &v
	return s
}

func (s *CreateGroupResponseBody) SetGbIp(v string) *CreateGroupResponseBody {
	s.GbIp = &v
	return s
}

func (s *CreateGroupResponseBody) SetGbPort(v int64) *CreateGroupResponseBody {
	s.GbPort = &v
	return s
}

func (s *CreateGroupResponseBody) SetId(v string) *CreateGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
