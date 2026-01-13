// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMeteringDailyDetailCmd interface {
	dara.Model
	String() string
	GoString() string
	SetBillPeriod(v string) *DownloadMeteringDailyDetailCmd
	GetBillPeriod() *string
	SetBillingItem(v string) *DownloadMeteringDailyDetailCmd
	GetBillingItem() *string
	SetEndTime(v string) *DownloadMeteringDailyDetailCmd
	GetEndTime() *string
	SetStartTime(v string) *DownloadMeteringDailyDetailCmd
	GetStartTime() *string
	SetSubAccountId(v string) *DownloadMeteringDailyDetailCmd
	GetSubAccountId() *string
}

type DownloadMeteringDailyDetailCmd struct {
	BillPeriod   *string `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	BillingItem  *string `json:"billingItem,omitempty" xml:"billingItem,omitempty"`
	EndTime      *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s DownloadMeteringDailyDetailCmd) String() string {
	return dara.Prettify(s)
}

func (s DownloadMeteringDailyDetailCmd) GoString() string {
	return s.String()
}

func (s *DownloadMeteringDailyDetailCmd) GetBillPeriod() *string {
	return s.BillPeriod
}

func (s *DownloadMeteringDailyDetailCmd) GetBillingItem() *string {
	return s.BillingItem
}

func (s *DownloadMeteringDailyDetailCmd) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadMeteringDailyDetailCmd) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadMeteringDailyDetailCmd) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *DownloadMeteringDailyDetailCmd) SetBillPeriod(v string) *DownloadMeteringDailyDetailCmd {
	s.BillPeriod = &v
	return s
}

func (s *DownloadMeteringDailyDetailCmd) SetBillingItem(v string) *DownloadMeteringDailyDetailCmd {
	s.BillingItem = &v
	return s
}

func (s *DownloadMeteringDailyDetailCmd) SetEndTime(v string) *DownloadMeteringDailyDetailCmd {
	s.EndTime = &v
	return s
}

func (s *DownloadMeteringDailyDetailCmd) SetStartTime(v string) *DownloadMeteringDailyDetailCmd {
	s.StartTime = &v
	return s
}

func (s *DownloadMeteringDailyDetailCmd) SetSubAccountId(v string) *DownloadMeteringDailyDetailCmd {
	s.SubAccountId = &v
	return s
}

func (s *DownloadMeteringDailyDetailCmd) Validate() error {
	return dara.Validate(s)
}
