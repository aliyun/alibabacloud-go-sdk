// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProvincesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListAllProvincesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllProvincesResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListAllProvincesResponseBody
	GetResult() []*string
	SetStatusCode(v int32) *ListAllProvincesResponseBody
	GetStatusCode() *int32
}

type ListAllProvincesResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 00534880-4397-5134-B212-1030B7A37C27
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListAllProvincesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllProvincesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllProvincesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllProvincesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllProvincesResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListAllProvincesResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllProvincesResponseBody) SetMessage(v string) *ListAllProvincesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllProvincesResponseBody) SetRequestId(v string) *ListAllProvincesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllProvincesResponseBody) SetResult(v []*string) *ListAllProvincesResponseBody {
	s.Result = v
	return s
}

func (s *ListAllProvincesResponseBody) SetStatusCode(v int32) *ListAllProvincesResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListAllProvincesResponseBody) Validate() error {
	return dara.Validate(s)
}
