// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVoiceFileAuditInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVoiceFileAuditInfoResponseBody
	GetCode() *string
	SetData(v []*QueryVoiceFileAuditInfoResponseBodyData) *QueryVoiceFileAuditInfoResponseBody
	GetData() []*QueryVoiceFileAuditInfoResponseBodyData
	SetMessage(v string) *QueryVoiceFileAuditInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryVoiceFileAuditInfoResponseBody
	GetRequestId() *string
}

type QueryVoiceFileAuditInfoResponseBody struct {
	// The response code.
	//
	// The value OK indicates that the request was successful. For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data []*QueryVoiceFileAuditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVoiceFileAuditInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVoiceFileAuditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVoiceFileAuditInfoResponseBody) GetData() []*QueryVoiceFileAuditInfoResponseBodyData {
	return s.Data
}

func (s *QueryVoiceFileAuditInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryVoiceFileAuditInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetCode(v string) *QueryVoiceFileAuditInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetData(v []*QueryVoiceFileAuditInfoResponseBodyData) *QueryVoiceFileAuditInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetMessage(v string) *QueryVoiceFileAuditInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetRequestId(v string) *QueryVoiceFileAuditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryVoiceFileAuditInfoResponseBodyData struct {
	// The review state of the voice file. Valid values:
	//
	// 	- **AUDIT_STATE_INIT**: The voice file was under review.
	//
	// 	- **AUDIT_STATE_PASS**: The voice file was approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The voice file was rejected.
	//
	// 	- **AUDIT_STATE_UPLOADING**: The voice file was approved and is being uploaded.
	//
	// 	- **AUDIT_STATE_REDOING**: The voice file was being reprocessed.
	//
	// 	- **AUDIT_SATE_CANCEL**: The review of the voice file was canceled.
	//
	// 	- **AUDIT_PAUSE**: The review of the voice file was suspended.
	//
	// 	- **AUDIT_ORDER_FINISHED**: The voice file was approved by the ticket system and was waiting for the review of the Internet service provider (ISP).
	//
	// example:
	//
	// AUDIT_STATE_NOT_PASS
	AuditState *string `json:"AuditState,omitempty" xml:"AuditState,omitempty"`
	// The reason why the voice file was rejected.
	//
	// example:
	//
	// This business is not supported
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	// The code of the voice file.
	//
	// example:
	//
	// 8501d2eb-efbb-471f-xxx8-****.wav
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
}

func (s QueryVoiceFileAuditInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryVoiceFileAuditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) GetAuditState() *string {
	return s.AuditState
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) GetRejectInfo() *string {
	return s.RejectInfo
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) SetAuditState(v string) *QueryVoiceFileAuditInfoResponseBodyData {
	s.AuditState = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) SetRejectInfo(v string) *QueryVoiceFileAuditInfoResponseBodyData {
	s.RejectInfo = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) SetVoiceCode(v string) *QueryVoiceFileAuditInfoResponseBodyData {
	s.VoiceCode = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
