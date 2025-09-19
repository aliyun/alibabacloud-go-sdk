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
	SetStartTimestamp(v int64) *GetDataTrendRequest
	GetStartTimestamp() *int64
}

type GetDataTrendRequest struct {
	// The type of the security data that you want to query. Valid values:
	//
	// 	- **HC_NEW**: the number of new baseline risks.
	//
	// 	- **HC_OPERATE**: the number of handled baseline risks.
	//
	// 	- **VUL_NEW**: the number of new vulnerabilities.
	//
	// 	- **VUL_OPERATE**: the number of handled vulnerabilities.
	//
	// 	- **SUSP_NEW**: the number of new alerts.
	//
	// 	- **SUSP_OPERATE**: the number of handled alerts.
	//
	// This parameter is required.
	//
	// example:
	//
	// HC_NEW,HC_OPERATE
	BizTypes *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721923200000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The interval of the data that you want to query. Unit: milliseconds.
	//
	// >  The minimum value is 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86400000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
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

func (s *GetDataTrendRequest) SetStartTimestamp(v int64) *GetDataTrendRequest {
	s.StartTimestamp = &v
	return s
}

func (s *GetDataTrendRequest) Validate() error {
	return dara.Validate(s)
}
