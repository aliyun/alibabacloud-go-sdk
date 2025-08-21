// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonGrafanaDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmonGrafanaDashboardsResponseBody
	GetCode() *string
	SetMessage(v string) *GetEmonGrafanaDashboardsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEmonGrafanaDashboardsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEmonGrafanaDashboardsResponseBody
	GetSuccess() *bool
}

type GetEmonGrafanaDashboardsResponseBody struct {
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
	// 1E9D9827-2092-4385-9DA1-FC5A8D1DB3F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonGrafanaDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmonGrafanaDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaDashboardsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEmonGrafanaDashboardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEmonGrafanaDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmonGrafanaDashboardsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetCode(v string) *GetEmonGrafanaDashboardsResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetMessage(v string) *GetEmonGrafanaDashboardsResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetRequestId(v string) *GetEmonGrafanaDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetSuccess(v bool) *GetEmonGrafanaDashboardsResponseBody {
	s.Success = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) Validate() error {
	return dara.Validate(s)
}
