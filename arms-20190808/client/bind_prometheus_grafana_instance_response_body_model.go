// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPrometheusGrafanaInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BindPrometheusGrafanaInstanceResponseBody
	GetCode() *int32
	SetData(v bool) *BindPrometheusGrafanaInstanceResponseBody
	GetData() *bool
	SetMessage(v string) *BindPrometheusGrafanaInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindPrometheusGrafanaInstanceResponseBody
	GetRequestId() *string
}

type BindPrometheusGrafanaInstanceResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindPrometheusGrafanaInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindPrometheusGrafanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BindPrometheusGrafanaInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindPrometheusGrafanaInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *BindPrometheusGrafanaInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindPrometheusGrafanaInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindPrometheusGrafanaInstanceResponseBody) SetCode(v int32) *BindPrometheusGrafanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponseBody) SetData(v bool) *BindPrometheusGrafanaInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponseBody) SetMessage(v string) *BindPrometheusGrafanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponseBody) SetRequestId(v string) *BindPrometheusGrafanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
