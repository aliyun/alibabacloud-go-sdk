// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentMonitorStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudAgentMonitorStatisticsRequest
	GetCnos() *string
	SetEndHour(v int64) *CloudAgentMonitorStatisticsRequest
	GetEndHour() *int64
	SetEnterpriseId(v int64) *CloudAgentMonitorStatisticsRequest
	GetEnterpriseId() *int64
	SetFields(v string) *CloudAgentMonitorStatisticsRequest
	GetFields() *string
	SetGno(v string) *CloudAgentMonitorStatisticsRequest
	GetGno() *string
	SetIsNeedQueueName(v string) *CloudAgentMonitorStatisticsRequest
	GetIsNeedQueueName() *string
	SetIsUseGno(v int64) *CloudAgentMonitorStatisticsRequest
	GetIsUseGno() *int64
	SetLimit(v int64) *CloudAgentMonitorStatisticsRequest
	GetLimit() *int64
	SetOffset(v int64) *CloudAgentMonitorStatisticsRequest
	GetOffset() *int64
	SetStartHour(v int64) *CloudAgentMonitorStatisticsRequest
	GetStartHour() *int64
}

type CloudAgentMonitorStatisticsRequest struct {
	// 说明：根据座席工号查询指定座席的统计项名称支持按照多个座席工号进行查询，多个座席工号之间使用英文逗号","分隔，默认查询账户下所有座席的统计信息
	//
	// example:
	//
	// 1111
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 统计时间段；正整数，1～24，最大值是24时
	//
	// example:
	//
	// 23
	EndHour *int64 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 统计项名称；取队列实时统计的统计字段名称，多个队列统计项之间使用英文逗号","分隔，默认查询所有统计项
	//
	// example:
	//
	// ""
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// 说明：根据外呼组编号查询指定座席的统计项名称支持按照多个外呼组编号进行查询，多个座席工号之间使用英文逗号","分隔，默认查询账户下所有座席的统计信息
	//
	// example:
	//
	// WH123
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 是否需要返回队列名称；0：不需要（默认） 1：需要
	//
	// example:
	//
	// 0
	IsNeedQueueName *string `json:"IsNeedQueueName,omitempty" xml:"IsNeedQueueName,omitempty"`
	// 是否需要外呼组查询；说明：0 关闭外呼组查询，1 开启外呼组查询。默认为关闭状态
	//
	// example:
	//
	// 0
	IsUseGno *int64 `json:"IsUseGno,omitempty" xml:"IsUseGno,omitempty"`
	// 条数；正整数，默认值是10，最大值是500
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 查询页码数的起始条数；正整数，默认值是0；与limit同步使用，例：offset=50,limit=10,表示查询第6页，每页10条数据，即总记录的第51-60条数据
	//
	// example:
	//
	// 0
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// 统计时间段；正整数，0～23，最大值是23时
	//
	// example:
	//
	// 0
	StartHour *int64 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s CloudAgentMonitorStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentMonitorStatisticsRequest) GoString() string {
	return s.String()
}

func (s *CloudAgentMonitorStatisticsRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudAgentMonitorStatisticsRequest) GetEndHour() *int64 {
	return s.EndHour
}

func (s *CloudAgentMonitorStatisticsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAgentMonitorStatisticsRequest) GetFields() *string {
	return s.Fields
}

func (s *CloudAgentMonitorStatisticsRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudAgentMonitorStatisticsRequest) GetIsNeedQueueName() *string {
	return s.IsNeedQueueName
}

func (s *CloudAgentMonitorStatisticsRequest) GetIsUseGno() *int64 {
	return s.IsUseGno
}

func (s *CloudAgentMonitorStatisticsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudAgentMonitorStatisticsRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *CloudAgentMonitorStatisticsRequest) GetStartHour() *int64 {
	return s.StartHour
}

func (s *CloudAgentMonitorStatisticsRequest) SetCnos(v string) *CloudAgentMonitorStatisticsRequest {
	s.Cnos = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetEndHour(v int64) *CloudAgentMonitorStatisticsRequest {
	s.EndHour = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetEnterpriseId(v int64) *CloudAgentMonitorStatisticsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetFields(v string) *CloudAgentMonitorStatisticsRequest {
	s.Fields = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetGno(v string) *CloudAgentMonitorStatisticsRequest {
	s.Gno = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetIsNeedQueueName(v string) *CloudAgentMonitorStatisticsRequest {
	s.IsNeedQueueName = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetIsUseGno(v int64) *CloudAgentMonitorStatisticsRequest {
	s.IsUseGno = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetLimit(v int64) *CloudAgentMonitorStatisticsRequest {
	s.Limit = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetOffset(v int64) *CloudAgentMonitorStatisticsRequest {
	s.Offset = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) SetStartHour(v int64) *CloudAgentMonitorStatisticsRequest {
	s.StartHour = &v
	return s
}

func (s *CloudAgentMonitorStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
