// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachHostAccountsFromHostShareKeyResponseBody
	GetRequestId() *string
	SetResults(v []*DetachHostAccountsFromHostShareKeyResponseBodyResults) *DetachHostAccountsFromHostShareKeyResponseBody
	GetResults() []*DetachHostAccountsFromHostShareKeyResponseBodyResults
}

type DetachHostAccountsFromHostShareKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostAccountsFromHostShareKeyResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostAccountsFromHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) GetResults() []*DetachHostAccountsFromHostShareKeyResponseBodyResults {
	return s.Results
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) SetRequestId(v string) *DetachHostAccountsFromHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) SetResults(v []*DetachHostAccountsFromHostShareKeyResponseBodyResults) *DetachHostAccountsFromHostShareKeyResponseBody {
	s.Results = v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetachHostAccountsFromHostShareKeyResponseBodyResults struct {
	// The error code returned. If **OK*	- is returned, the disassociation was successful. If a different error code is returned, the disassociation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	//
	// example:
	//
	// 12407
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the shared key.
	//
	// example:
	//
	// 11
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The host account does not exist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachHostAccountsFromHostShareKeyResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetCode(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetHostAccountId(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.HostAccountId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.HostShareKeyId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
