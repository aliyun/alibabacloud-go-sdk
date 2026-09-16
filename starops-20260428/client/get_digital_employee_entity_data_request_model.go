// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeEntityDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetDigitalEmployeeEntityDataRequest
	GetFrom() *int64
	SetQuery(v string) *GetDigitalEmployeeEntityDataRequest
	GetQuery() *string
	SetTo(v int64) *GetDigitalEmployeeEntityDataRequest
	GetTo() *int64
}

type GetDigitalEmployeeEntityDataRequest struct {
	// The UNIX timestamp, in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1770000000
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// Only the .entity or .topo semantic sources of the current digital employee can be queried. Explicit access to internal storage or external data sources is prohibited.
	//
	// This parameter is required.
	//
	// example:
	//
	// .entity with(type=\\"deployment\\", domain=\\"k8s\\") | limit 100
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The UNIX timestamp, in seconds. The value must be greater than the value of from.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1770003600
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetDigitalEmployeeEntityDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeEntityDataRequest) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeEntityDataRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetDigitalEmployeeEntityDataRequest) GetQuery() *string {
	return s.Query
}

func (s *GetDigitalEmployeeEntityDataRequest) GetTo() *int64 {
	return s.To
}

func (s *GetDigitalEmployeeEntityDataRequest) SetFrom(v int64) *GetDigitalEmployeeEntityDataRequest {
	s.From = &v
	return s
}

func (s *GetDigitalEmployeeEntityDataRequest) SetQuery(v string) *GetDigitalEmployeeEntityDataRequest {
	s.Query = &v
	return s
}

func (s *GetDigitalEmployeeEntityDataRequest) SetTo(v int64) *GetDigitalEmployeeEntityDataRequest {
	s.To = &v
	return s
}

func (s *GetDigitalEmployeeEntityDataRequest) Validate() error {
	return dara.Validate(s)
}
