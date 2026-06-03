// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGateVerifyExportLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveGateVerifyExportLogResponseBody
	GetRequestId() *string
	SetCode(v string) *RemoveGateVerifyExportLogResponseBody
	GetCode() *string
	SetData(v bool) *RemoveGateVerifyExportLogResponseBody
	GetData() *bool
}

type RemoveGateVerifyExportLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s RemoveGateVerifyExportLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveGateVerifyExportLogResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGateVerifyExportLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveGateVerifyExportLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveGateVerifyExportLogResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveGateVerifyExportLogResponseBody) SetRequestId(v string) *RemoveGateVerifyExportLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGateVerifyExportLogResponseBody) SetCode(v string) *RemoveGateVerifyExportLogResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveGateVerifyExportLogResponseBody) SetData(v bool) *RemoveGateVerifyExportLogResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveGateVerifyExportLogResponseBody) Validate() error {
	return dara.Validate(s)
}
