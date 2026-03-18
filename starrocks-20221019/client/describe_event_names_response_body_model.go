// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *DescribeEventNamesResponseBody
	GetData() []*string
	SetErrCode(v string) *DescribeEventNamesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeEventNamesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeEventNamesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeEventNamesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEventNamesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeEventNamesResponseBody
	GetTotal() *int32
}

type DescribeEventNamesResponseBody struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeEventNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventNamesResponseBody) GetData() []*string {
	return s.Data
}

func (s *DescribeEventNamesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeEventNamesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeEventNamesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeEventNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventNamesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventNamesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeEventNamesResponseBody) SetData(v []*string) *DescribeEventNamesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventNamesResponseBody) SetErrCode(v string) *DescribeEventNamesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeEventNamesResponseBody) SetErrMessage(v string) *DescribeEventNamesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeEventNamesResponseBody) SetHttpStatusCode(v int32) *DescribeEventNamesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEventNamesResponseBody) SetRequestId(v string) *DescribeEventNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventNamesResponseBody) SetSuccess(v bool) *DescribeEventNamesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventNamesResponseBody) SetTotal(v int32) *DescribeEventNamesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeEventNamesResponseBody) Validate() error {
	return dara.Validate(s)
}
