// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCompliancePackIdResponseBody
	GetCode() *string
	SetData(v string) *GetCompliancePackIdResponseBody
	GetData() *string
	SetMessage(v string) *GetCompliancePackIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCompliancePackIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCompliancePackIdResponseBody
	GetSuccess() *bool
}

type GetCompliancePackIdResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// cp-9g78b15xxxd0005d5a7
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 855FCC89-0B13-5FC0-AAD2-120878081C1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCompliancePackIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompliancePackIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCompliancePackIdResponseBody) GetData() *string {
	return s.Data
}

func (s *GetCompliancePackIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCompliancePackIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCompliancePackIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCompliancePackIdResponseBody) SetCode(v string) *GetCompliancePackIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetCompliancePackIdResponseBody) SetData(v string) *GetCompliancePackIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetCompliancePackIdResponseBody) SetMessage(v string) *GetCompliancePackIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetCompliancePackIdResponseBody) SetRequestId(v string) *GetCompliancePackIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompliancePackIdResponseBody) SetSuccess(v bool) *GetCompliancePackIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetCompliancePackIdResponseBody) Validate() error {
	return dara.Validate(s)
}
