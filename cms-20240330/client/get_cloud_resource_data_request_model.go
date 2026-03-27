// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudResourceDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int32) *GetCloudResourceDataRequest
	GetFrom() *int32
	SetQuery(v string) *GetCloudResourceDataRequest
	GetQuery() *string
	SetTo(v int32) *GetCloudResourceDataRequest
	GetTo() *int32
}

type GetCloudResourceDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1721767203
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// .entity with(domain=\\"acs\\", type=\\"acs.k8s.node\\") | limit 0, 10
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1721767203
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetCloudResourceDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceDataRequest) GoString() string {
	return s.String()
}

func (s *GetCloudResourceDataRequest) GetFrom() *int32 {
	return s.From
}

func (s *GetCloudResourceDataRequest) GetQuery() *string {
	return s.Query
}

func (s *GetCloudResourceDataRequest) GetTo() *int32 {
	return s.To
}

func (s *GetCloudResourceDataRequest) SetFrom(v int32) *GetCloudResourceDataRequest {
	s.From = &v
	return s
}

func (s *GetCloudResourceDataRequest) SetQuery(v string) *GetCloudResourceDataRequest {
	s.Query = &v
	return s
}

func (s *GetCloudResourceDataRequest) SetTo(v int32) *GetCloudResourceDataRequest {
	s.To = &v
	return s
}

func (s *GetCloudResourceDataRequest) Validate() error {
	return dara.Validate(s)
}
