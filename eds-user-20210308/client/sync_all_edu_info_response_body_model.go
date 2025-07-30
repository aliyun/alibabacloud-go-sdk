// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncAllEduInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncAllEduInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SyncAllEduInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SyncAllEduInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncAllEduInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncAllEduInfoResponseBody
	GetSuccess() *bool
}

type SyncAllEduInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncAllEduInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncAllEduInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncAllEduInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncAllEduInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SyncAllEduInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncAllEduInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncAllEduInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncAllEduInfoResponseBody) SetCode(v string) *SyncAllEduInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetHttpStatusCode(v int32) *SyncAllEduInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetMessage(v string) *SyncAllEduInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetRequestId(v string) *SyncAllEduInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetSuccess(v bool) *SyncAllEduInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
