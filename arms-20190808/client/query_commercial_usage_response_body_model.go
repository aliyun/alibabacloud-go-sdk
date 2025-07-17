// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommercialUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryCommercialUsageResponseBody
	GetCode() *int32
	SetData(v *QueryCommercialUsageResponseBodyData) *QueryCommercialUsageResponseBody
	GetData() *QueryCommercialUsageResponseBodyData
	SetHttpStatusCode(v int32) *QueryCommercialUsageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryCommercialUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCommercialUsageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCommercialUsageResponseBody
	GetSuccess() *bool
}

type QueryCommercialUsageResponseBody struct {
	// The response status. Valid values: 2XX: The request is successful. 3XX: A redirection message is returned. 4XX: The request is invalid. 5XX: A server error occurs.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *QueryCommercialUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 99A663CB-8D7B-4B0D-A006-03C8EE38E7BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCommercialUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCommercialUsageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommercialUsageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryCommercialUsageResponseBody) GetData() *QueryCommercialUsageResponseBodyData {
	return s.Data
}

func (s *QueryCommercialUsageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryCommercialUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCommercialUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCommercialUsageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCommercialUsageResponseBody) SetCode(v int32) *QueryCommercialUsageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCommercialUsageResponseBody) SetData(v *QueryCommercialUsageResponseBodyData) *QueryCommercialUsageResponseBody {
	s.Data = v
	return s
}

func (s *QueryCommercialUsageResponseBody) SetHttpStatusCode(v int32) *QueryCommercialUsageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryCommercialUsageResponseBody) SetMessage(v string) *QueryCommercialUsageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCommercialUsageResponseBody) SetRequestId(v string) *QueryCommercialUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommercialUsageResponseBody) SetSuccess(v bool) *QueryCommercialUsageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCommercialUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCommercialUsageResponseBodyData struct {
	// Indicates whether a multi-region query is complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Complete *bool `json:"Complete,omitempty" xml:"Complete,omitempty"`
	// The returned struct.
	Items []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s QueryCommercialUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCommercialUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCommercialUsageResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *QueryCommercialUsageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *QueryCommercialUsageResponseBodyData) SetComplete(v bool) *QueryCommercialUsageResponseBodyData {
	s.Complete = &v
	return s
}

func (s *QueryCommercialUsageResponseBodyData) SetItems(v []map[string]interface{}) *QueryCommercialUsageResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryCommercialUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
