// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexCurrentValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIndexCurrentValueResponseBody
	GetCode() *string
	SetData(v []map[string]interface{}) *GetIndexCurrentValueResponseBody
	GetData() []map[string]interface{}
	SetMessage(v string) *GetIndexCurrentValueResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIndexCurrentValueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIndexCurrentValueResponseBody
	GetSuccess() *bool
}

type GetIndexCurrentValueResponseBody struct {
	// Status code. A value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of data entries.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIndexCurrentValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexCurrentValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIndexCurrentValueResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *GetIndexCurrentValueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIndexCurrentValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexCurrentValueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIndexCurrentValueResponseBody) SetCode(v string) *GetIndexCurrentValueResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetData(v []map[string]interface{}) *GetIndexCurrentValueResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetMessage(v string) *GetIndexCurrentValueResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetRequestId(v string) *GetIndexCurrentValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetSuccess(v bool) *GetIndexCurrentValueResponseBody {
	s.Success = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) Validate() error {
	return dara.Validate(s)
}
