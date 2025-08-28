// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallDetailByCallIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCallDetailByCallIdResponseBody
	GetCode() *string
	SetData(v string) *QueryCallDetailByCallIdResponseBody
	GetData() *string
	SetMessage(v string) *QueryCallDetailByCallIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCallDetailByCallIdResponseBody
	GetRequestId() *string
}

type QueryCallDetailByCallIdResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the call, in the JSON format.
	//
	// 	- **caller**: the calling number.
	//
	// 	- **startDate**: the time when the call was started.
	//
	// 	- **stateDesc**: the description of the call state.
	//
	// 	- **duration**: the call duration. Unit: seconds. The value **0*	- indicates that the user was not connected.
	//
	// 	- **callerShowNumber**: the calling number displayed to the called party.
	//
	// 	- **gmtCreate**: the time when the call request was received.
	//
	// 	- **state**: the call state. The call state is returned by the Internet service provider (ISP) in real time. For more information about call states, see [ISP-returned error codes](https://help.aliyun.com/document_detail/55085.html).
	//
	// 	- **endDate**: the time when the call was ended.
	//
	// 	- **calleeShowNumber**: the number displayed to the called party.
	//
	// 	- **callee**: the called number.
	//
	// 	- **aRingTime**: the time when Line A started to ring, in the yyyy-MM-dd HH:mm:ss format.
	//
	// 	- **aEndTime**: the time when ringing on Line A ended, in the yyyy-MM-dd HH:mm:ss format.
	//
	// 	- **bRingTime**: the time when Line B started to ring, in the yyyy-MM-dd HH:mm:ss format.
	//
	// 	- **bEndTime**: the time when ringing on Line B ended, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// {"caller":"18767****","startDate":"","stateDesc":"502","duration":0,"callerShowNumber":"05344****","gmtCreate":"2017-11-27 20:09:06","state":"502","endDate":"","calleeShowNumber":"053447****","bRingTime":"2022-01-01 12:02:00"，"bEndTime":"2022-01-01 12:02:28"，"callee":"1373546****"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s QueryCallDetailByCallIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallDetailByCallIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCallDetailByCallIdResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCallDetailByCallIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCallDetailByCallIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallDetailByCallIdResponseBody) SetCode(v string) *QueryCallDetailByCallIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetData(v string) *QueryCallDetailByCallIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetMessage(v string) *QueryCallDetailByCallIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetRequestId(v string) *QueryCallDetailByCallIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) Validate() error {
	return dara.Validate(s)
}
