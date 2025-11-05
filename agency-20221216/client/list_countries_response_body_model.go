// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCountriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCountriesResponseBody
	GetCode() *string
	SetData(v []*string) *ListCountriesResponseBody
	GetData() []*string
	SetMessage(v string) *ListCountriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCountriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCountriesResponseBody
	GetSuccess() *bool
}

type ListCountriesResponseBody struct {
	// Error Code
	//
	// 	- 200: OK
	//
	// 	- 1109: System error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of Region Code
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Message information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// A747A00F-E096-5244-88B3-3E474BAE3AE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCountriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCountriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCountriesResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListCountriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCountriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCountriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCountriesResponseBody) SetCode(v string) *ListCountriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCountriesResponseBody) SetData(v []*string) *ListCountriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCountriesResponseBody) SetMessage(v string) *ListCountriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCountriesResponseBody) SetRequestId(v string) *ListCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCountriesResponseBody) SetSuccess(v bool) *ListCountriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCountriesResponseBody) Validate() error {
	return dara.Validate(s)
}
