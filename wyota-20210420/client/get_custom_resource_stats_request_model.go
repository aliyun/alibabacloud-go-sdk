// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomResourceStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMainBizType(v string) *GetCustomResourceStatsRequest
	GetMainBizType() *string
}

type GetCustomResourceStatsRequest struct {
	// The business type. Default value: enterprise.
	//
	// example:
	//
	// enterprise
	MainBizType *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
}

func (s GetCustomResourceStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResourceStatsRequest) GoString() string {
	return s.String()
}

func (s *GetCustomResourceStatsRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *GetCustomResourceStatsRequest) SetMainBizType(v string) *GetCustomResourceStatsRequest {
	s.MainBizType = &v
	return s
}

func (s *GetCustomResourceStatsRequest) Validate() error {
	return dara.Validate(s)
}
