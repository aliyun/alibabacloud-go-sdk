// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDataSourceResponseBody
	GetCode() *string
	SetDataSourceId(v string) *AddDataSourceResponseBody
	GetDataSourceId() *string
	SetMessage(v string) *AddDataSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataSourceResponseBody
	GetSuccess() *bool
}

type AddDataSourceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ds-000*******xzj
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 53167AD3-****-****-92C7-CF69A000BA45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDataSourceResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *AddDataSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataSourceResponseBody) SetCode(v string) *AddDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *AddDataSourceResponseBody) SetDataSourceId(v string) *AddDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *AddDataSourceResponseBody) SetMessage(v string) *AddDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *AddDataSourceResponseBody) SetRequestId(v string) *AddDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataSourceResponseBody) SetSuccess(v bool) *AddDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
