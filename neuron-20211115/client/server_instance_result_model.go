// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServerInstanceResult interface {
	dara.Model
	String() string
	GoString() string
	SetIntTotal(v int32) *ServerInstanceResult
	GetIntTotal() *int32
	SetIps(v []*string) *ServerInstanceResult
	GetIps() []*string
	SetRequestId(v string) *ServerInstanceResult
	GetRequestId() *string
	SetTotal(v int64) *ServerInstanceResult
	GetTotal() *int64
}

type ServerInstanceResult struct {
	IntTotal  *int32    `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ServerInstanceResult) String() string {
	return dara.Prettify(s)
}

func (s ServerInstanceResult) GoString() string {
	return s.String()
}

func (s *ServerInstanceResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *ServerInstanceResult) GetIps() []*string {
	return s.Ips
}

func (s *ServerInstanceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ServerInstanceResult) GetTotal() *int64 {
	return s.Total
}

func (s *ServerInstanceResult) SetIntTotal(v int32) *ServerInstanceResult {
	s.IntTotal = &v
	return s
}

func (s *ServerInstanceResult) SetIps(v []*string) *ServerInstanceResult {
	s.Ips = v
	return s
}

func (s *ServerInstanceResult) SetRequestId(v string) *ServerInstanceResult {
	s.RequestId = &v
	return s
}

func (s *ServerInstanceResult) SetTotal(v int64) *ServerInstanceResult {
	s.Total = &v
	return s
}

func (s *ServerInstanceResult) Validate() error {
	return dara.Validate(s)
}
