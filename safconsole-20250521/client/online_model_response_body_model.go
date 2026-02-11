// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *OnlineModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *OnlineModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *OnlineModelResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *OnlineModelResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *OnlineModelResponseBody
	GetSuccess() *bool
}

type OnlineModelResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OnlineModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineModelResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *OnlineModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *OnlineModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineModelResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *OnlineModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OnlineModelResponseBody) SetCode(v int64) *OnlineModelResponseBody {
	s.Code = &v
	return s
}

func (s *OnlineModelResponseBody) SetHttpStatusCode(v int64) *OnlineModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OnlineModelResponseBody) SetRequestId(v string) *OnlineModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineModelResponseBody) SetResultObject(v bool) *OnlineModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *OnlineModelResponseBody) SetSuccess(v bool) *OnlineModelResponseBody {
	s.Success = &v
	return s
}

func (s *OnlineModelResponseBody) Validate() error {
	return dara.Validate(s)
}
