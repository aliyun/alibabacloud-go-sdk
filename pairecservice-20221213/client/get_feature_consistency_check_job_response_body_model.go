// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConsistencyCheckJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetConfig() *string
	SetFeatureConsistencyCheckJobConfigId(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetFeatureConsistencyCheckJobConfigId() *string
	SetFeatureConsistencyCheckJobConfigName(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetFeatureConsistencyCheckJobConfigName() *string
	SetGmtEndTime(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetGmtEndTime() *string
	SetGmtStartTime(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetGmtStartTime() *string
	SetLogs(v []*string) *GetFeatureConsistencyCheckJobResponseBody
	GetLogs() []*string
	SetRequestId(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetFeatureConsistencyCheckJobResponseBody
	GetStatus() *string
}

type GetFeatureConsistencyCheckJobResponseBody struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 5
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// example:
	//
	// feature_consistency_check_1
	FeatureConsistencyCheckJobConfigName *string `json:"FeatureConsistencyCheckJobConfigName,omitempty" xml:"FeatureConsistencyCheckJobConfigName,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtStartTime *string   `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Logs         []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFeatureConsistencyCheckJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetFeatureConsistencyCheckJobConfigName() *string {
	return s.FeatureConsistencyCheckJobConfigName
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFeatureConsistencyCheckJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetConfig(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.Config = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetFeatureConsistencyCheckJobConfigId(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetFeatureConsistencyCheckJobConfigName(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.FeatureConsistencyCheckJobConfigName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetGmtEndTime(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.GmtEndTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetGmtStartTime(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.GmtStartTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetLogs(v []*string) *GetFeatureConsistencyCheckJobResponseBody {
	s.Logs = v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetRequestId(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetStatus(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) Validate() error {
	return dara.Validate(s)
}
