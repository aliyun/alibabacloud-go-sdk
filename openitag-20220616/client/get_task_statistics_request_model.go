// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatType(v string) *GetTaskStatisticsRequest
	GetStatType() *string
}

type GetTaskStatisticsRequest struct {
	// Statistics Type. Valid values:
	//
	// - OPERATORCELL: Operation cell.
	//
	// - ITEM: Single item.
	//
	// example:
	//
	// ITEM
	StatType *string `json:"StatType,omitempty" xml:"StatType,omitempty"`
}

func (s GetTaskStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatisticsRequest) GetStatType() *string {
	return s.StatType
}

func (s *GetTaskStatisticsRequest) SetStatType(v string) *GetTaskStatisticsRequest {
	s.StatType = &v
	return s
}

func (s *GetTaskStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
