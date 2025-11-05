// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteSubAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InviteSubAccountResponseBody
	GetCode() *string
	SetMessage(v string) *InviteSubAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *InviteSubAccountResponseBody
	GetRequestId() *string
	SetResults(v *InviteSubAccountResponseBodyResults) *InviteSubAccountResponseBody
	GetResults() *InviteSubAccountResponseBodyResults
	SetSuccess(v bool) *InviteSubAccountResponseBody
	GetSuccess() *bool
}

type InviteSubAccountResponseBody struct {
	// Error Code: </br>
	//
	// • 200 OK</br>
	//
	// • 1109 System Error</br>
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message</br>
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of invitation sending results
	Results *InviteSubAccountResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// Candidate Values: True/False, this value states if the current API calling action is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountResponseBody) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *InviteSubAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InviteSubAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InviteSubAccountResponseBody) GetResults() *InviteSubAccountResponseBodyResults {
	return s.Results
}

func (s *InviteSubAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InviteSubAccountResponseBody) SetCode(v string) *InviteSubAccountResponseBody {
	s.Code = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetMessage(v string) *InviteSubAccountResponseBody {
	s.Message = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetRequestId(v string) *InviteSubAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetResults(v *InviteSubAccountResponseBodyResults) *InviteSubAccountResponseBody {
	s.Results = v
	return s
}

func (s *InviteSubAccountResponseBody) SetSuccess(v bool) *InviteSubAccountResponseBody {
	s.Success = &v
	return s
}

func (s *InviteSubAccountResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InviteSubAccountResponseBodyResults struct {
	Result []*InviteSubAccountResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InviteSubAccountResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountResponseBodyResults) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResults) GetResult() []*InviteSubAccountResponseBodyResultsResult {
	return s.Result
}

func (s *InviteSubAccountResponseBodyResults) SetResult(v []*InviteSubAccountResponseBodyResultsResult) *InviteSubAccountResponseBodyResults {
	s.Result = v
	return s
}

func (s *InviteSubAccountResponseBodyResults) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InviteSubAccountResponseBodyResultsResult struct {
	// Error Code, 200 OK
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message, Notes of Code
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Returning Message of Invitation Results
	Result *InviteSubAccountResponseBodyResultsResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Always true.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubAccountResponseBodyResultsResult) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResultsResult) GetCode() *string {
	return s.Code
}

func (s *InviteSubAccountResponseBodyResultsResult) GetMessage() *string {
	return s.Message
}

func (s *InviteSubAccountResponseBodyResultsResult) GetResult() *InviteSubAccountResponseBodyResultsResultResult {
	return s.Result
}

func (s *InviteSubAccountResponseBodyResultsResult) GetSuccess() *bool {
	return s.Success
}

func (s *InviteSubAccountResponseBodyResultsResult) SetCode(v string) *InviteSubAccountResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetMessage(v string) *InviteSubAccountResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetResult(v *InviteSubAccountResponseBodyResultsResultResult) *InviteSubAccountResponseBodyResultsResult {
	s.Result = v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetSuccess(v bool) *InviteSubAccountResponseBodyResultsResult {
	s.Success = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InviteSubAccountResponseBodyResultsResultResult struct {
	// Valid days of registration URL, count on daily basis.
	//
	// example:
	//
	// 15
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Invitation ID, The invitation status tracking code.
	//
	// example:
	//
	// 12345
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	// URL for Partner Customer Registration.
	//
	// example:
	//
	// http://agency-intl.console.aliyun.com/customer/register?intl=true&fxinfo=-4uT%2FMWHnnUdvr5GXVd1AYK8luTnGgH3M7Y3lSCd5M1fxRwAkViTWtDJDpckh0HL
	RegUrl *string `json:"RegUrl,omitempty" xml:"RegUrl,omitempty"`
}

func (s InviteSubAccountResponseBodyResultsResultResult) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountResponseBodyResultsResultResult) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResultsResultResult) GetDays() *int32 {
	return s.Days
}

func (s *InviteSubAccountResponseBodyResultsResultResult) GetInviteId() *int64 {
	return s.InviteId
}

func (s *InviteSubAccountResponseBodyResultsResultResult) GetRegUrl() *string {
	return s.RegUrl
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetDays(v int32) *InviteSubAccountResponseBodyResultsResultResult {
	s.Days = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetInviteId(v int64) *InviteSubAccountResponseBodyResultsResultResult {
	s.InviteId = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetRegUrl(v string) *InviteSubAccountResponseBodyResultsResultResult {
	s.RegUrl = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) Validate() error {
	return dara.Validate(s)
}
