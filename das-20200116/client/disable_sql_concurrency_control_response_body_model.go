// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlConcurrencyControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableSqlConcurrencyControlResponseBody
	GetCode() *string
	SetData(v string) *DisableSqlConcurrencyControlResponseBody
	GetData() *string
	SetMessage(v string) *DisableSqlConcurrencyControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableSqlConcurrencyControlResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DisableSqlConcurrencyControlResponseBody
	GetSuccess() *string
}

type DisableSqlConcurrencyControlResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	//
	// example:
	//
	// Null
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSqlConcurrencyControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlConcurrencyControlResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSqlConcurrencyControlResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableSqlConcurrencyControlResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableSqlConcurrencyControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableSqlConcurrencyControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSqlConcurrencyControlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DisableSqlConcurrencyControlResponseBody) SetCode(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetData(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetMessage(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetRequestId(v string) *DisableSqlConcurrencyControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetSuccess(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) Validate() error {
	return dara.Validate(s)
}
