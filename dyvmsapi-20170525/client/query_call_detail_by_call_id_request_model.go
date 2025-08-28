// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallDetailByCallIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *QueryCallDetailByCallIdRequest
	GetCallId() *string
	SetOwnerId(v int64) *QueryCallDetailByCallIdRequest
	GetOwnerId() *int64
	SetProdId(v int64) *QueryCallDetailByCallIdRequest
	GetProdId() *int64
	SetQueryDate(v int64) *QueryCallDetailByCallIdRequest
	GetQueryDate() *int64
	SetResourceOwnerAccount(v string) *QueryCallDetailByCallIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCallDetailByCallIdRequest
	GetResourceOwnerId() *int64
}

type QueryCallDetailByCallIdRequest struct {
	// The unique ID of the call.
	//
	// >
	//
	// 	- The CallId parameter is included in the response parameters of the outbound call operation that you call to initiate a call.
	//
	// 	- The date when the call specified by CallId is started must be the same as the date specified by QueryDate.
	//
	// 	- The value of CallId must match the value of ProdId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 116014888060^10281631****
	CallId  *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service ID. Valid values:
	//
	// 	- **11000000300006**: voice notification. You can call the [SingleCallByVoice](https://help.aliyun.com/document_detail/393517.html) operation to send a voice notification of the voice notification file type to the specified number.
	//
	// 	- **11010000138001**: voice verification code. You can call the [SingleCallByTts](https://help.aliyun.com/document_detail/393519.html) operation to send a voice verification code or a text-to-speech (TTS) voice notification to the specified number.
	//
	// 	- **11000000300005**: IVR. You can call the [IvrCall](https://help.aliyun.com/document_detail/393521.html) operation to initiate an interactive voice call to the specified number.
	//
	// 	- **11000000300009**: Session Initiation Protocol (SIP) call.
	//
	// 	- **11030000180001**: intelligent outbound call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11000000300006
	ProdId *int64 `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	// The time at which call details are queried. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > The system queries the call records that are generated within 24 hours after the specified point in time. For example, if you specify the time 20:00:01 on November 21, 2022, the system queries the call records that are generated for the specified call ID from 20:00:01 on November 21, 2022, to 20:00:01 on November 22, 2022.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671552000000
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCallDetailByCallIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallDetailByCallIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdRequest) GetCallId() *string {
	return s.CallId
}

func (s *QueryCallDetailByCallIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCallDetailByCallIdRequest) GetProdId() *int64 {
	return s.ProdId
}

func (s *QueryCallDetailByCallIdRequest) GetQueryDate() *int64 {
	return s.QueryDate
}

func (s *QueryCallDetailByCallIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCallDetailByCallIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCallDetailByCallIdRequest) SetCallId(v string) *QueryCallDetailByCallIdRequest {
	s.CallId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetOwnerId(v int64) *QueryCallDetailByCallIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetProdId(v int64) *QueryCallDetailByCallIdRequest {
	s.ProdId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetQueryDate(v int64) *QueryCallDetailByCallIdRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetResourceOwnerAccount(v string) *QueryCallDetailByCallIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetResourceOwnerId(v int64) *QueryCallDetailByCallIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) Validate() error {
	return dara.Validate(s)
}
