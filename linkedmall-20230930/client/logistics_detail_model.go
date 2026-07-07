// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogisticsDetail interface {
	dara.Model
	String() string
	GoString() string
	SetOcurrTimeStr(v string) *LogisticsDetail
	GetOcurrTimeStr() *string
	SetStanderdDesc(v string) *LogisticsDetail
	GetStanderdDesc() *string
}

type LogisticsDetail struct {
	// Time when the logistics event occurred
	//
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	OcurrTimeStr *string `json:"ocurrTimeStr,omitempty" xml:"ocurrTimeStr,omitempty"`
	// Description of the logistics event
	//
	// example:
	//
	// 已签收
	StanderdDesc *string `json:"standerdDesc,omitempty" xml:"standerdDesc,omitempty"`
}

func (s LogisticsDetail) String() string {
	return dara.Prettify(s)
}

func (s LogisticsDetail) GoString() string {
	return s.String()
}

func (s *LogisticsDetail) GetOcurrTimeStr() *string {
	return s.OcurrTimeStr
}

func (s *LogisticsDetail) GetStanderdDesc() *string {
	return s.StanderdDesc
}

func (s *LogisticsDetail) SetOcurrTimeStr(v string) *LogisticsDetail {
	s.OcurrTimeStr = &v
	return s
}

func (s *LogisticsDetail) SetStanderdDesc(v string) *LogisticsDetail {
	s.StanderdDesc = &v
	return s
}

func (s *LogisticsDetail) Validate() error {
	return dara.Validate(s)
}
