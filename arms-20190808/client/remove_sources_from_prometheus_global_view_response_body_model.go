// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSourcesFromPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RemoveSourcesFromPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) *RemoveSourcesFromPrometheusGlobalViewResponseBody
	GetData() *RemoveSourcesFromPrometheusGlobalViewResponseBodyData
	SetMessage(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type RemoveSourcesFromPrometheusGlobalViewResponseBody struct {
	// The status code. The HTTP 200 status code indicates a successful request, while others indicate error conditions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *RemoveSourcesFromPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 9319A57D-2D9E-472A-B69B-CF3CD16D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) GetData() *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	return s.Data
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) SetCode(v int32) *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) SetData(v *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) SetMessage(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) SetRequestId(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveSourcesFromPrometheusGlobalViewResponseBodyData struct {
	// The Info-level information.
	//
	// example:
	//
	// {regionId: the region where the aggregation instance resides. globalViewClusterId: the ID of the aggregation instance.
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The additional information.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) GetInfo() *string {
	return s.Info
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) SetInfo(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) SetMsg(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
