// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSqlLogStatisticResponseBody
	GetCode() *string
	SetData(v *DescribeSqlLogStatisticResponseBodyData) *DescribeSqlLogStatisticResponseBody
	GetData() *DescribeSqlLogStatisticResponseBodyData
	SetMessage(v string) *DescribeSqlLogStatisticResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlLogStatisticResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSqlLogStatisticResponseBody
	GetSuccess() *string
}

type DescribeSqlLogStatisticResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeSqlLogStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request is successful, **Successful*	- is returned. If the request fails, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlLogStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogStatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSqlLogStatisticResponseBody) GetData() *DescribeSqlLogStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeSqlLogStatisticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlLogStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlLogStatisticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSqlLogStatisticResponseBody) SetCode(v string) *DescribeSqlLogStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBody) SetData(v *DescribeSqlLogStatisticResponseBodyData) *DescribeSqlLogStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlLogStatisticResponseBody) SetMessage(v string) *DescribeSqlLogStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBody) SetRequestId(v string) *DescribeSqlLogStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBody) SetSuccess(v string) *DescribeSqlLogStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSqlLogStatisticResponseBodyData struct {
	// The total cold storage data. Unit: bytes.
	//
	// example:
	//
	// 8585901
	ColdSqlSize *int64 `json:"ColdSqlSize,omitempty" xml:"ColdSqlSize,omitempty"`
	// The free cold storage data. Unit: bytes.
	//
	// example:
	//
	// 5041450
	FreeColdSqlSize *int64 `json:"FreeColdSqlSize,omitempty" xml:"FreeColdSqlSize,omitempty"`
	// The free hot storage data. Unit: bytes.
	//
	// example:
	//
	// 297245
	FreeHotSqlSize *int64 `json:"FreeHotSqlSize,omitempty" xml:"FreeHotSqlSize,omitempty"`
	// The total hot storage data. Unit: bytes.
	//
	// example:
	//
	// 1118042
	HotSqlSize *int64 `json:"HotSqlSize,omitempty" xml:"HotSqlSize,omitempty"`
	// The amount of data imported in the last day. Unit: bytes.
	//
	// example:
	//
	// 297245
	ImportSqlSize *int64 `json:"ImportSqlSize,omitempty" xml:"ImportSqlSize,omitempty"`
	// The timestamp in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1712568564928
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total storage data (cold data + hot data).
	//
	// example:
	//
	// 9703943
	TotalSqlSize *int64 `json:"TotalSqlSize,omitempty" xml:"TotalSqlSize,omitempty"`
}

func (s DescribeSqlLogStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetColdSqlSize() *int64 {
	return s.ColdSqlSize
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetFreeColdSqlSize() *int64 {
	return s.FreeColdSqlSize
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetFreeHotSqlSize() *int64 {
	return s.FreeHotSqlSize
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetHotSqlSize() *int64 {
	return s.HotSqlSize
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetImportSqlSize() *int64 {
	return s.ImportSqlSize
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSqlLogStatisticResponseBodyData) GetTotalSqlSize() *int64 {
	return s.TotalSqlSize
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetColdSqlSize(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.ColdSqlSize = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetFreeColdSqlSize(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.FreeColdSqlSize = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetFreeHotSqlSize(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.FreeHotSqlSize = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetHotSqlSize(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.HotSqlSize = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetImportSqlSize(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.ImportSqlSize = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetTimestamp(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) SetTotalSqlSize(v int64) *DescribeSqlLogStatisticResponseBodyData {
	s.TotalSqlSize = &v
	return s
}

func (s *DescribeSqlLogStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
