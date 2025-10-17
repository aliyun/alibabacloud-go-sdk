// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppResponseBody
	GetCode() *int32
	SetData(v *GetAppResponseBodyData) *GetAppResponseBody
	GetData() *GetAppResponseBodyData
	SetMessage(v string) *GetAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppResponseBody
	GetSuccess() *bool
}

type GetAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Not found: appName not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppResponseBody) GetData() *GetAppResponseBodyData {
	return s.Data
}

func (s *GetAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppResponseBody) SetCode(v int32) *GetAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppResponseBody) SetData(v *GetAppResponseBodyData) *GetAppResponseBody {
	s.Data = v
	return s
}

func (s *GetAppResponseBody) SetMessage(v string) *GetAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetSuccess(v bool) *GetAppResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppResponseBodyData struct {
	// AccessToken。
	//
	// example:
	//
	// 2f4ddeab8e344ed68e0402cf9b8502ffv3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 18582193685027xx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// example:
	//
	// 2
	ExecutorNum *int64 `json:"ExecutorNum,omitempty" xml:"ExecutorNum,omitempty"`
	// example:
	//
	// 3402
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 100
	JobNum             *int32 `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
	LabelRouteStrategy *int32 `json:"LabelRouteStrategy,omitempty" xml:"LabelRouteStrategy,omitempty"`
	// example:
	//
	// http://28.***.***.3:80
	Leader *string `json:"Leader,omitempty" xml:"Leader,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// 1000
	MaxJobs *int32  `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 18582193685027xx
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
}

func (s GetAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyData) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetAppResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetAppResponseBodyData) GetAppType() *int32 {
	return s.AppType
}

func (s *GetAppResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetAppResponseBodyData) GetEnableLog() *bool {
	return s.EnableLog
}

func (s *GetAppResponseBodyData) GetExecutorNum() *int64 {
	return s.ExecutorNum
}

func (s *GetAppResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetAppResponseBodyData) GetJobNum() *int32 {
	return s.JobNum
}

func (s *GetAppResponseBodyData) GetLabelRouteStrategy() *int32 {
	return s.LabelRouteStrategy
}

func (s *GetAppResponseBodyData) GetLeader() *string {
	return s.Leader
}

func (s *GetAppResponseBodyData) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *GetAppResponseBodyData) GetMaxJobs() *int32 {
	return s.MaxJobs
}

func (s *GetAppResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetAppResponseBodyData) GetUpdater() *string {
	return s.Updater
}

func (s *GetAppResponseBodyData) SetAccessToken(v string) *GetAppResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *GetAppResponseBodyData) SetAppName(v string) *GetAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyData) SetAppType(v int32) *GetAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetAppResponseBodyData) SetCreator(v string) *GetAppResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetAppResponseBodyData) SetEnableLog(v bool) *GetAppResponseBodyData {
	s.EnableLog = &v
	return s
}

func (s *GetAppResponseBodyData) SetExecutorNum(v int64) *GetAppResponseBodyData {
	s.ExecutorNum = &v
	return s
}

func (s *GetAppResponseBodyData) SetId(v int64) *GetAppResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAppResponseBodyData) SetJobNum(v int32) *GetAppResponseBodyData {
	s.JobNum = &v
	return s
}

func (s *GetAppResponseBodyData) SetLabelRouteStrategy(v int32) *GetAppResponseBodyData {
	s.LabelRouteStrategy = &v
	return s
}

func (s *GetAppResponseBodyData) SetLeader(v string) *GetAppResponseBodyData {
	s.Leader = &v
	return s
}

func (s *GetAppResponseBodyData) SetMaxConcurrency(v int32) *GetAppResponseBodyData {
	s.MaxConcurrency = &v
	return s
}

func (s *GetAppResponseBodyData) SetMaxJobs(v int32) *GetAppResponseBodyData {
	s.MaxJobs = &v
	return s
}

func (s *GetAppResponseBodyData) SetTitle(v string) *GetAppResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetAppResponseBodyData) SetUpdater(v string) *GetAppResponseBodyData {
	s.Updater = &v
	return s
}

func (s *GetAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
