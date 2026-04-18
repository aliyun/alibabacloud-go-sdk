// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryActiveUserCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDate(v string) *QueryHistoryActiveUserCountRequest
	GetDataDate() *string
}

type QueryHistoryActiveUserCountRequest struct {
	// example:
	//
	// 2023-01-01
	DataDate *string `json:"DataDate,omitempty" xml:"DataDate,omitempty"`
}

func (s QueryHistoryActiveUserCountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserCountRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserCountRequest) GetDataDate() *string {
	return s.DataDate
}

func (s *QueryHistoryActiveUserCountRequest) SetDataDate(v string) *QueryHistoryActiveUserCountRequest {
	s.DataDate = &v
	return s
}

func (s *QueryHistoryActiveUserCountRequest) Validate() error {
	return dara.Validate(s)
}
