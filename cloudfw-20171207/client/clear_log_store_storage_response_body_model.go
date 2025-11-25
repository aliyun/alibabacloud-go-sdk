// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearLogStoreStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpStatusCode(v int32) *ClearLogStoreStorageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ClearLogStoreStorageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ClearLogStoreStorageResponseBody
	GetSuccess() *bool
}

type ClearLogStoreStorageResponseBody struct {
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 21C27710-2DB1-5F2A-8588-72D0541B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearLogStoreStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearLogStoreStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ClearLogStoreStorageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ClearLogStoreStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearLogStoreStorageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ClearLogStoreStorageResponseBody) SetHttpStatusCode(v int32) *ClearLogStoreStorageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ClearLogStoreStorageResponseBody) SetRequestId(v string) *ClearLogStoreStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearLogStoreStorageResponseBody) SetSuccess(v bool) *ClearLogStoreStorageResponseBody {
	s.Success = &v
	return s
}

func (s *ClearLogStoreStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
