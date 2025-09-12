// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYzdInstanceTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v bool) *GetYzdInstanceTaskResultResponseBody
	GetCode() *bool
	SetData(v []*GetYzdInstanceTaskResultResponseBodyData) *GetYzdInstanceTaskResultResponseBody
	GetData() []*GetYzdInstanceTaskResultResponseBodyData
	SetMsg(v string) *GetYzdInstanceTaskResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetYzdInstanceTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetYzdInstanceTaskResultResponseBody
	GetSuccess() *bool
}

type GetYzdInstanceTaskResultResponseBody struct {
	Code      *bool                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetYzdInstanceTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                                     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetYzdInstanceTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultResponseBody) GetCode() *bool {
	return s.Code
}

func (s *GetYzdInstanceTaskResultResponseBody) GetData() []*GetYzdInstanceTaskResultResponseBodyData {
	return s.Data
}

func (s *GetYzdInstanceTaskResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetYzdInstanceTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYzdInstanceTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetYzdInstanceTaskResultResponseBody) SetCode(v bool) *GetYzdInstanceTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBody) SetData(v []*GetYzdInstanceTaskResultResponseBodyData) *GetYzdInstanceTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBody) SetMsg(v string) *GetYzdInstanceTaskResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBody) SetRequestId(v string) *GetYzdInstanceTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBody) SetSuccess(v bool) *GetYzdInstanceTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetYzdInstanceTaskResultResponseBodyData struct {
	Appid        *string                                                 `json:"appid,omitempty" xml:"appid,omitempty"`
	Bcid         *string                                                 `json:"bcid,omitempty" xml:"bcid,omitempty"`
	DatasetIds   *string                                                 `json:"datasetIds,omitempty" xml:"datasetIds,omitempty"`
	DownloadUrls []*GetYzdInstanceTaskResultResponseBodyDataDownloadUrls `json:"downloadUrls,omitempty" xml:"downloadUrls,omitempty" type:"Repeated"`
	ResultCnt    *string                                                 `json:"resultCnt,omitempty" xml:"resultCnt,omitempty"`
	Status       *string                                                 `json:"status,omitempty" xml:"status,omitempty"`
	TimeDuration *string                                                 `json:"timeDuration,omitempty" xml:"timeDuration,omitempty"`
}

func (s GetYzdInstanceTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetAppid() *string {
	return s.Appid
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetBcid() *string {
	return s.Bcid
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetDownloadUrls() []*GetYzdInstanceTaskResultResponseBodyDataDownloadUrls {
	return s.DownloadUrls
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetResultCnt() *string {
	return s.ResultCnt
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetYzdInstanceTaskResultResponseBodyData) GetTimeDuration() *string {
	return s.TimeDuration
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetAppid(v string) *GetYzdInstanceTaskResultResponseBodyData {
	s.Appid = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetBcid(v string) *GetYzdInstanceTaskResultResponseBodyData {
	s.Bcid = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetDatasetIds(v string) *GetYzdInstanceTaskResultResponseBodyData {
	s.DatasetIds = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetDownloadUrls(v []*GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) *GetYzdInstanceTaskResultResponseBodyData {
	s.DownloadUrls = v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetResultCnt(v string) *GetYzdInstanceTaskResultResponseBodyData {
	s.ResultCnt = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetStatus(v string) *GetYzdInstanceTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) SetTimeDuration(v string) *GetYzdInstanceTaskResultResponseBodyData {
	s.TimeDuration = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetYzdInstanceTaskResultResponseBodyDataDownloadUrls struct {
	Pwd *string `json:"pwd,omitempty" xml:"pwd,omitempty"`
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) GetPwd() *string {
	return s.Pwd
}

func (s *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) GetUrl() *string {
	return s.Url
}

func (s *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) SetPwd(v string) *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls {
	s.Pwd = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) SetUrl(v string) *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls {
	s.Url = &v
	return s
}

func (s *GetYzdInstanceTaskResultResponseBodyDataDownloadUrls) Validate() error {
	return dara.Validate(s)
}
