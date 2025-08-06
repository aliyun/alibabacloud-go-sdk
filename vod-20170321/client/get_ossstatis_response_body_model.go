// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxStorageUtilization(v int64) *GetOSSStatisResponseBody
	GetMaxStorageUtilization() *int64
	SetOssStatisList(v *GetOSSStatisResponseBodyOssStatisList) *GetOSSStatisResponseBody
	GetOssStatisList() *GetOSSStatisResponseBodyOssStatisList
	SetRequestId(v string) *GetOSSStatisResponseBody
	GetRequestId() *string
}

type GetOSSStatisResponseBody struct {
	MaxStorageUtilization *int64                                 `json:"MaxStorageUtilization,omitempty" xml:"MaxStorageUtilization,omitempty"`
	OssStatisList         *GetOSSStatisResponseBodyOssStatisList `json:"OssStatisList,omitempty" xml:"OssStatisList,omitempty" type:"Struct"`
	RequestId             *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOSSStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOSSStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSStatisResponseBody) GetMaxStorageUtilization() *int64 {
	return s.MaxStorageUtilization
}

func (s *GetOSSStatisResponseBody) GetOssStatisList() *GetOSSStatisResponseBodyOssStatisList {
	return s.OssStatisList
}

func (s *GetOSSStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOSSStatisResponseBody) SetMaxStorageUtilization(v int64) *GetOSSStatisResponseBody {
	s.MaxStorageUtilization = &v
	return s
}

func (s *GetOSSStatisResponseBody) SetOssStatisList(v *GetOSSStatisResponseBodyOssStatisList) *GetOSSStatisResponseBody {
	s.OssStatisList = v
	return s
}

func (s *GetOSSStatisResponseBody) SetRequestId(v string) *GetOSSStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOSSStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOSSStatisResponseBodyOssStatisList struct {
	OSSMetric []*GetOSSStatisResponseBodyOssStatisListOSSMetric `json:"OSSMetric,omitempty" xml:"OSSMetric,omitempty" type:"Repeated"`
}

func (s GetOSSStatisResponseBodyOssStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetOSSStatisResponseBodyOssStatisList) GoString() string {
	return s.String()
}

func (s *GetOSSStatisResponseBodyOssStatisList) GetOSSMetric() []*GetOSSStatisResponseBodyOssStatisListOSSMetric {
	return s.OSSMetric
}

func (s *GetOSSStatisResponseBodyOssStatisList) SetOSSMetric(v []*GetOSSStatisResponseBodyOssStatisListOSSMetric) *GetOSSStatisResponseBodyOssStatisList {
	s.OSSMetric = v
	return s
}

func (s *GetOSSStatisResponseBodyOssStatisList) Validate() error {
	return dara.Validate(s)
}

type GetOSSStatisResponseBodyOssStatisListOSSMetric struct {
	StatTime           *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
	StatTimeUTC        *string `json:"StatTimeUTC,omitempty" xml:"StatTimeUTC,omitempty"`
	StorageUtilization *int64  `json:"StorageUtilization,omitempty" xml:"StorageUtilization,omitempty"`
}

func (s GetOSSStatisResponseBodyOssStatisListOSSMetric) String() string {
	return dara.Prettify(s)
}

func (s GetOSSStatisResponseBodyOssStatisListOSSMetric) GoString() string {
	return s.String()
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) GetStatTime() *string {
	return s.StatTime
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) GetStatTimeUTC() *string {
	return s.StatTimeUTC
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) GetStorageUtilization() *int64 {
	return s.StorageUtilization
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) SetStatTime(v string) *GetOSSStatisResponseBodyOssStatisListOSSMetric {
	s.StatTime = &v
	return s
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) SetStatTimeUTC(v string) *GetOSSStatisResponseBodyOssStatisListOSSMetric {
	s.StatTimeUTC = &v
	return s
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) SetStorageUtilization(v int64) *GetOSSStatisResponseBodyOssStatisListOSSMetric {
	s.StorageUtilization = &v
	return s
}

func (s *GetOSSStatisResponseBodyOssStatisListOSSMetric) Validate() error {
	return dara.Validate(s)
}
