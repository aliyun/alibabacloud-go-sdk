// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExportLogStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExportLogStateResponseBody
	GetRequestId() *string
	SetCode(v string) *UpdateExportLogStateResponseBody
	GetCode() *string
	SetData(v bool) *UpdateExportLogStateResponseBody
	GetData() *bool
}

type UpdateExportLogStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateExportLogStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExportLogStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExportLogStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExportLogStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateExportLogStateResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateExportLogStateResponseBody) SetRequestId(v string) *UpdateExportLogStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExportLogStateResponseBody) SetCode(v string) *UpdateExportLogStateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExportLogStateResponseBody) SetData(v bool) *UpdateExportLogStateResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateExportLogStateResponseBody) Validate() error {
	return dara.Validate(s)
}
