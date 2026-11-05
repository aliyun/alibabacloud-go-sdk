// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypes(v string) *GetDataTrendRequest
	GetBizTypes() *string
	SetEndTimestamp(v int64) *GetDataTrendRequest
	GetEndTimestamp() *int64
	SetInterval(v int32) *GetDataTrendRequest
	GetInterval() *int32
	SetResourceDirectoryAccountId(v int64) *GetDataTrendRequest
	GetResourceDirectoryAccountId() *int64
	SetStartTimestamp(v int64) *GetDataTrendRequest
	GetStartTimestamp() *int64
}

type GetDataTrendRequest struct {
	// The security data type. Valid values:
	//
	// - **HC_NEW**: the number of new baseline risks
	//
	// - **HC_OPERATE**: the number of handled baseline risks
	//
	// - **VUL_NEW**: the number of new vulnerability risks
	//
	// - **VUL_OPERATE**: the number of handled vulnerability risks
	//
	// - **SUSP_NEW**: the number of new alerts
	//
	// - **SUSP_OPERATE**: the number of handled alerts
	//
	// > You can specify multiple values separated by commas (,). Example: HC_NEW,VUL_NEW.
	//
	// This parameter is required.
	//
	// example:
	//
	// HC_NEW,HC_OPERATE
	BizTypes *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	// The end timestamp of the query. Unit: milliseconds.
	//
	// > The value of StartTimestamp must be earlier than the value of EndTimestamp. Otherwise, the API returns HTTP 400 and the error code -101.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721923200000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The interval for querying data. Unit: milliseconds.
	//
	// > The minimum value is 1000. A maximum of 100 data points are returned. Excess data points are not returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86400000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The Alibaba Cloud account ID of member accounts in the resource folder.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The start timestamp of the query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1687334501000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s GetDataTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrendRequest) GoString() string {
	return s.String()
}

func (s *GetDataTrendRequest) GetBizTypes() *string {
	return s.BizTypes
}

func (s *GetDataTrendRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *GetDataTrendRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *GetDataTrendRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetDataTrendRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *GetDataTrendRequest) SetBizTypes(v string) *GetDataTrendRequest {
	s.BizTypes = &v
	return s
}

func (s *GetDataTrendRequest) SetEndTimestamp(v int64) *GetDataTrendRequest {
	s.EndTimestamp = &v
	return s
}

func (s *GetDataTrendRequest) SetInterval(v int32) *GetDataTrendRequest {
	s.Interval = &v
	return s
}

func (s *GetDataTrendRequest) SetResourceDirectoryAccountId(v int64) *GetDataTrendRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetDataTrendRequest) SetStartTimestamp(v int64) *GetDataTrendRequest {
	s.StartTimestamp = &v
	return s
}

func (s *GetDataTrendRequest) Validate() error {
	return dara.Validate(s)
}
