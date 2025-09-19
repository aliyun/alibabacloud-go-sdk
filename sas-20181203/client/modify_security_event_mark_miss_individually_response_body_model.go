// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityEventMarkMissIndividuallyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySecurityEventMarkMissIndividuallyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifySecurityEventMarkMissIndividuallyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifySecurityEventMarkMissIndividuallyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySecurityEventMarkMissIndividuallyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySecurityEventMarkMissIndividuallyResponseBody
	GetSuccess() *bool
	SetTimeCost(v int64) *ModifySecurityEventMarkMissIndividuallyResponseBody
	GetTimeCost() *int64
}

type ModifySecurityEventMarkMissIndividuallyResponseBody struct {
	// The status code of the alert event.
	//
	// 	- **200**: indicates that the alert event is successfully handled.
	//
	// 	- Other codes: indicates that the alert event fails to be handled.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A37B852F-E346-5FF2-82BD-D1F1DXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time that was consumed to process the request. Unit: milliseconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s ModifySecurityEventMarkMissIndividuallyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityEventMarkMissIndividuallyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) SetCode(v string) *ModifySecurityEventMarkMissIndividuallyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) SetHttpStatusCode(v int32) *ModifySecurityEventMarkMissIndividuallyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) SetMessage(v string) *ModifySecurityEventMarkMissIndividuallyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) SetRequestId(v string) *ModifySecurityEventMarkMissIndividuallyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) SetSuccess(v bool) *ModifySecurityEventMarkMissIndividuallyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) SetTimeCost(v int64) *ModifySecurityEventMarkMissIndividuallyResponseBody {
	s.TimeCost = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponseBody) Validate() error {
	return dara.Validate(s)
}
