// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySmsSignResponseBody
	GetCode() *string
	SetCreateDate(v string) *QuerySmsSignResponseBody
	GetCreateDate() *string
	SetMessage(v string) *QuerySmsSignResponseBody
	GetMessage() *string
	SetReason(v string) *QuerySmsSignResponseBody
	GetReason() *string
	SetRequestId(v string) *QuerySmsSignResponseBody
	GetRequestId() *string
	SetSignName(v string) *QuerySmsSignResponseBody
	GetSignName() *string
	SetSignStatus(v int32) *QuerySmsSignResponseBody
	GetSignStatus() *int32
}

type QuerySmsSignResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The date and time when the signature was created.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The remarks of the review. Valid values:
	//
	// 	- If the signature is in the **Approved*	- or **Pending Approval*	- state, No Remarks is returned.
	//
	// 	- If the signature is in the **Not Approved*	- state, the reason why the signature is rejected is returned.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The status of the signature. Valid values:
	//
	// 	- **0**: The signature is pending approval.
	//
	// 	- **1**: The signature is approved.
	//
	// 	- **2**: The signature is rejected. The Reason parameter indicates the reason why the signature is rejected.
	//
	// 	- **10**: The signature is cancelled.
	//
	// example:
	//
	// 1
	SignStatus *int32 `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
}

func (s QuerySmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsSignResponseBody) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QuerySmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsSignResponseBody) GetReason() *string {
	return s.Reason
}

func (s *QuerySmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsSignResponseBody) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsSignResponseBody) GetSignStatus() *int32 {
	return s.SignStatus
}

func (s *QuerySmsSignResponseBody) SetCode(v string) *QuerySmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetCreateDate(v string) *QuerySmsSignResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetMessage(v string) *QuerySmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetReason(v string) *QuerySmsSignResponseBody {
	s.Reason = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetRequestId(v string) *QuerySmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetSignName(v string) *QuerySmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetSignStatus(v int32) *QuerySmsSignResponseBody {
	s.SignStatus = &v
	return s
}

func (s *QuerySmsSignResponseBody) Validate() error {
	return dara.Validate(s)
}
