// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLogShipperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotTtl(v int32) *ResetLogShipperRequest
	GetHotTtl() *int32
	SetLogMetaList(v []*ResetLogShipperRequestLogMetaList) *ResetLogShipperRequest
	GetLogMetaList() []*ResetLogShipperRequestLogMetaList
	SetTtl(v int32) *ResetLogShipperRequest
	GetTtl() *int32
}

type ResetLogShipperRequest struct {
	// The global retention period of hot data.
	//
	// >  The value of this parameter must be at least 7 and smaller than the log retention period. Unit: days.
	//
	// example:
	//
	// 7
	HotTtl *int32 `json:"HotTtl,omitempty" xml:"HotTtl,omitempty"`
	// The settings of the log analysis feature.
	LogMetaList []*ResetLogShipperRequestLogMetaList `json:"LogMetaList,omitempty" xml:"LogMetaList,omitempty" type:"Repeated"`
	// The global log retention period.
	//
	// >  This parameter is supported only when the log analysis feature uses the pay-as-you-go billing method.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ResetLogShipperRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetLogShipperRequest) GoString() string {
	return s.String()
}

func (s *ResetLogShipperRequest) GetHotTtl() *int32 {
	return s.HotTtl
}

func (s *ResetLogShipperRequest) GetLogMetaList() []*ResetLogShipperRequestLogMetaList {
	return s.LogMetaList
}

func (s *ResetLogShipperRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *ResetLogShipperRequest) SetHotTtl(v int32) *ResetLogShipperRequest {
	s.HotTtl = &v
	return s
}

func (s *ResetLogShipperRequest) SetLogMetaList(v []*ResetLogShipperRequestLogMetaList) *ResetLogShipperRequest {
	s.LogMetaList = v
	return s
}

func (s *ResetLogShipperRequest) SetTtl(v int32) *ResetLogShipperRequest {
	s.Ttl = &v
	return s
}

func (s *ResetLogShipperRequest) Validate() error {
	if s.LogMetaList != nil {
		for _, item := range s.LogMetaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResetLogShipperRequestLogMetaList struct {
	// The Logstore that you want to configure.
	//
	// >  You can call the [DescribeLogMeta](~~DescribeLogMeta~~) operation to query the Logstore.
	//
	// example:
	//
	// sas-security-log
	ConfigLogStore *string `json:"ConfigLogStore,omitempty" xml:"ConfigLogStore,omitempty"`
	// The retention period of hot data in the Logstore.
	//
	// >  The value of this parameter must be at least 7 and smaller than the log retention period. Unit: days. If you specify this parameter for the Logstore, the global retention period of hot data specified by the HotTtl parameter is overwritten.
	//
	// example:
	//
	// 7
	HotTtl *int32 `json:"HotTtl,omitempty" xml:"HotTtl,omitempty"`
	// The status of the log analysis feature. Valid values:
	//
	// 	- **disabled**
	//
	// 	- **enabled**
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The log retention period of the Logstore.
	//
	// >  If you specify this parameter for the Logstore, the global log retention period specified by the Ttl parameter is overwritten.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ResetLogShipperRequestLogMetaList) String() string {
	return dara.Prettify(s)
}

func (s ResetLogShipperRequestLogMetaList) GoString() string {
	return s.String()
}

func (s *ResetLogShipperRequestLogMetaList) GetConfigLogStore() *string {
	return s.ConfigLogStore
}

func (s *ResetLogShipperRequestLogMetaList) GetHotTtl() *int32 {
	return s.HotTtl
}

func (s *ResetLogShipperRequestLogMetaList) GetStatus() *string {
	return s.Status
}

func (s *ResetLogShipperRequestLogMetaList) GetTtl() *int32 {
	return s.Ttl
}

func (s *ResetLogShipperRequestLogMetaList) SetConfigLogStore(v string) *ResetLogShipperRequestLogMetaList {
	s.ConfigLogStore = &v
	return s
}

func (s *ResetLogShipperRequestLogMetaList) SetHotTtl(v int32) *ResetLogShipperRequestLogMetaList {
	s.HotTtl = &v
	return s
}

func (s *ResetLogShipperRequestLogMetaList) SetStatus(v string) *ResetLogShipperRequestLogMetaList {
	s.Status = &v
	return s
}

func (s *ResetLogShipperRequestLogMetaList) SetTtl(v int32) *ResetLogShipperRequestLogMetaList {
	s.Ttl = &v
	return s
}

func (s *ResetLogShipperRequestLogMetaList) Validate() error {
	return dara.Validate(s)
}
