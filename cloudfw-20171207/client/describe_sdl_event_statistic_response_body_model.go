// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiSensitiveDataCount(v int64) *DescribeSdlEventStatisticResponseBody
	GetAiSensitiveDataCount() *int64
	SetAssetCount(v int64) *DescribeSdlEventStatisticResponseBody
	GetAssetCount() *int64
	SetRequestId(v string) *DescribeSdlEventStatisticResponseBody
	GetRequestId() *string
	SetSensitiveDataCount(v int64) *DescribeSdlEventStatisticResponseBody
	GetSensitiveDataCount() *int64
	SetTotalCount(v int64) *DescribeSdlEventStatisticResponseBody
	GetTotalCount() *int64
	SetTotalTraffic(v int64) *DescribeSdlEventStatisticResponseBody
	GetTotalTraffic() *int64
}

type DescribeSdlEventStatisticResponseBody struct {
	// example:
	//
	// 1
	AiSensitiveDataCount *int64 `json:"AiSensitiveDataCount,omitempty" xml:"AiSensitiveDataCount,omitempty"`
	// example:
	//
	// 32
	AssetCount *int64 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// example:
	//
	// C5DDD596-1191-5F36-A504-8733045A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	SensitiveDataCount *int64 `json:"SensitiveDataCount,omitempty" xml:"SensitiveDataCount,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 0
	TotalTraffic *int64 `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeSdlEventStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventStatisticResponseBody) GetAiSensitiveDataCount() *int64 {
	return s.AiSensitiveDataCount
}

func (s *DescribeSdlEventStatisticResponseBody) GetAssetCount() *int64 {
	return s.AssetCount
}

func (s *DescribeSdlEventStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSdlEventStatisticResponseBody) GetSensitiveDataCount() *int64 {
	return s.SensitiveDataCount
}

func (s *DescribeSdlEventStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSdlEventStatisticResponseBody) GetTotalTraffic() *int64 {
	return s.TotalTraffic
}

func (s *DescribeSdlEventStatisticResponseBody) SetAiSensitiveDataCount(v int64) *DescribeSdlEventStatisticResponseBody {
	s.AiSensitiveDataCount = &v
	return s
}

func (s *DescribeSdlEventStatisticResponseBody) SetAssetCount(v int64) *DescribeSdlEventStatisticResponseBody {
	s.AssetCount = &v
	return s
}

func (s *DescribeSdlEventStatisticResponseBody) SetRequestId(v string) *DescribeSdlEventStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdlEventStatisticResponseBody) SetSensitiveDataCount(v int64) *DescribeSdlEventStatisticResponseBody {
	s.SensitiveDataCount = &v
	return s
}

func (s *DescribeSdlEventStatisticResponseBody) SetTotalCount(v int64) *DescribeSdlEventStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSdlEventStatisticResponseBody) SetTotalTraffic(v int64) *DescribeSdlEventStatisticResponseBody {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeSdlEventStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
