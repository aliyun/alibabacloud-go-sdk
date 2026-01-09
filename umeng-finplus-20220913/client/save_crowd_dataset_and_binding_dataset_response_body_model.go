// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCrowdDatasetAndBindingDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveCrowdDatasetAndBindingDatasetResponseBody
	GetCode() *string
	SetData(v *SaveCrowdDatasetAndBindingDatasetResponseBodyData) *SaveCrowdDatasetAndBindingDatasetResponseBody
	GetData() *SaveCrowdDatasetAndBindingDatasetResponseBodyData
	SetMsg(v string) *SaveCrowdDatasetAndBindingDatasetResponseBody
	GetMsg() *string
	SetRequestId(v string) *SaveCrowdDatasetAndBindingDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveCrowdDatasetAndBindingDatasetResponseBody
	GetSuccess() *bool
}

type SaveCrowdDatasetAndBindingDatasetResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SaveCrowdDatasetAndBindingDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                            `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveCrowdDatasetAndBindingDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveCrowdDatasetAndBindingDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) GetData() *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	return s.Data
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) SetCode(v string) *SaveCrowdDatasetAndBindingDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) SetData(v *SaveCrowdDatasetAndBindingDatasetResponseBodyData) *SaveCrowdDatasetAndBindingDatasetResponseBody {
	s.Data = v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) SetMsg(v string) *SaveCrowdDatasetAndBindingDatasetResponseBody {
	s.Msg = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) SetRequestId(v string) *SaveCrowdDatasetAndBindingDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) SetSuccess(v bool) *SaveCrowdDatasetAndBindingDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveCrowdDatasetAndBindingDatasetResponseBodyData struct {
	AppId          *string   `json:"appId,omitempty" xml:"appId,omitempty"`
	CrowdDatasetId *string   `json:"crowdDatasetId,omitempty" xml:"crowdDatasetId,omitempty"`
	DatasetIds     []*string `json:"datasetIds,omitempty" xml:"datasetIds,omitempty" type:"Repeated"`
	Description    *string   `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string   `json:"name,omitempty" xml:"name,omitempty"`
	UploadStatus   *string   `json:"uploadStatus,omitempty" xml:"uploadStatus,omitempty"`
}

func (s SaveCrowdDatasetAndBindingDatasetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SaveCrowdDatasetAndBindingDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) GetCrowdDatasetId() *string {
	return s.CrowdDatasetId
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) GetDatasetIds() []*string {
	return s.DatasetIds
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) GetName() *string {
	return s.Name
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) GetUploadStatus() *string {
	return s.UploadStatus
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) SetAppId(v string) *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	s.AppId = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) SetCrowdDatasetId(v string) *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	s.CrowdDatasetId = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) SetDatasetIds(v []*string) *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	s.DatasetIds = v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) SetDescription(v string) *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	s.Description = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) SetName(v string) *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	s.Name = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) SetUploadStatus(v string) *SaveCrowdDatasetAndBindingDatasetResponseBodyData {
	s.UploadStatus = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
