// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitiatePptCreationV2ResponseBody
	GetCode() *string
	SetData(v *InitiatePptCreationV2ResponseBodyData) *InitiatePptCreationV2ResponseBody
	GetData() *InitiatePptCreationV2ResponseBodyData
	SetHttpStatusCode(v int32) *InitiatePptCreationV2ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InitiatePptCreationV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitiatePptCreationV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InitiatePptCreationV2ResponseBody
	GetSuccess() *bool
}

type InitiatePptCreationV2ResponseBody struct {
	// The error code.
	//
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *InitiatePptCreationV2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitiatePptCreationV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationV2ResponseBody) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitiatePptCreationV2ResponseBody) GetData() *InitiatePptCreationV2ResponseBodyData {
	return s.Data
}

func (s *InitiatePptCreationV2ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InitiatePptCreationV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitiatePptCreationV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitiatePptCreationV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InitiatePptCreationV2ResponseBody) SetCode(v string) *InitiatePptCreationV2ResponseBody {
	s.Code = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBody) SetData(v *InitiatePptCreationV2ResponseBodyData) *InitiatePptCreationV2ResponseBody {
	s.Data = v
	return s
}

func (s *InitiatePptCreationV2ResponseBody) SetHttpStatusCode(v int32) *InitiatePptCreationV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBody) SetMessage(v string) *InitiatePptCreationV2ResponseBody {
	s.Message = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBody) SetRequestId(v string) *InitiatePptCreationV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBody) SetSuccess(v bool) *InitiatePptCreationV2ResponseBody {
	s.Success = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitiatePptCreationV2ResponseBodyData struct {
	// An alert message related to service usage or status.
	//
	// example:
	//
	// 本月版本内的配送额度已经用尽，超额使用将走按量后付费，下个月配送额度将重新下发；请知晓
	Alert *string `json:"Alert,omitempty" xml:"Alert,omitempty"`
	// The AppKey for initializing the front-end component.
	//
	// example:
	//
	// S1X5ecouBztZelaQ
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 66b25058-d735-47e5-a534-5da93453d3df
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// The URL of the artifact cover.
	//
	// example:
	//
	// http://a.com/xxx.png
	PptArtifactCover *string `json:"PptArtifactCover,omitempty" xml:"PptArtifactCover,omitempty"`
	// The artifact ID. Use this ID to edit the artifact.
	//
	// example:
	//
	// 53059801
	PptArtifactId *string `json:"PptArtifactId,omitempty" xml:"PptArtifactId,omitempty"`
	// The process ID for creating the PPT, used to initialize the front-end component for artifact creation.
	//
	// example:
	//
	// 8485143
	PptProcessId *string `json:"PptProcessId,omitempty" xml:"PptProcessId,omitempty"`
	// The signature for initializing the front-end component.
	//
	// example:
	//
	// dBBGvT0Toje5887Qw+/IwwMNYfk=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s InitiatePptCreationV2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationV2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationV2ResponseBodyData) GetAlert() *string {
	return s.Alert
}

func (s *InitiatePptCreationV2ResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *InitiatePptCreationV2ResponseBodyData) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *InitiatePptCreationV2ResponseBodyData) GetPptArtifactCover() *string {
	return s.PptArtifactCover
}

func (s *InitiatePptCreationV2ResponseBodyData) GetPptArtifactId() *string {
	return s.PptArtifactId
}

func (s *InitiatePptCreationV2ResponseBodyData) GetPptProcessId() *string {
	return s.PptProcessId
}

func (s *InitiatePptCreationV2ResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *InitiatePptCreationV2ResponseBodyData) SetAlert(v string) *InitiatePptCreationV2ResponseBodyData {
	s.Alert = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) SetAppKey(v string) *InitiatePptCreationV2ResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) SetExportTaskId(v string) *InitiatePptCreationV2ResponseBodyData {
	s.ExportTaskId = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) SetPptArtifactCover(v string) *InitiatePptCreationV2ResponseBodyData {
	s.PptArtifactCover = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) SetPptArtifactId(v string) *InitiatePptCreationV2ResponseBodyData {
	s.PptArtifactId = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) SetPptProcessId(v string) *InitiatePptCreationV2ResponseBodyData {
	s.PptProcessId = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) SetSignature(v string) *InitiatePptCreationV2ResponseBodyData {
	s.Signature = &v
	return s
}

func (s *InitiatePptCreationV2ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
