// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int32) *GetEntityStoreDataRequest
	GetFrom() *int32
	SetQuery(v string) *GetEntityStoreDataRequest
	GetQuery() *string
	SetTo(v int32) *GetEntityStoreDataRequest
	GetTo() *int32
}

type GetEntityStoreDataRequest struct {
	// Start time of the query.
	//
	// Unix timestamp format, representing the number of seconds since 1970-1-1 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721767203
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// Query statement
	//
	// This parameter is required.
	//
	// example:
	//
	// .entity with(domain=\\"acs\\", type=\\"acs.k8s.node\\") | limit 0, 10
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// End time of the query.
	//
	// Unix timestamp format, representing the number of seconds since 1970-1-1 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721767283
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetEntityStoreDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataRequest) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataRequest) GetFrom() *int32 {
	return s.From
}

func (s *GetEntityStoreDataRequest) GetQuery() *string {
	return s.Query
}

func (s *GetEntityStoreDataRequest) GetTo() *int32 {
	return s.To
}

func (s *GetEntityStoreDataRequest) SetFrom(v int32) *GetEntityStoreDataRequest {
	s.From = &v
	return s
}

func (s *GetEntityStoreDataRequest) SetQuery(v string) *GetEntityStoreDataRequest {
	s.Query = &v
	return s
}

func (s *GetEntityStoreDataRequest) SetTo(v int32) *GetEntityStoreDataRequest {
	s.To = &v
	return s
}

func (s *GetEntityStoreDataRequest) Validate() error {
	return dara.Validate(s)
}
