// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataBatchIngestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataBatchIngestion(v *GetDataBatchIngestionResponseBodyDataBatchIngestion) *GetDataBatchIngestionResponseBody
	GetDataBatchIngestion() *GetDataBatchIngestionResponseBodyDataBatchIngestion
	SetRequestId(v string) *GetDataBatchIngestionResponseBody
	GetRequestId() *string
}

type GetDataBatchIngestionResponseBody struct {
	DataBatchIngestion *GetDataBatchIngestionResponseBodyDataBatchIngestion `json:"DataBatchIngestion,omitempty" xml:"DataBatchIngestion,omitempty" type:"Struct"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataBatchIngestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataBatchIngestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataBatchIngestionResponseBody) GetDataBatchIngestion() *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	return s.DataBatchIngestion
}

func (s *GetDataBatchIngestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataBatchIngestionResponseBody) SetDataBatchIngestion(v *GetDataBatchIngestionResponseBodyDataBatchIngestion) *GetDataBatchIngestionResponseBody {
	s.DataBatchIngestion = v
	return s
}

func (s *GetDataBatchIngestionResponseBody) SetRequestId(v string) *GetDataBatchIngestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataBatchIngestionResponseBody) Validate() error {
	if s.DataBatchIngestion != nil {
		if err := s.DataBatchIngestion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataBatchIngestionResponseBodyDataBatchIngestion struct {
	ApsaraDataIngestionIds []*string `json:"ApsaraDataIngestionIds,omitempty" xml:"ApsaraDataIngestionIds,omitempty" type:"Repeated"`
	// example:
	//
	// enabled。
	AutoScanNew *string `json:"AutoScanNew,omitempty" xml:"AutoScanNew,omitempty"`
	// example:
	//
	// 1733269771123。
	DataBatchIngestionEffectTime *string `json:"DataBatchIngestionEffectTime,omitempty" xml:"DataBatchIngestionEffectTime,omitempty"`
	// example:
	//
	// full。
	DataBatchIngestionMode *string `json:"DataBatchIngestionMode,omitempty" xml:"DataBatchIngestionMode,omitempty"`
	// example:
	//
	// 1733269771123。
	DataBatchIngestionSetTime *string `json:"DataBatchIngestionSetTime,omitempty" xml:"DataBatchIngestionSetTime,omitempty"`
	// example:
	//
	// pending。
	DataBatchIngestionStatus *string                                                              `json:"DataBatchIngestionStatus,omitempty" xml:"DataBatchIngestionStatus,omitempty"`
	DataIngestions           []*GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions `json:"DataIngestions,omitempty" xml:"DataIngestions,omitempty" type:"Repeated"`
	// example:
	//
	// true。
	DataSourceRecognizeEnabled *bool     `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	LogUserIds                 []*string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
	RecommendDataIngestionIds  []*string `json:"RecommendDataIngestionIds,omitempty" xml:"RecommendDataIngestionIds,omitempty" type:"Repeated"`
}

func (s GetDataBatchIngestionResponseBodyDataBatchIngestion) String() string {
	return dara.Prettify(s)
}

func (s GetDataBatchIngestionResponseBodyDataBatchIngestion) GoString() string {
	return s.String()
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetApsaraDataIngestionIds() []*string {
	return s.ApsaraDataIngestionIds
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetAutoScanNew() *string {
	return s.AutoScanNew
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetDataBatchIngestionEffectTime() *string {
	return s.DataBatchIngestionEffectTime
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetDataBatchIngestionMode() *string {
	return s.DataBatchIngestionMode
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetDataBatchIngestionSetTime() *string {
	return s.DataBatchIngestionSetTime
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetDataBatchIngestionStatus() *string {
	return s.DataBatchIngestionStatus
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetDataIngestions() []*GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions {
	return s.DataIngestions
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetLogUserIds() []*string {
	return s.LogUserIds
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) GetRecommendDataIngestionIds() []*string {
	return s.RecommendDataIngestionIds
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetApsaraDataIngestionIds(v []*string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.ApsaraDataIngestionIds = v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetAutoScanNew(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.AutoScanNew = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetDataBatchIngestionEffectTime(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.DataBatchIngestionEffectTime = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetDataBatchIngestionMode(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.DataBatchIngestionMode = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetDataBatchIngestionSetTime(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.DataBatchIngestionSetTime = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetDataBatchIngestionStatus(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.DataBatchIngestionStatus = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetDataIngestions(v []*GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.DataIngestions = v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetDataSourceRecognizeEnabled(v bool) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetLogUserIds(v []*string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.LogUserIds = v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) SetRecommendDataIngestionIds(v []*string) *GetDataBatchIngestionResponseBodyDataBatchIngestion {
	s.RecommendDataIngestionIds = v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestion) Validate() error {
	if s.DataIngestions != nil {
		for _, item := range s.DataIngestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions struct {
	// example:
	//
	// alibaba_cloud_sas_process_ingestion_173326*******。
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
	// example:
	//
	// enabled。
	DataIngestionStatus *string `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	// example:
	//
	// alibaba_cloud_sas_process_log_173326*******。
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// alibaba_cloud。
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) String() string {
	return dara.Prettify(s)
}

func (s GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) GoString() string {
	return s.String()
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) GetDataIngestionStatus() *string {
	return s.DataIngestionStatus
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) GetProductId() *string {
	return s.ProductId
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) GetVendorId() *string {
	return s.VendorId
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) SetDataIngestionId(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions {
	s.DataIngestionId = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) SetDataIngestionStatus(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions {
	s.DataIngestionStatus = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) SetDataSourceId(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions {
	s.DataSourceId = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) SetProductId(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions {
	s.ProductId = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) SetVendorId(v string) *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions {
	s.VendorId = &v
	return s
}

func (s *GetDataBatchIngestionResponseBodyDataBatchIngestionDataIngestions) Validate() error {
	return dara.Validate(s)
}
