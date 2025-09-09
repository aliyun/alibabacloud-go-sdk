// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManagePrivateRdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ManagePrivateRdsResponseBody
	GetData() *string
	SetRequestId(v string) *ManagePrivateRdsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ManagePrivateRdsResponseBody
	GetSuccess() *bool
}

type ManagePrivateRdsResponseBody struct {
	// The parameter result set returned for the operation that is called for the custom ApsaraDB RDS instance.
	//
	// example:
	//
	// {"requestId":"E63C810A-4A13-47B6-BA67-C0E23A******"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0237BCD2-2C7A-4F86-A766-657AF6******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ManagePrivateRdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManagePrivateRdsResponseBody) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsResponseBody) GetData() *string {
	return s.Data
}

func (s *ManagePrivateRdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManagePrivateRdsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ManagePrivateRdsResponseBody) SetData(v string) *ManagePrivateRdsResponseBody {
	s.Data = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) SetRequestId(v string) *ManagePrivateRdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) SetSuccess(v bool) *ManagePrivateRdsResponseBody {
	s.Success = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) Validate() error {
	return dara.Validate(s)
}
