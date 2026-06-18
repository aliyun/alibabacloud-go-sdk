// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcuLvsIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcuLvsIpResponseBody
	GetCode() *string
	SetData(v string) *GetMcuLvsIpResponseBody
	GetData() *string
	SetMessage(v string) *GetMcuLvsIpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcuLvsIpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcuLvsIpResponseBody
	GetSuccess() *bool
}

type GetMcuLvsIpResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// IP information. The value is a JSON string of the List type.
	//
	// example:
	//
	// { "xmculvs":[ 0:{ "port":00 "ip":"0.0.0.0" "name":"lvs1" } ]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcuLvsIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcuLvsIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcuLvsIpResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcuLvsIpResponseBody) GetData() *string {
	return s.Data
}

func (s *GetMcuLvsIpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcuLvsIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcuLvsIpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcuLvsIpResponseBody) SetCode(v string) *GetMcuLvsIpResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetData(v string) *GetMcuLvsIpResponseBody {
	s.Data = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetMessage(v string) *GetMcuLvsIpResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetRequestId(v string) *GetMcuLvsIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetSuccess(v bool) *GetMcuLvsIpResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) Validate() error {
	return dara.Validate(s)
}
