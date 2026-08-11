// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQpsStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *GetQpsStatsRequest
	GetQuery() *string
	SetRegionId(v string) *GetQpsStatsRequest
	GetRegionId() *string
}

type GetQpsStatsRequest struct {
	// The query condition. The value is a string in JSON format.
	//
	// >Different query conditions return different protected objects. For more information, see **Query parameter description**.
	//
	// example:
	//
	// {}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetQpsStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQpsStatsRequest) GoString() string {
	return s.String()
}

func (s *GetQpsStatsRequest) GetQuery() *string {
	return s.Query
}

func (s *GetQpsStatsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQpsStatsRequest) SetQuery(v string) *GetQpsStatsRequest {
	s.Query = &v
	return s
}

func (s *GetQpsStatsRequest) SetRegionId(v string) *GetQpsStatsRequest {
	s.RegionId = &v
	return s
}

func (s *GetQpsStatsRequest) Validate() error {
	return dara.Validate(s)
}
