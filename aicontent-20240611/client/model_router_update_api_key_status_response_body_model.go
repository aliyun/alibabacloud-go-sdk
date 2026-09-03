// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateApiKeyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ModelRouterUpdateApiKeyStatusResponseBody
	GetCode() *int64
	SetData(v *ModelRouterUpdateApiKeyStatusResponseBodyData) *ModelRouterUpdateApiKeyStatusResponseBody
	GetData() *ModelRouterUpdateApiKeyStatusResponseBodyData
	SetMessage(v string) *ModelRouterUpdateApiKeyStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModelRouterUpdateApiKeyStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateApiKeyStatusResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateApiKeyStatusResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The data object.
	Data *ModelRouterUpdateApiKeyStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterUpdateApiKeyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateApiKeyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) GetData() *ModelRouterUpdateApiKeyStatusResponseBodyData {
	return s.Data
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) SetCode(v int64) *ModelRouterUpdateApiKeyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) SetData(v *ModelRouterUpdateApiKeyStatusResponseBodyData) *ModelRouterUpdateApiKeyStatusResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) SetMessage(v string) *ModelRouterUpdateApiKeyStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) SetRequestId(v string) *ModelRouterUpdateApiKeyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) SetSuccess(v bool) *ModelRouterUpdateApiKeyStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterUpdateApiKeyStatusResponseBodyData struct {
	// ID
	//
	// example:
	//
	// 3220
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The status of the API key. Valid values:
	//
	// - active: The API key is valid.
	//
	// - disabled: The API key is invalid.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterUpdateApiKeyStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateApiKeyStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateApiKeyStatusResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterUpdateApiKeyStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ModelRouterUpdateApiKeyStatusResponseBodyData) SetId(v int64) *ModelRouterUpdateApiKeyStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBodyData) SetStatus(v string) *ModelRouterUpdateApiKeyStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
