// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePptArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePptArtifactResponseBody
	GetCode() *string
	SetData(v *DeletePptArtifactResponseBodyData) *DeletePptArtifactResponseBody
	GetData() *DeletePptArtifactResponseBodyData
	SetHttpStatusCode(v int32) *DeletePptArtifactResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePptArtifactResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePptArtifactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePptArtifactResponseBody
	GetSuccess() *bool
}

type DeletePptArtifactResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeletePptArtifactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeletePptArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePptArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePptArtifactResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePptArtifactResponseBody) GetData() *DeletePptArtifactResponseBodyData {
	return s.Data
}

func (s *DeletePptArtifactResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePptArtifactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePptArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePptArtifactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePptArtifactResponseBody) SetCode(v string) *DeletePptArtifactResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePptArtifactResponseBody) SetData(v *DeletePptArtifactResponseBodyData) *DeletePptArtifactResponseBody {
	s.Data = v
	return s
}

func (s *DeletePptArtifactResponseBody) SetHttpStatusCode(v int32) *DeletePptArtifactResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePptArtifactResponseBody) SetMessage(v string) *DeletePptArtifactResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePptArtifactResponseBody) SetRequestId(v string) *DeletePptArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePptArtifactResponseBody) SetSuccess(v bool) *DeletePptArtifactResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePptArtifactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeletePptArtifactResponseBodyData struct {
	// example:
	//
	// 5233498
	PptArtifactId *int32 `json:"PptArtifactId,omitempty" xml:"PptArtifactId,omitempty"`
}

func (s DeletePptArtifactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeletePptArtifactResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePptArtifactResponseBodyData) GetPptArtifactId() *int32 {
	return s.PptArtifactId
}

func (s *DeletePptArtifactResponseBodyData) SetPptArtifactId(v int32) *DeletePptArtifactResponseBodyData {
	s.PptArtifactId = &v
	return s
}

func (s *DeletePptArtifactResponseBodyData) Validate() error {
	return dara.Validate(s)
}
