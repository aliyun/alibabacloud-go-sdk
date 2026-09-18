// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSkillAuditRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SaveSkillAuditRecordResponseBody
	GetData() *bool
	SetErrCode(v string) *SaveSkillAuditRecordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SaveSkillAuditRecordResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SaveSkillAuditRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveSkillAuditRecordResponseBody
	GetSuccess() *bool
}

type SaveSkillAuditRecordResponseBody struct {
	// The data body returned by the operation. For information about the fields, see the child parameters.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The error message code.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 016D6CE5-51C6-5767-A8F9-D2818FC56509
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false. If false is returned, check errCode and errMessage for troubleshooting.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveSkillAuditRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSkillAuditRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSkillAuditRecordResponseBody) GetData() *bool {
	return s.Data
}

func (s *SaveSkillAuditRecordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SaveSkillAuditRecordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SaveSkillAuditRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSkillAuditRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveSkillAuditRecordResponseBody) SetData(v bool) *SaveSkillAuditRecordResponseBody {
	s.Data = &v
	return s
}

func (s *SaveSkillAuditRecordResponseBody) SetErrCode(v string) *SaveSkillAuditRecordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SaveSkillAuditRecordResponseBody) SetErrMessage(v string) *SaveSkillAuditRecordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SaveSkillAuditRecordResponseBody) SetRequestId(v string) *SaveSkillAuditRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSkillAuditRecordResponseBody) SetSuccess(v bool) *SaveSkillAuditRecordResponseBody {
	s.Success = &v
	return s
}

func (s *SaveSkillAuditRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
