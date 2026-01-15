// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *OfflineModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *OfflineModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *OfflineModelResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *OfflineModelResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *OfflineModelResponseBody
	GetSuccess() *bool
}

type OfflineModelResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineModelResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *OfflineModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *OfflineModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineModelResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *OfflineModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineModelResponseBody) SetCode(v int64) *OfflineModelResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineModelResponseBody) SetHttpStatusCode(v int64) *OfflineModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflineModelResponseBody) SetRequestId(v string) *OfflineModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineModelResponseBody) SetResultObject(v bool) *OfflineModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *OfflineModelResponseBody) SetSuccess(v bool) *OfflineModelResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineModelResponseBody) Validate() error {
	return dara.Validate(s)
}
