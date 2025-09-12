// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrowdDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCrowdDatasetResponseBody
	GetCode() *string
	SetData(v *GetCrowdDatasetResponseBodyData) *GetCrowdDatasetResponseBody
	GetData() *GetCrowdDatasetResponseBodyData
	SetMsg(v string) *GetCrowdDatasetResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetCrowdDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCrowdDatasetResponseBody
	GetSuccess() *bool
}

type GetCrowdDatasetResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCrowdDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                          `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCrowdDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCrowdDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrowdDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCrowdDatasetResponseBody) GetData() *GetCrowdDatasetResponseBodyData {
	return s.Data
}

func (s *GetCrowdDatasetResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetCrowdDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCrowdDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCrowdDatasetResponseBody) SetCode(v string) *GetCrowdDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *GetCrowdDatasetResponseBody) SetData(v *GetCrowdDatasetResponseBodyData) *GetCrowdDatasetResponseBody {
	s.Data = v
	return s
}

func (s *GetCrowdDatasetResponseBody) SetMsg(v string) *GetCrowdDatasetResponseBody {
	s.Msg = &v
	return s
}

func (s *GetCrowdDatasetResponseBody) SetRequestId(v string) *GetCrowdDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrowdDatasetResponseBody) SetSuccess(v bool) *GetCrowdDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *GetCrowdDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCrowdDatasetResponseBodyData struct {
	AppId          *string `json:"appId,omitempty" xml:"appId,omitempty"`
	CrowdDatasetId *string `json:"crowdDatasetId,omitempty" xml:"crowdDatasetId,omitempty"`
	DatasetIds     *string `json:"datasetIds,omitempty" xml:"datasetIds,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	UploadStatus   *string `json:"uploadStatus,omitempty" xml:"uploadStatus,omitempty"`
}

func (s GetCrowdDatasetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCrowdDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCrowdDatasetResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetCrowdDatasetResponseBodyData) GetCrowdDatasetId() *string {
	return s.CrowdDatasetId
}

func (s *GetCrowdDatasetResponseBodyData) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *GetCrowdDatasetResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetCrowdDatasetResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCrowdDatasetResponseBodyData) GetUploadStatus() *string {
	return s.UploadStatus
}

func (s *GetCrowdDatasetResponseBodyData) SetAppId(v string) *GetCrowdDatasetResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetCrowdDatasetResponseBodyData) SetCrowdDatasetId(v string) *GetCrowdDatasetResponseBodyData {
	s.CrowdDatasetId = &v
	return s
}

func (s *GetCrowdDatasetResponseBodyData) SetDatasetIds(v string) *GetCrowdDatasetResponseBodyData {
	s.DatasetIds = &v
	return s
}

func (s *GetCrowdDatasetResponseBodyData) SetDescription(v string) *GetCrowdDatasetResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCrowdDatasetResponseBodyData) SetName(v string) *GetCrowdDatasetResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCrowdDatasetResponseBodyData) SetUploadStatus(v string) *GetCrowdDatasetResponseBodyData {
	s.UploadStatus = &v
	return s
}

func (s *GetCrowdDatasetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
