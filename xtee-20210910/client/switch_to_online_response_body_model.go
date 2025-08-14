// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchToOnlineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SwitchToOnlineResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *SwitchToOnlineResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SwitchToOnlineResponseBody
	GetMessage() *string
	SetRequestId(v string) *SwitchToOnlineResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SwitchToOnlineResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *SwitchToOnlineResponseBody
	GetSuccess() *bool
}

type SwitchToOnlineResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Indicates whether this operation was successful, `true` means success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SwitchToOnlineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchToOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchToOnlineResponseBody) GetCode() *string {
	return s.Code
}

func (s *SwitchToOnlineResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SwitchToOnlineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SwitchToOnlineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchToOnlineResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SwitchToOnlineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchToOnlineResponseBody) SetCode(v string) *SwitchToOnlineResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchToOnlineResponseBody) SetHttpStatusCode(v string) *SwitchToOnlineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SwitchToOnlineResponseBody) SetMessage(v string) *SwitchToOnlineResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchToOnlineResponseBody) SetRequestId(v string) *SwitchToOnlineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchToOnlineResponseBody) SetResultObject(v bool) *SwitchToOnlineResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SwitchToOnlineResponseBody) SetSuccess(v bool) *SwitchToOnlineResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchToOnlineResponseBody) Validate() error {
	return dara.Validate(s)
}
