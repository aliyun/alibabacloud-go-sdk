// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectionReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInspectionReportResponseBody
	GetCode() *string
	SetData(v interface{}) *GetInspectionReportResponseBody
	GetData() interface{}
	SetMessage(v string) *GetInspectionReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInspectionReportResponseBody
	GetRequestId() *string
}

type GetInspectionReportResponseBody struct {
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInspectionReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInspectionReportResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetInspectionReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInspectionReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInspectionReportResponseBody) SetCode(v string) *GetInspectionReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetInspectionReportResponseBody) SetData(v interface{}) *GetInspectionReportResponseBody {
	s.Data = v
	return s
}

func (s *GetInspectionReportResponseBody) SetMessage(v string) *GetInspectionReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetInspectionReportResponseBody) SetRequestId(v string) *GetInspectionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInspectionReportResponseBody) Validate() error {
	return dara.Validate(s)
}
