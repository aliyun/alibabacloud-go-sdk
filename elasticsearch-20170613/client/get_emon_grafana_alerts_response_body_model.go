// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonGrafanaAlertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmonGrafanaAlertsResponseBody
	GetCode() *string
	SetMessage(v string) *GetEmonGrafanaAlertsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEmonGrafanaAlertsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEmonGrafanaAlertsResponseBody
	GetSuccess() *bool
}

type GetEmonGrafanaAlertsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 08FA74C7-5654-4309-9729-D555AF587B7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonGrafanaAlertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmonGrafanaAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaAlertsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEmonGrafanaAlertsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEmonGrafanaAlertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmonGrafanaAlertsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEmonGrafanaAlertsResponseBody) SetCode(v string) *GetEmonGrafanaAlertsResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) SetMessage(v string) *GetEmonGrafanaAlertsResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) SetRequestId(v string) *GetEmonGrafanaAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) SetSuccess(v bool) *GetEmonGrafanaAlertsResponseBody {
	s.Success = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) Validate() error {
	return dara.Validate(s)
}
