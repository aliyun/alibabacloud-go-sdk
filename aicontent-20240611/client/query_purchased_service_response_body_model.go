// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPurchasedServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AliyunConsoleServiceInfoDTO) *QueryPurchasedServiceResponseBody
	GetData() []*AliyunConsoleServiceInfoDTO
	SetErrCode(v string) *QueryPurchasedServiceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryPurchasedServiceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryPurchasedServiceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryPurchasedServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPurchasedServiceResponseBody
	GetSuccess() *bool
}

type QueryPurchasedServiceResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleServiceInfoDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPurchasedServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPurchasedServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPurchasedServiceResponseBody) GetData() []*AliyunConsoleServiceInfoDTO {
	return s.Data
}

func (s *QueryPurchasedServiceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryPurchasedServiceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryPurchasedServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryPurchasedServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPurchasedServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPurchasedServiceResponseBody) SetData(v []*AliyunConsoleServiceInfoDTO) *QueryPurchasedServiceResponseBody {
	s.Data = v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetErrCode(v string) *QueryPurchasedServiceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetErrMessage(v string) *QueryPurchasedServiceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetHttpStatusCode(v int32) *QueryPurchasedServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetRequestId(v string) *QueryPurchasedServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetSuccess(v bool) *QueryPurchasedServiceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
