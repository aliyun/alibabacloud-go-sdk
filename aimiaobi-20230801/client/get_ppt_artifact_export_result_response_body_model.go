// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptArtifactExportResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPptArtifactExportResultResponseBody
	GetCode() *string
	SetData(v *GetPptArtifactExportResultResponseBodyData) *GetPptArtifactExportResultResponseBody
	GetData() *GetPptArtifactExportResultResponseBodyData
	SetHttpStatusCode(v int32) *GetPptArtifactExportResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPptArtifactExportResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPptArtifactExportResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPptArtifactExportResultResponseBody
	GetSuccess() *bool
}

type GetPptArtifactExportResultResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPptArtifactExportResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPptArtifactExportResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactExportResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPptArtifactExportResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPptArtifactExportResultResponseBody) GetData() *GetPptArtifactExportResultResponseBodyData {
	return s.Data
}

func (s *GetPptArtifactExportResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPptArtifactExportResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPptArtifactExportResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPptArtifactExportResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPptArtifactExportResultResponseBody) SetCode(v string) *GetPptArtifactExportResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetPptArtifactExportResultResponseBody) SetData(v *GetPptArtifactExportResultResponseBodyData) *GetPptArtifactExportResultResponseBody {
	s.Data = v
	return s
}

func (s *GetPptArtifactExportResultResponseBody) SetHttpStatusCode(v int32) *GetPptArtifactExportResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPptArtifactExportResultResponseBody) SetMessage(v string) *GetPptArtifactExportResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetPptArtifactExportResultResponseBody) SetRequestId(v string) *GetPptArtifactExportResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPptArtifactExportResultResponseBody) SetSuccess(v bool) *GetPptArtifactExportResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetPptArtifactExportResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPptArtifactExportResultResponseBodyData struct {
	ExportFileLink []*string `json:"ExportFileLink,omitempty" xml:"ExportFileLink,omitempty" type:"Repeated"`
}

func (s GetPptArtifactExportResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactExportResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPptArtifactExportResultResponseBodyData) GetExportFileLink() []*string {
	return s.ExportFileLink
}

func (s *GetPptArtifactExportResultResponseBodyData) SetExportFileLink(v []*string) *GetPptArtifactExportResultResponseBodyData {
	s.ExportFileLink = v
	return s
}

func (s *GetPptArtifactExportResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
