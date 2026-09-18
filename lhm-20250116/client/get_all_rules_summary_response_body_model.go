// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllRulesSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetAllRulesSummaryResponseBody
	GetData() *string
	SetErrCode(v string) *GetAllRulesSummaryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetAllRulesSummaryResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetAllRulesSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAllRulesSummaryResponseBody
	GetSuccess() *bool
}

type GetAllRulesSummaryResponseBody struct {
	// The response data.
	//
	// example:
	//
	// []
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The fault error message encoding.
	//
	// example:
	//
	// None
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// not supported.pos 3084, line 96, column 1, token IDENTIFIER settings
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E871A612-DBD2-53D9-B2A0-723EC30B1823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAllRulesSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllRulesSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllRulesSummaryResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAllRulesSummaryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetAllRulesSummaryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetAllRulesSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllRulesSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAllRulesSummaryResponseBody) SetData(v string) *GetAllRulesSummaryResponseBody {
	s.Data = &v
	return s
}

func (s *GetAllRulesSummaryResponseBody) SetErrCode(v string) *GetAllRulesSummaryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAllRulesSummaryResponseBody) SetErrMessage(v string) *GetAllRulesSummaryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAllRulesSummaryResponseBody) SetRequestId(v string) *GetAllRulesSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllRulesSummaryResponseBody) SetSuccess(v bool) *GetAllRulesSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetAllRulesSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
