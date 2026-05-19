// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteSubResellerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InviteSubResellerResponseBody
	GetCode() *string
	SetMessage(v string) *InviteSubResellerResponseBody
	GetMessage() *string
	SetRequestId(v string) *InviteSubResellerResponseBody
	GetRequestId() *string
	SetResults(v []*InviteSubResellerResponseBodyResults) *InviteSubResellerResponseBody
	GetResults() []*InviteSubResellerResponseBodyResults
	SetSuccess(v bool) *InviteSubResellerResponseBody
	GetSuccess() *bool
}

type InviteSubResellerResponseBody struct {
	// Result code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message.
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 邀请结果信息
	Results []*InviteSubResellerResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubResellerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InviteSubResellerResponseBody) GoString() string {
	return s.String()
}

func (s *InviteSubResellerResponseBody) GetCode() *string {
	return s.Code
}

func (s *InviteSubResellerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InviteSubResellerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InviteSubResellerResponseBody) GetResults() []*InviteSubResellerResponseBodyResults {
	return s.Results
}

func (s *InviteSubResellerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InviteSubResellerResponseBody) SetCode(v string) *InviteSubResellerResponseBody {
	s.Code = &v
	return s
}

func (s *InviteSubResellerResponseBody) SetMessage(v string) *InviteSubResellerResponseBody {
	s.Message = &v
	return s
}

func (s *InviteSubResellerResponseBody) SetRequestId(v string) *InviteSubResellerResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteSubResellerResponseBody) SetResults(v []*InviteSubResellerResponseBodyResults) *InviteSubResellerResponseBody {
	s.Results = v
	return s
}

func (s *InviteSubResellerResponseBody) SetSuccess(v bool) *InviteSubResellerResponseBody {
	s.Success = &v
	return s
}

func (s *InviteSubResellerResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InviteSubResellerResponseBodyResults struct {
	// Error code, 200 OK
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Prompt message, explanation of the code
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Returned invitation result information
	Result *InviteSubResellerResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Always true
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubResellerResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s InviteSubResellerResponseBodyResults) GoString() string {
	return s.String()
}

func (s *InviteSubResellerResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *InviteSubResellerResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *InviteSubResellerResponseBodyResults) GetResult() *InviteSubResellerResponseBodyResultsResult {
	return s.Result
}

func (s *InviteSubResellerResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *InviteSubResellerResponseBodyResults) SetCode(v string) *InviteSubResellerResponseBodyResults {
	s.Code = &v
	return s
}

func (s *InviteSubResellerResponseBodyResults) SetMessage(v string) *InviteSubResellerResponseBodyResults {
	s.Message = &v
	return s
}

func (s *InviteSubResellerResponseBodyResults) SetResult(v *InviteSubResellerResponseBodyResultsResult) *InviteSubResellerResponseBodyResults {
	s.Result = v
	return s
}

func (s *InviteSubResellerResponseBodyResults) SetSuccess(v bool) *InviteSubResellerResponseBodyResults {
	s.Success = &v
	return s
}

func (s *InviteSubResellerResponseBodyResults) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InviteSubResellerResponseBodyResultsResult struct {
	// Validity period of the registration URL, in days
	//
	// example:
	//
	// 15
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Invitation ID, used for querying the invitation status
	//
	// example:
	//
	// 12345
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	// T2 Reseller registration URL
	//
	// example:
	//
	// http://agency-intl.console.aliyun.com/customer/register?intl=true&fxinfo=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	RegUrl *string `json:"RegUrl,omitempty" xml:"RegUrl,omitempty"`
}

func (s InviteSubResellerResponseBodyResultsResult) String() string {
	return dara.Prettify(s)
}

func (s InviteSubResellerResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *InviteSubResellerResponseBodyResultsResult) GetDays() *int32 {
	return s.Days
}

func (s *InviteSubResellerResponseBodyResultsResult) GetInviteId() *int64 {
	return s.InviteId
}

func (s *InviteSubResellerResponseBodyResultsResult) GetRegUrl() *string {
	return s.RegUrl
}

func (s *InviteSubResellerResponseBodyResultsResult) SetDays(v int32) *InviteSubResellerResponseBodyResultsResult {
	s.Days = &v
	return s
}

func (s *InviteSubResellerResponseBodyResultsResult) SetInviteId(v int64) *InviteSubResellerResponseBodyResultsResult {
	s.InviteId = &v
	return s
}

func (s *InviteSubResellerResponseBodyResultsResult) SetRegUrl(v string) *InviteSubResellerResponseBodyResultsResult {
	s.RegUrl = &v
	return s
}

func (s *InviteSubResellerResponseBodyResultsResult) Validate() error {
	return dara.Validate(s)
}
