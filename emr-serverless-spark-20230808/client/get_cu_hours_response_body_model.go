// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCuHoursResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCuHoursResponseBodyData) *GetCuHoursResponseBody
	GetData() *GetCuHoursResponseBodyData
	SetRequestId(v string) *GetCuHoursResponseBody
	GetRequestId() *string
}

type GetCuHoursResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {
	//
	//     "cuHours": "{2025-01-09 00:00:00=2.033333, 2025-01-09 01:00:00=2.033333, 2025-01-09 02:00:00=2.033333, 2025-01-09 03:00:00=2.033333, 2025-01-09 04:00:00=2.033333, 2025-01-09 05:00:00=2.033333, 2025-01-09 06:00:00=2.033333, 2025-01-09 07:00:00=2.033333, 2025-01-09 08:00:00=2.033333, 2025-01-09 09:00:00=1.933333, 2025-01-09 10:00:00=2.133333, 2025-01-09 11:00:00=3.100000, 2025-01-09 12:00:00=2.900000}"
	//
	// }
	Data *GetCuHoursResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCuHoursResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCuHoursResponseBody) GoString() string {
	return s.String()
}

func (s *GetCuHoursResponseBody) GetData() *GetCuHoursResponseBodyData {
	return s.Data
}

func (s *GetCuHoursResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCuHoursResponseBody) SetData(v *GetCuHoursResponseBodyData) *GetCuHoursResponseBody {
	s.Data = v
	return s
}

func (s *GetCuHoursResponseBody) SetRequestId(v string) *GetCuHoursResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCuHoursResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCuHoursResponseBodyData struct {
	// The number of CU-hours consumed by a queue during a specified cycle. The value is an estimated value. Refer to your Alibaba Cloud bill for the actual number of consumed CU-hours.
	//
	// example:
	//
	// {2025-01-09 00:00:00=2.033333, 2025-01-09 01:00:00=2.033333, 2025-01-09 02:00:00=2.033333, 2025-01-09 03:00:00=2.033333, 2025-01-09 04:00:00=2.033333, 2025-01-09 05:00:00=2.033333, 2025-01-09 06:00:00=2.033333, 2025-01-09 07:00:00=2.033333, 2025-01-09 08:00:00=2.033333, 2025-01-09 09:00:00=1.933333, 2025-01-09 10:00:00=2.133333, 2025-01-09 11:00:00=3.100000, 2025-01-09 12:00:00=2.900000}
	CuHours *string `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
}

func (s GetCuHoursResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCuHoursResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCuHoursResponseBodyData) GetCuHours() *string {
	return s.CuHours
}

func (s *GetCuHoursResponseBodyData) SetCuHours(v string) *GetCuHoursResponseBodyData {
	s.CuHours = &v
	return s
}

func (s *GetCuHoursResponseBodyData) Validate() error {
	return dara.Validate(s)
}
