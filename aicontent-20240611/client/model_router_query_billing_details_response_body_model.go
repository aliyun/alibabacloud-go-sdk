// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BillingDetailListRespDTO) *ModelRouterQueryBillingDetailsResponseBody
	GetData() *BillingDetailListRespDTO
	SetErrCode(v string) *ModelRouterQueryBillingDetailsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryBillingDetailsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryBillingDetailsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryBillingDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryBillingDetailsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryBillingDetailsResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *BillingDetailListRespDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryBillingDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingDetailsResponseBody) GetData() *BillingDetailListRespDTO {
	return s.Data
}

func (s *ModelRouterQueryBillingDetailsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryBillingDetailsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryBillingDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryBillingDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryBillingDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryBillingDetailsResponseBody) SetData(v *BillingDetailListRespDTO) *ModelRouterQueryBillingDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponseBody) SetErrCode(v string) *ModelRouterQueryBillingDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponseBody) SetErrMessage(v string) *ModelRouterQueryBillingDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryBillingDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponseBody) SetRequestId(v string) *ModelRouterQueryBillingDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponseBody) SetSuccess(v bool) *ModelRouterQueryBillingDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
