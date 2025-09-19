// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeHoneypotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeHoneypotNodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpgradeHoneypotNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpgradeHoneypotNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeHoneypotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeHoneypotNodeResponseBody
	GetSuccess() *bool
}

type UpgradeHoneypotNodeResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F832F2D0-C4CD-507B-8C14-CE1F25A7*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeHoneypotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeHoneypotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeHoneypotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeHoneypotNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpgradeHoneypotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeHoneypotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeHoneypotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeHoneypotNodeResponseBody) SetCode(v string) *UpgradeHoneypotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeHoneypotNodeResponseBody) SetHttpStatusCode(v int32) *UpgradeHoneypotNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeHoneypotNodeResponseBody) SetMessage(v string) *UpgradeHoneypotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeHoneypotNodeResponseBody) SetRequestId(v string) *UpgradeHoneypotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeHoneypotNodeResponseBody) SetSuccess(v bool) *UpgradeHoneypotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeHoneypotNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
