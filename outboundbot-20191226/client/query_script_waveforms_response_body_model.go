// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptWaveformsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryScriptWaveformsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QueryScriptWaveformsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryScriptWaveformsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryScriptWaveformsResponseBody
	GetRequestId() *string
	SetScriptWaveforms(v []*QueryScriptWaveformsResponseBodyScriptWaveforms) *QueryScriptWaveformsResponseBody
	GetScriptWaveforms() []*QueryScriptWaveformsResponseBodyScriptWaveforms
	SetSuccess(v bool) *QueryScriptWaveformsResponseBody
	GetSuccess() *bool
}

type QueryScriptWaveformsResponseBody struct {
	Code            *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode  *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message         *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScriptWaveforms []*QueryScriptWaveformsResponseBodyScriptWaveforms `json:"ScriptWaveforms,omitempty" xml:"ScriptWaveforms,omitempty" type:"Repeated"`
	Success         *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryScriptWaveformsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptWaveformsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScriptWaveformsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryScriptWaveformsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryScriptWaveformsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryScriptWaveformsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryScriptWaveformsResponseBody) GetScriptWaveforms() []*QueryScriptWaveformsResponseBodyScriptWaveforms {
	return s.ScriptWaveforms
}

func (s *QueryScriptWaveformsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryScriptWaveformsResponseBody) SetCode(v string) *QueryScriptWaveformsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryScriptWaveformsResponseBody) SetHttpStatusCode(v int32) *QueryScriptWaveformsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryScriptWaveformsResponseBody) SetMessage(v string) *QueryScriptWaveformsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryScriptWaveformsResponseBody) SetRequestId(v string) *QueryScriptWaveformsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScriptWaveformsResponseBody) SetScriptWaveforms(v []*QueryScriptWaveformsResponseBodyScriptWaveforms) *QueryScriptWaveformsResponseBody {
	s.ScriptWaveforms = v
	return s
}

func (s *QueryScriptWaveformsResponseBody) SetSuccess(v bool) *QueryScriptWaveformsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryScriptWaveformsResponseBody) Validate() error {
	if s.ScriptWaveforms != nil {
		for _, item := range s.ScriptWaveforms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryScriptWaveformsResponseBodyScriptWaveforms struct {
	FileId           *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ScriptContent    *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	ScriptId         *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	ScriptWaveformId *string `json:"ScriptWaveformId,omitempty" xml:"ScriptWaveformId,omitempty"`
}

func (s QueryScriptWaveformsResponseBodyScriptWaveforms) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptWaveformsResponseBodyScriptWaveforms) GoString() string {
	return s.String()
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) GetFileId() *string {
	return s.FileId
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) GetFileName() *string {
	return s.FileName
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) GetScriptId() *string {
	return s.ScriptId
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) GetScriptWaveformId() *string {
	return s.ScriptWaveformId
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) SetFileId(v string) *QueryScriptWaveformsResponseBodyScriptWaveforms {
	s.FileId = &v
	return s
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) SetFileName(v string) *QueryScriptWaveformsResponseBodyScriptWaveforms {
	s.FileName = &v
	return s
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) SetScriptContent(v string) *QueryScriptWaveformsResponseBodyScriptWaveforms {
	s.ScriptContent = &v
	return s
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) SetScriptId(v string) *QueryScriptWaveformsResponseBodyScriptWaveforms {
	s.ScriptId = &v
	return s
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) SetScriptWaveformId(v string) *QueryScriptWaveformsResponseBodyScriptWaveforms {
	s.ScriptWaveformId = &v
	return s
}

func (s *QueryScriptWaveformsResponseBodyScriptWaveforms) Validate() error {
	return dara.Validate(s)
}
