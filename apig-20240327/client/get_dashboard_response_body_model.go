// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDashboardResponseBody
	GetCode() *int32
	SetData(v *GetDashboardResponseBodyData) *GetDashboardResponseBody
	GetData() *GetDashboardResponseBodyData
	SetErrorCode(v string) *GetDashboardResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetDashboardResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDashboardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDashboardResponseBody
	GetSuccess() *bool
}

type GetDashboardResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetDashboardResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Ok
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetDashboardResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDashboardResponseBody) GetData() *GetDashboardResponseBodyData {
	return s.Data
}

func (s *GetDashboardResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDashboardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDashboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDashboardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDashboardResponseBody) SetCode(v int32) *GetDashboardResponseBody {
	s.Code = &v
	return s
}

func (s *GetDashboardResponseBody) SetData(v *GetDashboardResponseBodyData) *GetDashboardResponseBody {
	s.Data = v
	return s
}

func (s *GetDashboardResponseBody) SetErrorCode(v string) *GetDashboardResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDashboardResponseBody) SetMessage(v string) *GetDashboardResponseBody {
	s.Message = &v
	return s
}

func (s *GetDashboardResponseBody) SetRequestId(v string) *GetDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDashboardResponseBody) SetSuccess(v bool) *GetDashboardResponseBody {
	s.Success = &v
	return s
}

func (s *GetDashboardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDashboardResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// gw-co370icmjeu****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The dashboard name.
	//
	// example:
	//
	// PLUGIN
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The dashboard title.
	//
	// example:
	//
	// APIG Plugin
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The dashboard URL.
	//
	// example:
	//
	// https://sls.console.aliyun.com/lognext/project/xxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDashboardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDashboardResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetDashboardResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetDashboardResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetDashboardResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDashboardResponseBodyData) SetGatewayId(v string) *GetDashboardResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetDashboardResponseBodyData) SetName(v string) *GetDashboardResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDashboardResponseBodyData) SetTitle(v string) *GetDashboardResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetDashboardResponseBodyData) SetUrl(v string) *GetDashboardResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDashboardResponseBodyData) Validate() error {
	return dara.Validate(s)
}
