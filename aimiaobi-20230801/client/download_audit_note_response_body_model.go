// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadAuditNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadAuditNoteResponseBody
	GetCode() *string
	SetData(v string) *DownloadAuditNoteResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DownloadAuditNoteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DownloadAuditNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadAuditNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadAuditNoteResponseBody
	GetSuccess() *bool
}

type DownloadAuditNoteResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// http://download.xxx.yyy.zzz
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadAuditNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadAuditNoteResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadAuditNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadAuditNoteResponseBody) GetData() *string {
	return s.Data
}

func (s *DownloadAuditNoteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DownloadAuditNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadAuditNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadAuditNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadAuditNoteResponseBody) SetCode(v string) *DownloadAuditNoteResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadAuditNoteResponseBody) SetData(v string) *DownloadAuditNoteResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadAuditNoteResponseBody) SetHttpStatusCode(v int32) *DownloadAuditNoteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadAuditNoteResponseBody) SetMessage(v string) *DownloadAuditNoteResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadAuditNoteResponseBody) SetRequestId(v string) *DownloadAuditNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadAuditNoteResponseBody) SetSuccess(v bool) *DownloadAuditNoteResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadAuditNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
