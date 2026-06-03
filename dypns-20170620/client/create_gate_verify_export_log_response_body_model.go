// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGateVerifyExportLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateGateVerifyExportLogResponseBody
	GetRequestId() *string
	SetCode(v string) *CreateGateVerifyExportLogResponseBody
	GetCode() *string
	SetData(v bool) *CreateGateVerifyExportLogResponseBody
	GetData() *bool
}

type CreateGateVerifyExportLogResponseBody struct {
	// example:
	//
	// C19D3BCD-2233-59DF-B459-6E6587C24405
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateGateVerifyExportLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGateVerifyExportLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGateVerifyExportLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGateVerifyExportLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGateVerifyExportLogResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateGateVerifyExportLogResponseBody) SetRequestId(v string) *CreateGateVerifyExportLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGateVerifyExportLogResponseBody) SetCode(v string) *CreateGateVerifyExportLogResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGateVerifyExportLogResponseBody) SetData(v bool) *CreateGateVerifyExportLogResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGateVerifyExportLogResponseBody) Validate() error {
	return dara.Validate(s)
}
