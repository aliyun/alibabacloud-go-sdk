// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachHostAccountsToHostShareKeyResponseBody
	GetRequestId() *string
	SetResults(v []*AttachHostAccountsToHostShareKeyResponseBodyResults) *AttachHostAccountsToHostShareKeyResponseBody
	GetResults() []*AttachHostAccountsToHostShareKeyResponseBodyResults
}

type AttachHostAccountsToHostShareKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostAccountsToHostShareKeyResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostAccountsToHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) GetResults() []*AttachHostAccountsToHostShareKeyResponseBodyResults {
	return s.Results
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) SetRequestId(v string) *AttachHostAccountsToHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) SetResults(v []*AttachHostAccountsToHostShareKeyResponseBodyResults) *AttachHostAccountsToHostShareKeyResponseBody {
	s.Results = v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) Validate() error {
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

type AttachHostAccountsToHostShareKeyResponseBodyResults struct {
	// The error code returned. If **OK*	- is returned, the association was successful. If another error code is returned, the association failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	//
	// example:
	//
	// 1201
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the shared key.
	//
	// example:
	//
	// 10267
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The host account does not exist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachHostAccountsToHostShareKeyResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetCode(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetHostAccountId(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.HostAccountId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetHostShareKeyId(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.HostShareKeyId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetMessage(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
